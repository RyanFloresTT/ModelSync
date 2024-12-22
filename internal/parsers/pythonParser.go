package parsers

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/RyanFloresTT/ModelSync/pkg/models"
)

type PythonParser struct{}

func (p PythonParser) Parse(filePath string) (models.IR, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return models.IR{}, fmt.Errorf("failed to read file: %v", err)
	}

	// Example regex for Python class
	re := regexp.MustCompile(`class (\w+):\s*def __init__\(self.*?\):\s*([\s\S]*?)\n\s*`)
	matches := re.FindStringSubmatch(string(src))
	if len(matches) < 3 {
		return models.IR{}, fmt.Errorf("failed to parse Python class")
	}

	ir := models.IR{Name: matches[1]}
	fields := strings.Split(matches[2], "\n")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if field == "" || !strings.HasPrefix(field, "self.") {
			continue
		}

		parts := strings.Split(field, "=")
		if len(parts) == 2 {
			ir.Fields = append(ir.Fields, models.Field{
				Name: strings.TrimPrefix(strings.TrimSpace(parts[0]), "self."),
				Type: "unknown", // Python doesn’t enforce types; you may infer from conventions
			})
		}
	}

	return ir, nil
}
