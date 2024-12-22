package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fsnotify/fsnotify"
)

// IR represents the generic structure of a parsed file.
type IR struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Watcher monitors file changes
func WatchFile(config Config) {
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

// GenerateCode parses the file and writes generated files
func GenerateCode(config Config) {
	// Parse file (mock IR for demo)
	ir := IR{
		Name: "BookResponse",
		Fields: []Field{
			{"Title", "string"},
			{"Author", "string"},
		},
	}

	// Read templates and generate code
	for _, target := range config.Targets {
		tmpl, err := template.ParseFiles(target.TemplatePath)
		if err != nil {
			log.Fatal(err)
		}

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

		err = tmpl.Execute(output, ir)
		if err != nil {
			log.Fatalf("Failed to write to file %s: %v", target.Output, err)
		}

		fmt.Printf("Generated code for %s and wrote to %s\n", target.Language, target.Output)
	}
}

// Config structure
type Config struct {
	WatchFile string   `json:"watchFile"`
	Targets   []Target `json:"targets"`
}

type Target struct {
	Language     string `json:"language"`
	Output       string `json:"output"`
	TemplatePath string `json:"templatePath"`
}

// LoadConfig reads configuration from a file
func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	json.Unmarshal(data, &config)
	return config
}

func main() {
	config := LoadConfig("config.json")
	go WatchFile(config)

	// Keep the program running
	select {}
}
