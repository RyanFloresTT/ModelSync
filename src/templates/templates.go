package main

import (
	"fmt"
	"strings"
	"text/template"
)

// Templates contains pre-defined templates for various languages
var Templates = map[string]string{
	"typescript": `// Auto-generated TypeScript interface
export interface {{.Name}} {
{{range .Fields}}	{{.Name}}: {{.Type}};
{{end}}}
`,

	"csharp": `// Auto-generated C# class
public class {{.Name}} {
{{range .Fields}}	public {{.Type}} {{.Name}} { get; set; }
{{end}}}
`,

	"cpp": `// Auto-generated C++ struct
struct {{.Name}} {
{{range .Fields}}	{{.Type}} {{.Name}};
{{end}}};
`,

	"python": `# Auto-generated Python class
class {{.Name}}:
	def __init__(self):
{{range .Fields}}		self.{{.Name}}: {{.Type}} = None
{{end}}`,

	"java": `// Auto-generated Java class
public class {{.Name}} {
{{range .Fields}}	private {{.Type}} {{.Name}};

	public {{.Type}} get{{.Name | title}}() {
		return {{.Name}};
	}

	public void set{{.Name | title}}({{.Type}} {{.Name}}) {
		this.{{.Name}} = {{.Name}};
	}
{{end}}}
`,
}

// TypeMappings maps generic IR types to language-specific types
var TypeMappings = map[string]map[string]string{
	"typescript": {
		"string": "string",
		"int":    "number",
		"float":  "number",
		"bool":   "boolean",
	},
	"csharp": {
		"string": "string",
		"int":    "int",
		"float":  "float",
		"bool":   "bool",
	},
	"cpp": {
		"string": "std::string",
		"int":    "int",
		"float":  "float",
		"bool":   "bool",
	},
	"python": {
		"string": "str",
		"int":    "int",
		"float":  "float",
		"bool":   "bool",
	},
	"java": {
		"string": "String",
		"int":    "int",
		"float":  "float",
		"bool":   "boolean",
	},
}

// GetTemplate retrieves a template string by language
func GetTemplate(language string) (*template.Template, error) {
	tmpl, exists := Templates[language]
	if !exists {
		return nil, fmt.Errorf("template for language %s not found", language)
	}
	return template.New(language).Funcs(template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
	}).Parse(tmpl)
}
