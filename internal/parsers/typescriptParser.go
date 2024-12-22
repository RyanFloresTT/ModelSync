package parsers

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/RyanFloresTT/ModelSync/pkg/models"
)

type TypeScriptParser struct{}

func (t TypeScriptParser) Parse(filePath string) (models.IR, error) {
	// Simple regex-based parser for TypeScript interfaces
	src, err := os.ReadFile(filePath)
	if err != nil {
		return models.IR{}, fmt.Errorf("failed to read file: %v", err)
	}

	// Example regex for TypeScript interface (adjust as needed)
	re := regexp.MustCompile(`interface (\w+)\s*{\s*([\s\S]*?)\s*}`)
	matches := re.FindStringSubmatch(string(src))
	if len(matches) < 3 {
		return models.IR{}, fmt.Errorf("failed to parse TypeScript interface")
	}

	ir := models.IR{Name: matches[1]}
	fields := strings.Split(matches[2], "\n")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if field == "" || strings.HasSuffix(field, "}") {
			continue
		}

		parts := strings.Split(field, ":")
		if len(parts) == 2 {
			ir.Fields = append(ir.Fields, models.Field{
				Name: strings.TrimSpace(parts[0]),
				Type: strings.TrimSpace(parts[1]),
			})
		}
	}

	return ir, nil
}
