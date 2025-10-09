package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	go_zero "github.com/chimerakang/simple-admin-core/rpc/types/go_zero"
	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// TestGenerateAPI_EndToEnd tests complete .api file generation from Proto
func TestGenerateAPI_EndToEnd(t *testing.T) {
	t.Run("basic CRUD service generation", func(t *testing.T) {
		// Create test proto file descriptor with proper HTTP annotations
		req := createUserServiceProtoRequest(t)

		// Create protogen plugin
		plugin, err := protogen.Options{}.New(req)
		require.NoError(t, err)
		require.NotEmpty(t, plugin.Files)

		// Generate API file
		generator := NewGenerator(plugin.Files[0])
		content, err := generator.Generate()

		// Verify generation succeeds
		require.NoError(t, err)
		require.NotEmpty(t, content)

		// Log the generated content for manual inspection
		t.Logf("Generated .api content:\n%s", content)

		// Verify basic structure
		assert.Contains(t, content, "syntax = \"v1\"", "Should have syntax declaration")
		assert.Contains(t, content, "import \"../base.api\"", "Should have base import")
		assert.Contains(t, content, "type (", "Should have type section")
		assert.Contains(t, content, "service Core {", "Should have service declaration")

		// Verify at least some types are generated
		assert.Contains(t, content, "CreateUserReq", "Should have CreateUserReq type")
		assert.Contains(t, content, "GetUserReq", "Should have GetUserReq type")

		// Verify output is properly formatted (no trailing whitespace)
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			trimmed := strings.TrimRight(line, " \t")
			assert.Equal(t, trimmed, line, "Line %d should not have trailing whitespace", i+1)
		}
	})
}

// TestGenerateAPI_TypeConversion tests type definitions are correctly generated
func TestGenerateAPI_TypeConversion(t *testing.T) {
	t.Run("message to .api type conversion", func(t *testing.T) {
		req := createTypeConversionProtoRequest(t)

		plugin, err := protogen.Options{}.New(req)
		require.NoError(t, err)
		require.NotEmpty(t, plugin.Files)

		generator := NewGenerator(plugin.Files[0])
		content, err := generator.Generate()

		require.NoError(t, err)
		t.Logf("Generated types:\n%s", content)

		// Verify different field types
		assert.Contains(t, content, "string", "Should have string type")
		assert.Contains(t, content, "int64", "Should have int64 type")
		assert.Contains(t, content, "bool", "Should have bool type")

		// Verify repeated fields generate arrays
		assert.Contains(t, content, "[]", "Should have array type for repeated fields")

		// Verify optional fields generate pointers
		assert.Contains(t, content, "*", "Should have pointer for optional fields")
		assert.Contains(t, content, ",optional", "Should have optional tag")
	})
}

// TestTemplateGeneration tests the template generation component
func TestTemplateGeneration(t *testing.T) {
	t.Run("template generates valid .api structure", func(t *testing.T) {
		templateGen := NewTemplateGenerator()

		// Create test data
		data := &TemplateData{
			APIInfo: &go_zero.ApiInfo{
				Title:   "Test API",
				Desc:    "Test Description",
				Author:  "Test Author",
				Version: "v1.0",
			},
			Types: []*TypeDefinition{
				{
					Name: "TestReq",
					Definition: "    TestReq {\n" +
						"        Id int64 `json:\"id\"`\n" +
						"        Name string `json:\"name\"`\n" +
						"    }\n",
				},
				{
					Name: "TestResp",
					Definition: "    TestResp {\n" +
						"        Success bool `json:\"success\"`\n" +
						"    }\n",
				},
			},
			ServiceGroups: []*ServiceGroup{
				{
					ServerOptions: &model.ServerOptions{
						JWT:    "Auth",
						Group:  "test",
						Prefix: "/api/v1",
					},
					Methods: []*model.Method{
						{
							Name:         "TestMethod",
							RequestType:  "TestReq",
							ResponseType: "TestResp",
							HTTPRule: &model.HTTPRule{
								Method: "post",
								Path:   "/test",
								Body:   "*",
							},
							Options: &model.ServerOptions{
								JWT:    "Auth",
								Group:  "test",
								Prefix: "/api/v1",
							},
						},
					},
				},
			},
		}

		// Generate content
		content, err := templateGen.Generate(data)
		require.NoError(t, err)
		require.NotEmpty(t, content)

		t.Logf("Template output:\n%s", content)

		// Verify structure
		assert.Contains(t, content, "syntax = \"v1\"")
		assert.Contains(t, content, "info(")
		assert.Contains(t, content, "title: \"Test API\"")
		assert.Contains(t, content, "type (")
		assert.Contains(t, content, "TestReq {")
		assert.Contains(t, content, "TestResp {")
		assert.Contains(t, content, "@server(")
		assert.Contains(t, content, "jwt: Auth")
		assert.Contains(t, content, "group: test")
		assert.Contains(t, content, "service Core {")
		assert.Contains(t, content, "@handler testMethod")
		assert.Contains(t, content, "post /test (TestReq) returns (TestResp)")
	})
}

// TestServiceGrouping tests method grouping by @server configuration
func TestServiceGrouping(t *testing.T) {
	t.Run("methods grouped by server options", func(t *testing.T) {
		grouper := NewServiceGrouper()

		// Create test methods with different options
		methods := []*model.Method{
			{
				Name: "PublicMethod",
				Options: &model.ServerOptions{
					Group: "public",
				},
			},
			{
				Name: "ProtectedMethod1",
				Options: &model.ServerOptions{
					JWT:   "Auth",
					Group: "user",
				},
			},
			{
				Name: "ProtectedMethod2",
				Options: &model.ServerOptions{
					JWT:   "Auth",
					Group: "user",
				},
			},
			{
				Name: "AdminMethod",
				Options: &model.ServerOptions{
					JWT:        "Auth",
					Group:      "admin",
					Middleware: []string{"Authority"},
				},
			},
		}

		// Group methods
		groups := grouper.GroupMethods(methods)

		// Verify grouping
		assert.NotEmpty(t, groups)
		t.Logf("Number of groups: %d", len(groups))

		// Methods with same options should be in same group
		userGroup := findGroup(groups, "user")
		if userGroup != nil {
			assert.Len(t, userGroup.Methods, 2, "User group should have 2 methods")
		}

		// Admin group should be separate
		adminGroup := findGroup(groups, "admin")
		if adminGroup != nil {
			assert.Len(t, adminGroup.Methods, 1, "Admin group should have 1 method")
			assert.NotEmpty(t, adminGroup.ServerOptions.Middleware)
		}
	})
}

// TestValidation tests template data validation
func TestValidation(t *testing.T) {
	t.Run("duplicate type names", func(t *testing.T) {
		templateGen := NewTemplateGenerator()

		data := &TemplateData{
			Types: []*TypeDefinition{
				{Name: "User", Definition: "User {}"},
				{Name: "User", Definition: "User {}"},
			},
		}

		err := templateGen.ValidateData(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "duplicate type name")
	})

	t.Run("duplicate handler names", func(t *testing.T) {
		templateGen := NewTemplateGenerator()

		data := &TemplateData{
			ServiceGroups: []*ServiceGroup{
				{
					ServerOptions: &model.ServerOptions{Group: "user"},
					Methods: []*model.Method{
						{Name: "GetUser", Options: &model.ServerOptions{Group: "user"}},
						{Name: "GetUser", Options: &model.ServerOptions{Group: "user"}},
					},
				},
			},
		}

		err := templateGen.ValidateData(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "duplicate handler name")
	})
}

// TestGenerateWithDefaults tests generation with default values
func TestGenerateWithDefaults(t *testing.T) {
	t.Run("generation with missing API info", func(t *testing.T) {
		templateGen := NewTemplateGenerator()

		data := &TemplateData{
			Types: []*TypeDefinition{
				{
					Name:       "TestReq",
					Definition: "    TestReq {\n        Id int64 `json:\"id\"`\n    }\n",
				},
			},
		}

		// Generate with defaults
		content, err := templateGen.GenerateWithDefaults(data)
		require.NoError(t, err)
		require.NotEmpty(t, content)

		// Should have default API info
		assert.Contains(t, content, "info(")
		assert.Contains(t, content, "Auto-generated API from Proto")
	})
}

// TestOutputFormat tests the output format and whitespace handling
func TestOutputFormat(t *testing.T) {
	t.Run("output has proper formatting", func(t *testing.T) {
		templateGen := NewTemplateGenerator()

		data := &TemplateData{
			APIInfo: &go_zero.ApiInfo{
				Title:   "Test",
				Desc:    "Test",
				Author:  "Test",
				Version: "v1.0",
			},
		}

		content, err := templateGen.Generate(data)
		require.NoError(t, err)

		// Check no trailing whitespace
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			assert.Equal(t, strings.TrimRight(line, " \t"), line,
				"Line %d has trailing whitespace: %q", i+1, line)
		}

		// Check no excessive blank lines (max 2 consecutive)
		blankCount := 0
		for i, line := range lines {
			if strings.TrimSpace(line) == "" {
				blankCount++
				assert.LessOrEqual(t, blankCount, 2,
					"More than 2 consecutive blank lines at line %d", i+1)
			} else {
				blankCount = 0
			}
		}

		// Check content ends with single newline
		assert.True(t, strings.HasSuffix(content, "\n"), "Content should end with newline")
		assert.False(t, strings.HasSuffix(content, "\n\n"), "Content should not have trailing blank lines")
	})
}

// Helper function to create a user service proto request with HTTP annotations
func createUserServiceProtoRequest(t *testing.T) *pluginpb.CodeGeneratorRequest {
	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("user.proto"),
		Package: proto.String("user"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/test/user"),
		},
	}

	// Add file-level API info options
	proto.SetExtension(fileDesc.Options, go_zero.E_ApiInfo, &go_zero.ApiInfo{
		Title:   "User Service API",
		Desc:    "Simple user management service",
		Author:  "Test Author",
		Email:   "test@example.com",
		Version: "v1.0",
	})

	// Add messages
	fileDesc.MessageType = []*descriptorpb.DescriptorProto{
		{
			Name: proto.String("CreateUserReq"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:   proto.String("name"),
					Number: proto.Int32(1),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
				{
					Name:   proto.String("email"),
					Number: proto.Int32(2),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
			},
		},
		{
			Name: proto.String("CreateUserResp"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:   proto.String("id"),
					Number: proto.Int32(1),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
			},
		},
		{
			Name: proto.String("GetUserReq"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:   proto.String("id"),
					Number: proto.Int32(1),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
			},
		},
		{
			Name: proto.String("GetUserResp"),
			Field: []*descriptorpb.FieldDescriptorProto{
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
			},
		},
	}

	// Add service with HTTP annotations
	serviceOpts := &descriptorpb.ServiceOptions{}
	proto.SetExtension(serviceOpts, go_zero.E_Jwt, "Auth")
	proto.SetExtension(serviceOpts, go_zero.E_Group, "user")
	proto.SetExtension(serviceOpts, go_zero.E_Prefix, "/api/v1")

	createUserOpts := &descriptorpb.MethodOptions{}
	proto.SetExtension(createUserOpts, annotations.E_Http, &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Post{
			Post: "/user",
		},
		Body: "*",
	})

	getUserOpts := &descriptorpb.MethodOptions{}
	proto.SetExtension(getUserOpts, annotations.E_Http, &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Get{
			Get: "/user/{id}",
		},
	})

	service := &descriptorpb.ServiceDescriptorProto{
		Name:    proto.String("UserService"),
		Options: serviceOpts,
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:       proto.String("CreateUser"),
				InputType:  proto.String(".user.CreateUserReq"),
				OutputType: proto.String(".user.CreateUserResp"),
				Options:    createUserOpts,
			},
			{
				Name:       proto.String("GetUser"),
				InputType:  proto.String(".user.GetUserReq"),
				OutputType: proto.String(".user.GetUserResp"),
				Options:    getUserOpts,
			},
		},
	}

	fileDesc.Service = []*descriptorpb.ServiceDescriptorProto{service}

	return &pluginpb.CodeGeneratorRequest{
		ProtoFile:      []*descriptorpb.FileDescriptorProto{fileDesc},
		FileToGenerate: []string{"user.proto"},
	}
}

// Helper function to create proto request for type conversion testing
func createTypeConversionProtoRequest(t *testing.T) *pluginpb.CodeGeneratorRequest {
	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("types.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/test/types"),
		},
	}

	// Add file-level options
	proto.SetExtension(fileDesc.Options, go_zero.E_ApiInfo, &go_zero.ApiInfo{
		Title:   "Type Test API",
		Desc:    "Test type conversions",
		Author:  "Test",
		Version: "v1.0",
	})

	// Message with various field types
	fileDesc.MessageType = []*descriptorpb.DescriptorProto{
		{
			Name: proto.String("TypeTestReq"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:   proto.String("string_field"),
					Number: proto.Int32(1),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
				{
					Name:   proto.String("int64_field"),
					Number: proto.Int32(2),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
				{
					Name:   proto.String("bool_field"),
					Number: proto.Int32(3),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
				{
					Name:            proto.String("optional_field"),
					Number:          proto.Int32(4),
					Type:            descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					Label:           descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Proto3Optional:  proto.Bool(true),
				},
				{
					Name:   proto.String("repeated_field"),
					Number: proto.Int32(5),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
				},
			},
		},
		{
			Name: proto.String("TypeTestResp"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:   proto.String("success"),
					Number: proto.Int32(1),
					Type:   descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum(),
					Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				},
			},
		},
	}

	// Add service
	serviceOpts := &descriptorpb.ServiceOptions{}
	proto.SetExtension(serviceOpts, go_zero.E_Group, "test")

	methodOpts := &descriptorpb.MethodOptions{}
	proto.SetExtension(methodOpts, annotations.E_Http, &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Post{
			Post: "/test",
		},
		Body: "*",
	})

	service := &descriptorpb.ServiceDescriptorProto{
		Name:    proto.String("TestService"),
		Options: serviceOpts,
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:       proto.String("TestTypes"),
				InputType:  proto.String(".test.TypeTestReq"),
				OutputType: proto.String(".test.TypeTestResp"),
				Options:    methodOpts,
			},
		},
	}

	fileDesc.Service = []*descriptorpb.ServiceDescriptorProto{service}

	return &pluginpb.CodeGeneratorRequest{
		ProtoFile:      []*descriptorpb.FileDescriptorProto{fileDesc},
		FileToGenerate: []string{"types.proto"},
	}
}

// Helper function to find a service group by group name
func findGroup(groups []*ServiceGroup, groupName string) *ServiceGroup {
	for _, group := range groups {
		if group.ServerOptions.Group == groupName {
			return group
		}
	}
	return nil
}
