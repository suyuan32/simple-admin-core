package generator

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// TypeConverter converts Proto types to Go-Zero .api types
type TypeConverter struct {
	convertedTypes map[string]*model.Message // Track converted types to avoid duplicates
}

// NewTypeConverter creates a new TypeConverter instance
func NewTypeConverter() *TypeConverter {
	return &TypeConverter{
		convertedTypes: make(map[string]*model.Message),
	}
}

// ConvertMessage converts a Proto message to Go-Zero type definition
// Returns nil if the message has already been converted
func (c *TypeConverter) ConvertMessage(msg *protogen.Message) *model.Message {
	fullName := string(msg.Desc.FullName())

	// Check if already converted
	if existing, ok := c.convertedTypes[fullName]; ok {
		return existing
	}

	message := &model.Message{
		Name:   string(msg.Desc.Name()),
		Fields: []*model.Field{},
	}

	// Convert each field
	for _, field := range msg.Fields {
		convertedField := c.convertField(field)
		if convertedField != nil {
			message.Fields = append(message.Fields, convertedField)
		}
	}

	// Mark as converted
	c.convertedTypes[fullName] = message

	return message
}

// convertField converts a Proto field to a Go-Zero field
func (c *TypeConverter) convertField(field *protogen.Field) *model.Field {
	f := &model.Field{
		Name:      c.toGoFieldName(string(field.Desc.Name())),
		JSONTag:   field.Desc.JSONName(),
		Optional:  field.Desc.HasOptionalKeyword(),
		Repeated:  field.Desc.Cardinality() == protoreflect.Repeated && !field.Desc.IsMap(),
	}

	// Convert type
	f.Type, f.ProtoType = c.convertType(field)

	return f
}

// convertType converts Proto type to Go type
// Returns (goType, protoType)
func (c *TypeConverter) convertType(field *protogen.Field) (string, string) {
	kind := field.Desc.Kind()

	// Handle maps specially
	if field.Desc.IsMap() {
		return c.convertMapType(field), "map"
	}

	switch kind {
	case protoreflect.BoolKind:
		return "bool", "bool"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32", "int32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64", "int64"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32", "uint32"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64", "uint64"
	case protoreflect.FloatKind:
		return "float32", "float"
	case protoreflect.DoubleKind:
		return "float64", "double"
	case protoreflect.StringKind:
		return "string", "string"
	case protoreflect.BytesKind:
		return "[]byte", "bytes"
	case protoreflect.EnumKind:
		// In Go-Zero, enums are typically represented as strings or int32
		// We'll use string for better readability
		return "string", "enum"
	case protoreflect.MessageKind:
		// Custom message type
		msgName := string(field.Message.Desc.Name())
		return msgName, "message"
	default:
		return "interface{}", "unknown"
	}
}

// convertMapType handles Proto map<K,V> types
// Example: map<string, int32> -> map[string]int32
func (c *TypeConverter) convertMapType(field *protogen.Field) string {
	mapKey := field.Message.Fields[0]   // Key field
	mapValue := field.Message.Fields[1] // Value field

	keyType, _ := c.convertType(mapKey)
	valueType, _ := c.convertType(mapValue)

	return fmt.Sprintf("map[%s]%s", keyType, valueType)
}

// GenerateTypeDefinition generates Go-Zero .api type definition string
func (c *TypeConverter) GenerateTypeDefinition(msg *model.Message) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("    %s {\n", msg.Name))

	for _, field := range msg.Fields {
		b.WriteString(c.generateFieldLine(field))
	}

	b.WriteString("    }\n")

	return b.String()
}

// generateFieldLine generates a single field line in .api format
func (c *TypeConverter) generateFieldLine(field *model.Field) string {
	fieldType := c.getGoZeroType(field)

	// Generate tags
	tags := fmt.Sprintf("`json:\"%s", field.JSONTag)
	if field.Optional {
		tags += ",optional"
	}
	if field.JSONTag == "-" {
		tags = "`json:\"-\"`"
	} else {
		tags += "\"`"
	}

	return fmt.Sprintf("        %s %s %s\n",
		field.Name,
		fieldType,
		tags)
}

// getGoZeroType returns the final Go-Zero type string for a field
func (c *TypeConverter) getGoZeroType(field *model.Field) string {
	baseType := field.Type

	// Handle repeated fields (arrays)
	if field.Repeated {
		return "[]" + baseType
	}

	// Handle optional fields (pointers)
	// Note: Go-Zero typically uses pointers for optional fields
	if field.Optional {
		// Don't add pointer for slice or map types
		if strings.HasPrefix(baseType, "[]") || strings.HasPrefix(baseType, "map[") {
			return baseType
		}
		return "*" + baseType
	}

	return baseType
}

// toGoFieldName converts snake_case to PascalCase for Go field names
func (c *TypeConverter) toGoFieldName(snakeCase string) string {
	if snakeCase == "" {
		return ""
	}

	parts := strings.Split(snakeCase, "_")
	result := ""

	for _, part := range parts {
		if len(part) > 0 {
			result += strings.ToUpper(part[:1]) + part[1:]
		}
	}

	return result
}

// GetAllConvertedTypes returns all converted message types
func (c *TypeConverter) GetAllConvertedTypes() []*model.Message {
	result := make([]*model.Message, 0, len(c.convertedTypes))
	for _, msg := range c.convertedTypes {
		result = append(result, msg)
	}
	return result
}

// Reset clears the converted types cache
func (c *TypeConverter) Reset() {
	c.convertedTypes = make(map[string]*model.Message)
}

// ConvertAllMessages converts all messages in a file
// This is useful for batch conversion
func (c *TypeConverter) ConvertAllMessages(file *protogen.File) []*model.Message {
	var messages []*model.Message

	for _, msg := range file.Messages {
		converted := c.ConvertMessage(msg)
		if converted != nil {
			messages = append(messages, converted)
		}

		// Recursively convert nested messages
		for _, nested := range msg.Messages {
			nestedConverted := c.ConvertMessage(nested)
			if nestedConverted != nil {
				messages = append(messages, nestedConverted)
			}
		}
	}

	return messages
}

// IsScalarType checks if a Proto type is a scalar type
func (c *TypeConverter) IsScalarType(kind protoreflect.Kind) bool {
	switch kind {
	case protoreflect.BoolKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind,
		protoreflect.FloatKind, protoreflect.DoubleKind,
		protoreflect.StringKind, protoreflect.BytesKind:
		return true
	default:
		return false
	}
}
