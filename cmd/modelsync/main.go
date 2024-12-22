package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/RyanFloresTT/ModelSync/models"
	"github.com/RyanFloresTT/ModelSync/parsers"
	"github.com/RyanFloresTT/ModelSync/templates"
	"github.com/fsnotify/fsnotify"
)

// mapType maps a generic type to the target language-specific type
func mapType(language, fieldType string) string {
	mappedType, exists := templates.TypeMappings[language][fieldType]
	if exists {
		return mappedType
	}
	return fieldType // Fallback to original if no mapping exists
}

// applyTypeMapping applies type mapping to the IR fields for a specific language
func applyTypeMapping(ir models.IR, language string) models.IR {
	mappedFields := make([]models.Field, len(ir.Fields))
	for i, field := range ir.Fields {
		mappedFields[i] = models.Field{
			Name: field.Name,
			Type: mapType(language, field.Type),
		}
	}
	return models.IR{
		Name:   ir.Name,
		Fields: mappedFields,
	}
}

// InitProject initializes a new project directory with a default config file
func InitProject(configName string) {
	defaultConfig := models.Config{
		WatchFile: "models/sample.go",
		Targets: []models.Target{
			{
				Language: "typescript",
				Output:   "frontend/src/models/BookResponse.tsx",
			},
			{
				Language: "csharp",
				Output:   "shared/Models/BookResponse.cs",
			},
		},
	}

	data, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		log.Fatalf("Failed to serialize default config: %v", err)
	}

	if err := os.WriteFile(configName, data, 0644); err != nil {
		log.Fatalf("Failed to write config file: %v", err)
	}

	fmt.Printf("Initialized project with config file: %s\n", configName)
	os.Exit(0)
}

// WatchFile monitors the specified file for changes and triggers code generation
func WatchFile(config models.Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(config.WatchFile)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("File changed:", event.Name)
				GenerateCode(config)
			}
		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
}

// WatchConfig monitors the configuration file for changes and reloads the program
func WatchConfig(configPath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(configPath)
	if err != nil {
		log.Fatalf("Failed to watch config file: %v", err)
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("Configuration file changed:", event.Name)
				fmt.Println("Reloading configuration and restarting watchers...")

				// Reload configuration
				newConfig := LoadConfig(configPath)

				// Restart watchers with new configuration
				go WatchFile(newConfig)
			}
		case err := <-watcher.Errors:
			log.Println("Error watching configuration file:", err)
		}
	}
}

// GenerateCode parses the specified file and generates target files
func GenerateCode(config models.Config) {
	// Parse the file using the appropriate parser
	parser, err := parsers.GetParser(config.WatchFile)
	if err != nil {
		log.Fatalf("Error selecting parser: %v", err)
	}

	ir, err := parser.Parse(config.WatchFile)
	if err != nil {
		log.Fatalf("Failed to parse file %s: %v", config.WatchFile, err)
	}

	// Read templates and generate code
	for _, target := range config.Targets {
		tmpl, err := templates.GetTemplate(target.Language)
		if err != nil {
			log.Fatalf("Failed to load template for %s: %v", target.Language, err)
		}

		// Apply type mappings to the IR for the target language
		mappedIR := applyTypeMapping(ir, target.Language)

		// Ensure the directory for the target file exists
		err = os.MkdirAll(filepath.Dir(target.Output), os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory for %s: %v", target.Output, err)
		}

		// Check if the file exists, create it if not
		if _, err := os.Stat(target.Output); os.IsNotExist(err) {
			file, err := os.Create(target.Output)
			if err != nil {
				log.Fatalf("Failed to create file %s: %v", target.Output, err)
			}
			file.Close() // Close immediately, will be reopened for writing below
			fmt.Printf("Created new file: %s\n", target.Output)
		} else {
			fmt.Printf("File exists, overwriting: %s\n", target.Output)
		}

		// Write generated content to the file
		output, err := os.OpenFile(target.Output, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to open file %s for writing: %v", target.Output, err)
		}
		defer output.Close()

		err = tmpl.Execute(output, mappedIR)
		if err != nil {
			log.Fatalf("Failed to write to file %s: %v", target.Output, err)
		}

		fmt.Printf("Generated code for %s and wrote to %s\n", target.Language, target.Output)
	}
}

// LoadConfig reads the configuration from a file, creating a default one if none exists
func LoadConfig(path string) models.Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Configuration file %s does not exist. Creating a default one.\n", path)
		InitProject(path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}
	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse configuration file: %v", err)
	}
	return config
}

func main() {
	// Define default configuration file name
	configFile := "syncConfig.json"

	// Load configuration from file, creating a default one if it doesn't exist
	config := LoadConfig(configFile)

	// Start watching the specified file for changes
	go WatchFile(config)

	// Start watching the configuration file for changes
	go WatchConfig(configFile)

	// Keep the program running indefinitely
	select {}
}
