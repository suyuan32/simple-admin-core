package model

import (
	"fmt"
	"strings"
)

// Message represents a Proto message (type definition in .api)
type Message struct {
	Name   string
	Fields []*Field
}

// Field represents a message field
type Field struct {
	Name      string
	Type      string
	ProtoType string
	JSONTag   string
	Optional  bool
	Repeated  bool
}

// GoZeroType returns the Go type for Go-Zero .api file
func (f *Field) GoZeroType() string {
	baseType := ""

	switch f.ProtoType {
	case "string":
		baseType = "string"
	case "int32":
		baseType = "int32"
	case "int64":
		baseType = "int64"
	case "uint32":
		baseType = "uint32"
	case "uint64":
		baseType = "uint64"
	case "bool":
		baseType = "bool"
	case "float":
		baseType = "float32"
	case "double":
		baseType = "float64"
	case "bytes":
		baseType = "[]byte"
	default:
		// Custom message type
		baseType = f.Type
	}

	// Handle repeated fields
	if f.Repeated {
		return "[]" + baseType
	}

	// Handle optional fields (use pointer)
	if f.Optional {
		return "*" + baseType
	}

	return baseType
}

// FieldName returns the capitalized field name for Go struct
func (f *Field) FieldName() string {
	if len(f.Name) == 0 {
		return ""
	}
	// Capitalize first letter
	return strings.ToUpper(f.Name[:1]) + f.Name[1:]
}

// Tags returns the struct tags for this field
func (f *Field) Tags() string {
	tag := fmt.Sprintf("`json:\"%s", f.JSONTag)
	if f.Optional {
		tag += ",optional"
	}
	tag += "\"`"
	return tag
}
