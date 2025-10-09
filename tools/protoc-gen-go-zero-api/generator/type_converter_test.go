package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// TestNewTypeConverter tests TypeConverter initialization
func TestNewTypeConverter(t *testing.T) {
	converter := NewTypeConverter()
	require.NotNil(t, converter)
	assert.NotNil(t, converter.convertedTypes)
	assert.Empty(t, converter.convertedTypes)
}

// TestToGoFieldName tests snake_case to PascalCase conversion
func TestToGoFieldName(t *testing.T) {
	converter := NewTypeConverter()

	tests := []struct {
		name      string
		input     string
		expected  string
	}{
		{
			name:     "simple snake_case",
			input:    "user_name",
			expected: "UserName",
		},
		{
			name:     "multiple underscores",
			input:    "user_full_name",
			expected: "UserFullName",
		},
		{
			name:     "single word",
			input:    "username",
			expected: "Username",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single letter",
			input:    "a",
			expected: "A",
		},
		{
			name:     "with numbers",
			input:    "user_id_123",
			expected: "UserId123",
		},
		{
			name:     "consecutive underscores",
			input:    "user__name",
			expected: "UserName",
		},
		{
			name:     "leading underscore",
			input:    "_user_name",
			expected: "UserName",
		},
		{
			name:     "trailing underscore",
			input:    "user_name_",
			expected: "UserName",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.toGoFieldName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsScalarType tests scalar type detection
func TestIsScalarType(t *testing.T) {
	converter := NewTypeConverter()

	tests := []struct {
		name     string
		kind     protoreflect.Kind
		expected bool
	}{
		{"bool", protoreflect.BoolKind, true},
		{"int32", protoreflect.Int32Kind, true},
		{"int64", protoreflect.Int64Kind, true},
		{"uint32", protoreflect.Uint32Kind, true},
		{"uint64", protoreflect.Uint64Kind, true},
		{"sint32", protoreflect.Sint32Kind, true},
		{"sint64", protoreflect.Sint64Kind, true},
		{"fixed32", protoreflect.Fixed32Kind, true},
		{"fixed64", protoreflect.Fixed64Kind, true},
		{"sfixed32", protoreflect.Sfixed32Kind, true},
		{"sfixed64", protoreflect.Sfixed64Kind, true},
		{"float", protoreflect.FloatKind, true},
		{"double", protoreflect.DoubleKind, true},
		{"string", protoreflect.StringKind, true},
		{"bytes", protoreflect.BytesKind, true},
		{"enum", protoreflect.EnumKind, false},
		{"message", protoreflect.MessageKind, false},
		{"group", protoreflect.GroupKind, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.IsScalarType(tt.kind)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestGetGoZeroType tests Go-Zero type generation for fields
func TestGetGoZeroType(t *testing.T) {
	converter := NewTypeConverter()

	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{
			name: "basic string",
			field: &model.Field{
				Type:     "string",
				Optional: false,
				Repeated: false,
			},
			expected: "string",
		},
		{
			name: "optional string",
			field: &model.Field{
				Type:     "string",
				Optional: true,
				Repeated: false,
			},
			expected: "*string",
		},
		{
			name: "repeated string",
			field: &model.Field{
				Type:     "string",
				Optional: false,
				Repeated: true,
			},
			expected: "[]string",
		},
		{
			name: "repeated int32",
			field: &model.Field{
				Type:     "int32",
				Optional: false,
				Repeated: true,
			},
			expected: "[]int32",
		},
		{
			name: "optional int64",
			field: &model.Field{
				Type:     "int64",
				Optional: true,
				Repeated: false,
			},
			expected: "*int64",
		},
		{
			name: "map type",
			field: &model.Field{
				Type:     "map[string]int32",
				Optional: false,
				Repeated: false,
			},
			expected: "map[string]int32",
		},
		{
			name: "optional map type (should not add pointer)",
			field: &model.Field{
				Type:     "map[string]int32",
				Optional: true,
				Repeated: false,
			},
			expected: "map[string]int32",
		},
		{
			name: "optional slice type (should not add pointer)",
			field: &model.Field{
				Type:     "[]string",
				Optional: true,
				Repeated: false,
			},
			expected: "[]string",
		},
		{
			name: "custom message type",
			field: &model.Field{
				Type:     "User",
				Optional: false,
				Repeated: false,
			},
			expected: "User",
		},
		{
			name: "optional custom message type",
			field: &model.Field{
				Type:     "User",
				Optional: true,
				Repeated: false,
			},
			expected: "*User",
		},
		{
			name: "repeated custom message type",
			field: &model.Field{
				Type:     "User",
				Optional: false,
				Repeated: true,
			},
			expected: "[]User",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.getGoZeroType(tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestGenerateFieldLine tests field line generation for .api format
func TestGenerateFieldLine(t *testing.T) {
	converter := NewTypeConverter()

	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{
			name: "basic field",
			field: &model.Field{
				Name:     "UserName",
				Type:     "string",
				JSONTag:  "userName",
				Optional: false,
				Repeated: false,
			},
			expected: "        UserName string `json:\"userName\"`\n",
		},
		{
			name: "optional field",
			field: &model.Field{
				Name:     "Email",
				Type:     "string",
				JSONTag:  "email",
				Optional: true,
				Repeated: false,
			},
			expected: "        Email *string `json:\"email,optional\"`\n",
		},
		{
			name: "repeated field",
			field: &model.Field{
				Name:     "Tags",
				Type:     "string",
				JSONTag:  "tags",
				Optional: false,
				Repeated: true,
			},
			expected: "        Tags []string `json:\"tags\"`\n",
		},
		{
			name: "ignored field",
			field: &model.Field{
				Name:     "Internal",
				Type:     "string",
				JSONTag:  "-",
				Optional: false,
				Repeated: false,
			},
			expected: "        Internal string `json:\"-\"`\n",
		},
		{
			name: "int32 field",
			field: &model.Field{
				Name:     "Age",
				Type:     "int32",
				JSONTag:  "age",
				Optional: false,
				Repeated: false,
			},
			expected: "        Age int32 `json:\"age\"`\n",
		},
		{
			name: "map field",
			field: &model.Field{
				Name:     "Metadata",
				Type:     "map[string]string",
				JSONTag:  "metadata",
				Optional: false,
				Repeated: false,
			},
			expected: "        Metadata map[string]string `json:\"metadata\"`\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.generateFieldLine(tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestGenerateTypeDefinition tests complete type definition generation
func TestGenerateTypeDefinition(t *testing.T) {
	converter := NewTypeConverter()

	tests := []struct {
		name     string
		message  *model.Message
		expected string
	}{
		{
			name: "simple message",
			message: &model.Message{
				Name: "User",
				Fields: []*model.Field{
					{
						Name:     "Id",
						Type:     "int64",
						JSONTag:  "id",
						Optional: false,
						Repeated: false,
					},
					{
						Name:     "Name",
						Type:     "string",
						JSONTag:  "name",
						Optional: false,
						Repeated: false,
					},
				},
			},
			expected: "    User {\n" +
				"        Id int64 `json:\"id\"`\n" +
				"        Name string `json:\"name\"`\n" +
				"    }\n",
		},
		{
			name: "message with optional fields",
			message: &model.Message{
				Name: "UpdateUserReq",
				Fields: []*model.Field{
					{
						Name:     "Id",
						Type:     "int64",
						JSONTag:  "id",
						Optional: false,
						Repeated: false,
					},
					{
						Name:     "Name",
						Type:     "string",
						JSONTag:  "name",
						Optional: true,
						Repeated: false,
					},
					{
						Name:     "Email",
						Type:     "string",
						JSONTag:  "email",
						Optional: true,
						Repeated: false,
					},
				},
			},
			expected: "    UpdateUserReq {\n" +
				"        Id int64 `json:\"id\"`\n" +
				"        Name *string `json:\"name,optional\"`\n" +
				"        Email *string `json:\"email,optional\"`\n" +
				"    }\n",
		},
		{
			name: "message with repeated fields",
			message: &model.Message{
				Name: "UserList",
				Fields: []*model.Field{
					{
						Name:     "Users",
						Type:     "User",
						JSONTag:  "users",
						Optional: false,
						Repeated: true,
					},
					{
						Name:     "Total",
						Type:     "int64",
						JSONTag:  "total",
						Optional: false,
						Repeated: false,
					},
				},
			},
			expected: "    UserList {\n" +
				"        Users []User `json:\"users\"`\n" +
				"        Total int64 `json:\"total\"`\n" +
				"    }\n",
		},
		{
			name: "empty message",
			message: &model.Message{
				Name:   "EmptyMessage",
				Fields: []*model.Field{},
			},
			expected: "    EmptyMessage {\n" +
				"    }\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.GenerateTypeDefinition(tt.message)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestReset tests the Reset functionality
func TestReset(t *testing.T) {
	converter := NewTypeConverter()

	// Add some converted types
	converter.convertedTypes["test.User"] = &model.Message{
		Name:   "User",
		Fields: []*model.Field{},
	}
	converter.convertedTypes["test.Role"] = &model.Message{
		Name:   "Role",
		Fields: []*model.Field{},
	}

	assert.Len(t, converter.convertedTypes, 2)

	// Reset
	converter.Reset()

	assert.Empty(t, converter.convertedTypes)
	assert.NotNil(t, converter.convertedTypes)
}

// TestGetAllConvertedTypes tests retrieving all converted types
func TestGetAllConvertedTypes(t *testing.T) {
	converter := NewTypeConverter()

	// Add some converted types
	user := &model.Message{
		Name:   "User",
		Fields: []*model.Field{},
	}
	role := &model.Message{
		Name:   "Role",
		Fields: []*model.Field{},
	}

	converter.convertedTypes["test.User"] = user
	converter.convertedTypes["test.Role"] = role

	result := converter.GetAllConvertedTypes()

	assert.Len(t, result, 2)
	// Result order is not guaranteed due to map iteration
	names := make(map[string]bool)
	for _, msg := range result {
		names[msg.Name] = true
	}
	assert.True(t, names["User"])
	assert.True(t, names["Role"])
}

// Helper function to create a minimal proto plugin request for testing
func createTestPluginRequest(protoContent string) (*pluginpb.CodeGeneratorRequest, error) {
	// Create a file descriptor proto
	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/test/testpb"),
		},
	}

	// Add a simple message for testing
	messageDesc := &descriptorpb.DescriptorProto{
		Name: proto.String("TestMessage"),
	}

	// Add fields to the message
	messageDesc.Field = []*descriptorpb.FieldDescriptorProto{
		{
			Name:   proto.String("id"),
			Number: proto.Int32(1),
			Type:   descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
			Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		},
		{
			Name:   proto.String("name"),
			Number: proto.Int32(2),
			Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		},
		{
			Name:   proto.String("tags"),
			Number: proto.Int32(3),
			Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		},
	}

	fileDesc.MessageType = []*descriptorpb.DescriptorProto{messageDesc}

	// Create the code generator request
	request := &pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{fileDesc},
		FileToGenerate: []string{"test.proto"},
	}

	return request, nil
}

// TestConvertMessage tests message conversion with mock protogen.Message
func TestConvertMessage(t *testing.T) {
	// Create a test plugin request
	request, err := createTestPluginRequest("")
	require.NoError(t, err)

	// Create protogen plugin
	plugin, err := protogen.Options{}.New(request)
	require.NoError(t, err)

	converter := NewTypeConverter()

	// Test converting the first file's first message
	if len(plugin.Files) > 0 && len(plugin.Files[0].Messages) > 0 {
		msg := plugin.Files[0].Messages[0]
		result := converter.ConvertMessage(msg)

		require.NotNil(t, result)
		assert.Equal(t, "TestMessage", result.Name)
		assert.Len(t, result.Fields, 3)

		// Check first field (id)
		assert.Equal(t, "Id", result.Fields[0].Name)
		assert.Equal(t, "int64", result.Fields[0].Type)
		assert.False(t, result.Fields[0].Repeated)

		// Check second field (name)
		assert.Equal(t, "Name", result.Fields[1].Name)
		assert.Equal(t, "string", result.Fields[1].Type)
		assert.False(t, result.Fields[1].Repeated)

		// Check third field (tags - repeated)
		assert.Equal(t, "Tags", result.Fields[2].Name)
		assert.Equal(t, "string", result.Fields[2].Type)
		assert.True(t, result.Fields[2].Repeated)

		// Test duplicate conversion returns same instance
		result2 := converter.ConvertMessage(msg)
		assert.Equal(t, result, result2)
	}
}

// TestConvertMessage_CachingBehavior tests that duplicate conversions are cached
func TestConvertMessage_CachingBehavior(t *testing.T) {
	request, err := createTestPluginRequest("")
	require.NoError(t, err)

	plugin, err := protogen.Options{}.New(request)
	require.NoError(t, err)

	converter := NewTypeConverter()

	if len(plugin.Files) > 0 && len(plugin.Files[0].Messages) > 0 {
		msg := plugin.Files[0].Messages[0]

		// First conversion
		result1 := converter.ConvertMessage(msg)
		require.NotNil(t, result1)

		// Second conversion should return cached result
		result2 := converter.ConvertMessage(msg)
		require.NotNil(t, result2)

		// Should be the exact same pointer
		assert.True(t, result1 == result2, "Expected same pointer for cached conversion")

		// Verify it's in the cache
		fullName := string(msg.Desc.FullName())
		cached, ok := converter.convertedTypes[fullName]
		assert.True(t, ok)
		assert.True(t, cached == result1)
	}
}

// TestConvertType tests type conversion for various proto types
func TestConvertType(t *testing.T) {
	request, err := createTestPluginRequest("")
	require.NoError(t, err)

	plugin, err := protogen.Options{}.New(request)
	require.NoError(t, err)

	converter := NewTypeConverter()

	if len(plugin.Files) > 0 && len(plugin.Files[0].Messages) > 0 {
		msg := plugin.Files[0].Messages[0]

		for _, field := range msg.Fields {
			goType, protoType := converter.convertType(field)

			switch string(field.Desc.Name()) {
			case "id":
				assert.Equal(t, "int64", goType)
				assert.Equal(t, "int64", protoType)
			case "name":
				assert.Equal(t, "string", goType)
				assert.Equal(t, "string", protoType)
			case "tags":
				assert.Equal(t, "string", goType)
				assert.Equal(t, "string", protoType)
			}
		}
	}
}

// TestConvertType_AllTypes tests all proto type conversions comprehensively
func TestConvertType_AllTypes(t *testing.T) {
	// Create a comprehensive test proto with all types
	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test_all_types.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/test/testpb"),
		},
	}

	// Create enum for testing
	enumDesc := &descriptorpb.EnumDescriptorProto{
		Name: proto.String("Status"),
		Value: []*descriptorpb.EnumValueDescriptorProto{
			{Name: proto.String("UNKNOWN"), Number: proto.Int32(0)},
			{Name: proto.String("ACTIVE"), Number: proto.Int32(1)},
		},
	}

	// Create a nested message for testing
	nestedMessageDesc := &descriptorpb.DescriptorProto{
		Name: proto.String("NestedMessage"),
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("value"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
		},
	}

	// Create main message with all field types
	messageDesc := &descriptorpb.DescriptorProto{
		Name: proto.String("AllTypesMessage"),
		Field: []*descriptorpb.FieldDescriptorProto{
			// Bool
			{
				Name:   proto.String("bool_field"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// Int32 variants
			{
				Name:   proto.String("int32_field"),
				Number: proto.Int32(2),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("sint32_field"),
				Number: proto.Int32(3),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_SINT32.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("sfixed32_field"),
				Number: proto.Int32(4),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_SFIXED32.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// Int64 variants
			{
				Name:   proto.String("int64_field"),
				Number: proto.Int32(5),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("sint64_field"),
				Number: proto.Int32(6),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_SINT64.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("sfixed64_field"),
				Number: proto.Int32(7),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_SFIXED64.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// Uint32 variants
			{
				Name:   proto.String("uint32_field"),
				Number: proto.Int32(8),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("fixed32_field"),
				Number: proto.Int32(9),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_FIXED32.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// Uint64 variants
			{
				Name:   proto.String("uint64_field"),
				Number: proto.Int32(10),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_UINT64.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("fixed64_field"),
				Number: proto.Int32(11),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_FIXED64.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// Float types
			{
				Name:   proto.String("float_field"),
				Number: proto.Int32(12),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_FLOAT.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("double_field"),
				Number: proto.Int32(13),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// String and bytes
			{
				Name:   proto.String("string_field"),
				Number: proto.Int32(14),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("bytes_field"),
				Number: proto.Int32(15),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_BYTES.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// Enum
			{
				Name:     proto.String("enum_field"),
				Number:   proto.Int32(16),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
				TypeName: proto.String(".test.Status"),
				Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			// Message (use nested type reference)
			{
				Name:     proto.String("message_field"),
				Number:   proto.Int32(17),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
				TypeName: proto.String(".test.AllTypesMessage.NestedMessage"),
				Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
		},
		NestedType: []*descriptorpb.DescriptorProto{nestedMessageDesc},
		EnumType:   []*descriptorpb.EnumDescriptorProto{enumDesc},
	}

	fileDesc.MessageType = []*descriptorpb.DescriptorProto{messageDesc}
	fileDesc.EnumType = []*descriptorpb.EnumDescriptorProto{enumDesc}

	request := &pluginpb.CodeGeneratorRequest{
		ProtoFile:      []*descriptorpb.FileDescriptorProto{fileDesc},
		FileToGenerate: []string{"test_all_types.proto"},
	}

	plugin, err := protogen.Options{}.New(request)
	require.NoError(t, err)

	converter := NewTypeConverter()

	expectedTypes := map[string]struct {
		goType    string
		protoType string
	}{
		"bool_field":      {"bool", "bool"},
		"int32_field":     {"int32", "int32"},
		"sint32_field":    {"int32", "int32"},
		"sfixed32_field":  {"int32", "int32"},
		"int64_field":     {"int64", "int64"},
		"sint64_field":    {"int64", "int64"},
		"sfixed64_field":  {"int64", "int64"},
		"uint32_field":    {"uint32", "uint32"},
		"fixed32_field":   {"uint32", "uint32"},
		"uint64_field":    {"uint64", "uint64"},
		"fixed64_field":   {"uint64", "uint64"},
		"float_field":     {"float32", "float"},
		"double_field":    {"float64", "double"},
		"string_field":    {"string", "string"},
		"bytes_field":     {"[]byte", "bytes"},
		"enum_field":      {"string", "enum"},
		"message_field":   {"NestedMessage", "message"},
	}

	if len(plugin.Files) > 0 && len(plugin.Files[0].Messages) > 0 {
		msg := plugin.Files[0].Messages[0]

		for _, field := range msg.Fields {
			fieldName := string(field.Desc.Name())
			expected, ok := expectedTypes[fieldName]
			require.True(t, ok, "Unexpected field: %s", fieldName)

			goType, protoType := converter.convertType(field)

			assert.Equal(t, expected.goType, goType, "Field %s: go type mismatch", fieldName)
			assert.Equal(t, expected.protoType, protoType, "Field %s: proto type mismatch", fieldName)
		}
	}
}

// TestConvertMapType tests map type conversion
func TestConvertMapType(t *testing.T) {
	// Create a test proto with a map field
	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test_map.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/test/testpb"),
		},
	}

	// Create map entry message (this is how protobuf represents maps)
	mapEntryDesc := &descriptorpb.DescriptorProto{
		Name: proto.String("MetadataEntry"),
		Options: &descriptorpb.MessageOptions{
			MapEntry: proto.Bool(true),
		},
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("key"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
			{
				Name:   proto.String("value"),
				Number: proto.Int32(2),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
		},
	}

	// Create main message with map field
	messageDesc := &descriptorpb.DescriptorProto{
		Name: proto.String("TestMapMessage"),
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:     proto.String("metadata"),
				Number:   proto.Int32(1),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
				TypeName: proto.String(".test.TestMapMessage.MetadataEntry"),
				Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
			},
		},
		NestedType: []*descriptorpb.DescriptorProto{mapEntryDesc},
	}

	fileDesc.MessageType = []*descriptorpb.DescriptorProto{messageDesc}

	request := &pluginpb.CodeGeneratorRequest{
		ProtoFile:      []*descriptorpb.FileDescriptorProto{fileDesc},
		FileToGenerate: []string{"test_map.proto"},
	}

	plugin, err := protogen.Options{}.New(request)
	require.NoError(t, err)

	converter := NewTypeConverter()

	if len(plugin.Files) > 0 && len(plugin.Files[0].Messages) > 0 {
		msg := plugin.Files[0].Messages[0]

		if len(msg.Fields) > 0 {
			field := msg.Fields[0]

			// Check if it's recognized as a map
			if field.Desc.IsMap() {
				mapType := converter.convertMapType(field)
				assert.Equal(t, "map[string]int32", mapType)
			}
		}
	}
}

// TestConvertAllMessages tests batch conversion of all messages in a file
func TestConvertAllMessages(t *testing.T) {
	// Create a test proto with multiple messages
	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test_multi.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/test/testpb"),
		},
	}

	// First message
	msg1 := &descriptorpb.DescriptorProto{
		Name: proto.String("User"),
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("id"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
		},
	}

	// Second message
	msg2 := &descriptorpb.DescriptorProto{
		Name: proto.String("Role"),
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("name"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
		},
	}

	// Third message with nested message
	nestedMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("Address"),
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("street"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
		},
	}

	msg3 := &descriptorpb.DescriptorProto{
		Name: proto.String("Company"),
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("name"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			},
		},
		NestedType: []*descriptorpb.DescriptorProto{nestedMsg},
	}

	fileDesc.MessageType = []*descriptorpb.DescriptorProto{msg1, msg2, msg3}

	request := &pluginpb.CodeGeneratorRequest{
		ProtoFile:      []*descriptorpb.FileDescriptorProto{fileDesc},
		FileToGenerate: []string{"test_multi.proto"},
	}

	plugin, err := protogen.Options{}.New(request)
	require.NoError(t, err)

	converter := NewTypeConverter()

	if len(plugin.Files) > 0 {
		file := plugin.Files[0]
		messages := converter.ConvertAllMessages(file)

		// Should convert all messages including nested ones
		assert.GreaterOrEqual(t, len(messages), 3)

		// Check message names
		names := make(map[string]bool)
		for _, msg := range messages {
			names[msg.Name] = true
		}

		assert.True(t, names["User"])
		assert.True(t, names["Role"])
		assert.True(t, names["Company"])
		assert.True(t, names["Address"]) // Nested message
	}
}

// TestConvertField tests field conversion
func TestConvertField(t *testing.T) {
	request, err := createTestPluginRequest("")
	require.NoError(t, err)

	plugin, err := protogen.Options{}.New(request)
	require.NoError(t, err)

	converter := NewTypeConverter()

	if len(plugin.Files) > 0 && len(plugin.Files[0].Messages) > 0 {
		msg := plugin.Files[0].Messages[0]

		for _, field := range msg.Fields {
			convertedField := converter.convertField(field)
			require.NotNil(t, convertedField)

			// Field should have a name
			assert.NotEmpty(t, convertedField.Name)

			// Field should have a type
			assert.NotEmpty(t, convertedField.Type)

			// Field should have a JSON tag
			assert.NotEmpty(t, convertedField.JSONTag)

			// Check specific field properties
			switch string(field.Desc.Name()) {
			case "tags":
				assert.True(t, convertedField.Repeated)
			default:
				assert.False(t, convertedField.Repeated)
			}
		}
	}
}

// Benchmark tests for performance
func BenchmarkToGoFieldName(b *testing.B) {
	converter := NewTypeConverter()
	for i := 0; i < b.N; i++ {
		converter.toGoFieldName("user_full_name")
	}
}

func BenchmarkConvertMessage(b *testing.B) {
	request, _ := createTestPluginRequest("")
	plugin, _ := protogen.Options{}.New(request)

	if len(plugin.Files) > 0 && len(plugin.Files[0].Messages) > 0 {
		msg := plugin.Files[0].Messages[0]

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			converter := NewTypeConverter()
			converter.ConvertMessage(msg)
		}
	}
}

func BenchmarkGenerateTypeDefinition(b *testing.B) {
	converter := NewTypeConverter()
	message := &model.Message{
		Name: "User",
		Fields: []*model.Field{
			{Name: "Id", Type: "int64", JSONTag: "id"},
			{Name: "Name", Type: "string", JSONTag: "name"},
			{Name: "Email", Type: "string", JSONTag: "email", Optional: true},
			{Name: "Tags", Type: "string", JSONTag: "tags", Repeated: true},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		converter.GenerateTypeDefinition(message)
	}
}
