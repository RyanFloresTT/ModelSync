package parsers

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/RyanFloresTT/ModelSync/models"
)

// ParseGoFile parses a Go file and extracts struct information
func ParseGoFile(filePath string) models.IR {
	// Open the file
	src, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Create the AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, src, parser.AllErrors)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	// Initialize the IR
	var ir models.IR

	// Traverse the AST to find struct types
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for type declarations
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// Check if it's a struct type
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// Populate the IR with struct name and fields
		ir.Name = typeSpec.Name.Name
		for _, field := range structType.Fields.List {
			fieldType := extractType(field.Type)
			for _, fieldName := range field.Names {
				ir.Fields = append(ir.Fields, models.Field{
					Name: fieldName.Name,
					Type: fieldType,
				})
			}
		}
		return false
	})

	return ir
}

type GoParser struct{}

func (g GoParser) Parse(filePath string) (models.IR, error) {
	// Implement parsing logic for Go structs (as shown earlier)
	return ParseGoFile(filePath), nil
}
