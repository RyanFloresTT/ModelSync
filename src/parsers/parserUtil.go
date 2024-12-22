package parsers

import (
	"fmt"
	"go/ast"
	"path/filepath"

	"github.com/RyanFloresTT/ModelSync/models"
)

func GetParser(filePath string) (Parser, error) {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".go":
		return GoParser{}, nil
	case ".ts":
		return TypeScriptParser{}, nil
	case ".py":
		return PythonParser{}, nil
	default:
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}
}

// extractType extracts the type of a field as a string
func extractType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		return "[]" + extractType(t.Elt)
	case *ast.StarExpr:
		return "*" + extractType(t.X)
	default:
		return "unknown"
	}
}

type Parser interface {
	Parse(filePath string) (models.IR, error)
}
