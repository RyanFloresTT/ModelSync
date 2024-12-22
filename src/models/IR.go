package models

// IR represents the intermediate representation of a parsed file
type IR struct {
	Name   string
	Fields []Field
}

// Field represents a single field in a struct, class, or interface
type Field struct {
	Name string
	Type string
}
