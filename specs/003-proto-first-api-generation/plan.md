# Technical Plan: Proto-First API Generation

**Related Spec**: [spec.md](./spec.md)
**Created**: 2025-10-08
**Status**: Draft
**Estimated Effort**: 60-80 hours

## Architecture Overview

### System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                   Developer Workflow                             │
│                                                                  │
│  1. Edit Proto file (rpc/desc/user.proto)                      │
│     - Add google.api.http annotations                           │
│     - Add go_zero.* custom options (jwt, middleware)           │
│                                                                  │
│  2. Run: make gen-proto-api                                     │
│     └─> protoc --go-zero-api_out=api/desc \                    │
│             --proto_path=. \                                     │
│             rpc/desc/**/*.proto                                 │
│                                                                  │
│  3. Auto-generated outputs:                                     │
│     ✅ api/desc/core/user.api (Go-Zero REST API definition)    │
│     ✅ Types with correct field mappings                        │
│     ✅ @server annotations (jwt, middleware, group)            │
│     ✅ Route handlers from RPC methods                          │
└─────────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│              protoc-gen-go-zero-api Plugin                       │
│                                                                  │
│  ┌──────────────────────────────────────────────────────┐      │
│  │  1. Proto File Parser                                 │      │
│  │     - Parse service definitions                       │      │
│  │     - Extract methods and messages                    │      │
│  │     - Read file-level metadata                        │      │
│  └──────────────────────────────────────────────────────┘      │
│                         │                                        │
│                         ▼                                        │
│  ┌──────────────────────────────────────────────────────┐      │
│  │  2. HTTP Annotation Parser                            │      │
│  │     - Extract google.api.http options                 │      │
│  │     - Parse HTTP method (GET/POST/PUT/DELETE)        │      │
│  │     - Parse path with parameters                      │      │
│  │     - Parse body mapping                              │      │
│  │     - Handle additional_bindings                      │      │
│  └──────────────────────────────────────────────────────┘      │
│                         │                                        │
│                         ▼                                        │
│  ┌──────────────────────────────────────────────────────┐      │
│  │  3. Go-Zero Options Parser                            │      │
│  │     - Extract go_zero.jwt option                      │      │
│  │     - Extract go_zero.middleware option               │      │
│  │     - Extract go_zero.group option                    │      │
│  │     - Extract go_zero.public (method-level)          │      │
│  │     - Extract go_zero.api_info (file-level)          │      │
│  └──────────────────────────────────────────────────────┘      │
│                         │                                        │
│                         ▼                                        │
│  ┌──────────────────────────────────────────────────────┐      │
│  │  4. Type Converter                                    │      │
│  │     - Proto types → Go-Zero types                     │      │
│  │     - Field name transformation (snake_case → json)  │      │
│  │     - Handle nested messages                          │      │
│  │     - Generate optional fields                        │      │
│  └──────────────────────────────────────────────────────┘      │
│                         │                                        │
│                         ▼                                        │
│  ┌──────────────────────────────────────────────────────┐      │
│  │  5. Service Grouper                                   │      │
│  │     - Group methods by @server config                 │      │
│  │     - Separate public vs protected endpoints          │      │
│  │     - Handle method-specific middleware               │      │
│  └──────────────────────────────────────────────────────┘      │
│                         │                                        │
│                         ▼                                        │
│  ┌──────────────────────────────────────────────────────┐      │
│  │  6. .api File Generator                               │      │
│  │     - Generate info() section                         │      │
│  │     - Generate type definitions                       │      │
│  │     - Generate @server blocks                         │      │
│  │     - Generate service definitions                    │      │
│  │     - Format output                                   │      │
│  └──────────────────────────────────────────────────────┘      │
│                         │                                        │
└─────────────────────────┼────────────────────────────────────────┘
                          │
                          ▼
                ┌─────────────────────┐
                │  api/desc/core/     │
                │    user.api         │
                │    role.api         │
                │    menu.api         │
                └─────────────────────┘
```

### Data Flow

```
Proto Definition
       │
       ├─> google.api.http ────┐
       ├─> go_zero.jwt ────────┤
       ├─> go_zero.middleware ─┤──> Plugin Parser
       ├─> go_zero.group ──────┤
       └─> go_zero.api_info ───┘
                │
                ▼
         Internal AST Model
                │
                ├─> Services []Service
                │      ├─> Name
                │      ├─> Methods []Method
                │      │      ├─> Name
                │      │      ├─> HTTPRule
                │      │      ├─> Request/Response
                │      │      └─> Options
                │      └─> Options (JWT, Middleware, Group)
                │
                └─> Messages []Message
                       ├─> Name
                       └─> Fields []Field
                │
                ▼
       .api File Generator
                │
                ├─> info() section
                ├─> type definitions
                ├─> @server blocks (grouped)
                └─> service methods
                │
                ▼
         Generated .api File
```

### Technology Stack

- **Language**: Go 1.21+
- **Proto Parsing**: `google.golang.org/protobuf/compiler/protogen` v1.31.0+
- **Proto Annotations**: `google.golang.org/genproto/googleapis/api/annotations`
- **Code Generation**: Text templates with `text/template`
- **Testing**: `github.com/stretchr/testify` for assertions
- **Build Tool**: Standard `go build` with Makefile integration

## Implementation Details

### Phase 1: Project Setup and Custom Proto Options (8-12 hours)

#### Task 1.1: Create Plugin Project Structure (2 hours)

**File**: `tools/protoc-gen-go-zero-api/`

**Directory Structure**:
```
tools/protoc-gen-go-zero-api/
├── main.go              # Plugin entry point
├── generator/
│   ├── generator.go     # Main generator logic
│   ├── parser.go        # Proto parsing
│   ├── http_parser.go   # HTTP annotation parser
│   ├── options_parser.go # Go-Zero options parser
│   ├── type_converter.go # Type conversion
│   ├── grouper.go       # Service grouping logic
│   └── template.go      # .api template
├── model/
│   ├── service.go       # Service model
│   ├── method.go        # Method model
│   ├── message.go       # Message model
│   └── options.go       # Options model
├── test/
│   ├── fixtures/        # Test Proto files
│   │   ├── user.proto
│   │   └── expected_user.api
│   └── generator_test.go
├── go.mod
└── go.sum
```

**Code Example** - `main.go`:
```go
package main

import (
    "flag"
    "fmt"

    "google.golang.org/protobuf/compiler/protogen"
    "google.golang.org/protobuf/types/pluginpb"
)

func main() {
    var flags flag.FlagSet

    protogen.Options{
        ParamFunc: flags.Set,
    }.Run(func(gen *protogen.Plugin) error {
        gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

        for _, f := range gen.Files {
            if !f.Generate {
                continue
            }
            generateFile(gen, f)
        }
        return nil
    })
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
    // Generate .api file
    filename := getAPIFilename(file)
    g := gen.NewGeneratedFile(filename, file.GoImportPath)

    // Parse and generate
    generator := NewGenerator(file)
    content, err := generator.Generate()
    if err != nil {
        gen.Error(err)
        return
    }

    g.P(content)
}

func getAPIFilename(file *protogen.File) string {
    // rpc/desc/user.proto -> api/desc/core/user.api
    name := file.GeneratedFilenamePrefix
    // Transform path: rpc/desc/user -> api/desc/core/user
    return transformPath(name) + ".api"
}
```

**Testing**:
```bash
cd tools/protoc-gen-go-zero-api
go mod init github.com/suyuan32/simple-admin-core/tools/protoc-gen-go-zero-api
go get google.golang.org/protobuf/compiler/protogen
go build -o protoc-gen-go-zero-api
```

#### Task 1.2: Define Go-Zero Custom Proto Options (3 hours)

**File**: `rpc/desc/go_zero/options.proto`

**Approach**:
Create custom Proto extensions for Go-Zero specific features following Google's extension pattern.

**Code Example**:
```protobuf
// rpc/desc/go_zero/options.proto
syntax = "proto3";

package go_zero;

import "google/protobuf/descriptor.proto";

option go_package = "github.com/suyuan32/simple-admin-core/rpc/types/go_zero";

// Service-level options
extend google.protobuf.ServiceOptions {
  // JWT authentication config name (e.g., "Auth")
  string jwt = 50001;

  // Comma-separated middleware list (e.g., "Authority,RateLimit")
  string middleware = 50002;

  // Route group name (e.g., "user", "admin")
  string group = 50003;

  // Route prefix (e.g., "/api/v1")
  string prefix = 50004;
}

// Method-level options
extend google.protobuf.MethodOptions {
  // Mark method as public (no JWT required)
  bool public = 50011;

  // Method-specific middleware (overrides service-level)
  string middleware = 50012;
}

// File-level options
extend google.protobuf.FileOptions {
  // API metadata for info() section
  ApiInfo api_info = 50021;
}

// API metadata message
message ApiInfo {
  string title = 1;
  string desc = 2;
  string author = 3;
  string email = 4;
  string version = 5;
}
```

**Generate Go code**:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
       rpc/desc/go_zero/options.proto
```

**Testing**:
```go
// test/options_test.go
func TestGoZeroOptions(t *testing.T) {
    // Verify options are accessible
    opts := &descriptorpb.ServiceOptions{}
    proto.SetExtension(opts, go_zero.E_Jwt, "Auth")

    jwt := proto.GetExtension(opts, go_zero.E_Jwt).(string)
    assert.Equal(t, "Auth", jwt)
}
```

#### Task 1.3: Create Internal Model Structures (3 hours)

**File**: `tools/protoc-gen-go-zero-api/model/`

**Approach**:
Define internal models to represent parsed Proto structures before generation.

**Code Example** - `model/service.go`:
```go
package model

// Service represents a Proto service
type Service struct {
    Name        string
    Methods     []*Method
    Options     *ServerOptions
    PackageName string
}

// ServerOptions represents @server block options
type ServerOptions struct {
    JWT        string   // JWT config name
    Middleware []string // Middleware list
    Group      string   // Route group
    Prefix     string   // Route prefix
}

func (s *ServerOptions) Signature() string {
    // Generate unique key for grouping
    return fmt.Sprintf("jwt:%s|mw:%s|group:%s",
        s.JWT,
        strings.Join(s.Middleware, ","),
        s.Group)
}

func (s *ServerOptions) IsEmpty() bool {
    return s.JWT == "" && len(s.Middleware) == 0 && s.Group == ""
}
```

**Code Example** - `model/method.go`:
```go
package model

// Method represents a Proto RPC method
type Method struct {
    Name         string
    RequestType  string
    ResponseType string
    HTTPRule     *HTTPRule
    Options      *MethodOptions
}

// HTTPRule represents google.api.http annotation
type HTTPRule struct {
    Method string // GET, POST, PUT, DELETE, PATCH
    Path   string // e.g., "/user/{id}"
    Body   string // e.g., "*" or "user"
}

// MethodOptions represents method-level options
type MethodOptions struct {
    Public     bool     // No JWT required
    Middleware []string // Method-specific middleware
}

func (m *Method) HandlerName() string {
    // CreateUser -> createUser
    if len(m.Name) == 0 {
        return ""
    }
    return strings.ToLower(m.Name[:1]) + m.Name[1:]
}
```

**Code Example** - `model/message.go`:
```go
package model

// Message represents a Proto message (type definition)
type Message struct {
    Name   string
    Fields []*Field
}

// Field represents a message field
type Field struct {
    Name       string
    Type       string
    ProtoType  string
    JSONTag    string
    Optional   bool
    Repeated   bool
}

func (f *Field) GoZeroType() string {
    // Convert Proto type to Go type
    switch f.ProtoType {
    case "string":
        return "string"
    case "int32", "int64", "uint32", "uint64":
        return f.ProtoType
    case "bool":
        return "bool"
    case "bytes":
        return "[]byte"
    default:
        // Custom message type
        return f.Type
    }
}
```

**Testing**:
```go
func TestServiceModel(t *testing.T) {
    opts := &ServerOptions{
        JWT:        "Auth",
        Middleware: []string{"Authority", "RateLimit"},
        Group:      "user",
    }

    sig := opts.Signature()
    assert.Contains(t, sig, "jwt:Auth")
    assert.Contains(t, sig, "mw:Authority,RateLimit")
}

func TestMethodHandlerName(t *testing.T) {
    method := &Method{Name: "CreateUser"}
    assert.Equal(t, "createUser", method.HandlerName())
}
```

### Phase 2: Parser Implementation (16-20 hours)

#### Task 2.1: Implement HTTP Annotation Parser (6 hours)

**File**: `tools/protoc-gen-go-zero-api/generator/http_parser.go`

**Approach**:
Extract `google.api.http` annotations from Proto MethodOptions.

**Code Example**:
```go
package generator

import (
    "fmt"
    "strings"

    "google.golang.org/genproto/googleapis/api/annotations"
    "google.golang.org/protobuf/compiler/protogen"
    "google.golang.org/protobuf/proto"

    "github.com/suyuan32/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

type HTTPParser struct{}

func NewHTTPParser() *HTTPParser {
    return &HTTPParser{}
}

func (p *HTTPParser) Parse(method *protogen.Method) (*model.HTTPRule, error) {
    // Get google.api.http option
    opts := method.Desc.Options()
    if opts == nil {
        return nil, nil
    }

    httpRule := proto.GetExtension(opts, annotations.E_Http)
    if httpRule == nil {
        return nil, nil
    }

    rule := httpRule.(*annotations.HttpRule)

    // Extract HTTP method and path
    var httpMethod, path, body string

    switch pattern := rule.Pattern.(type) {
    case *annotations.HttpRule_Get:
        httpMethod = "get"
        path = pattern.Get
    case *annotations.HttpRule_Post:
        httpMethod = "post"
        path = pattern.Post
    case *annotations.HttpRule_Put:
        httpMethod = "put"
        path = pattern.Put
    case *annotations.HttpRule_Delete:
        httpMethod = "delete"
        path = pattern.Delete
    case *annotations.HttpRule_Patch:
        httpMethod = "patch"
        path = pattern.Patch
    default:
        return nil, fmt.Errorf("unsupported HTTP pattern for method %s", method.Desc.Name())
    }

    // Get body mapping
    body = rule.Body
    if body == "" {
        body = "*" // Default to entire request
    }

    // Transform path: /user/{id} -> /user/:id (Go-Zero format)
    path = p.transformPath(path)

    return &model.HTTPRule{
        Method: httpMethod,
        Path:   path,
        Body:   body,
    }, nil
}

func (p *HTTPParser) transformPath(path string) string {
    // Transform {param} to :param
    // /user/{id}/posts/{post_id} -> /user/:id/posts/:post_id
    result := path
    for {
        start := strings.Index(result, "{")
        if start == -1 {
            break
        }
        end := strings.Index(result[start:], "}")
        if end == -1 {
            break
        }
        end += start

        param := result[start+1 : end]
        result = result[:start] + ":" + param + result[end+1:]
    }
    return result
}

func (p *HTTPParser) ParseAdditionalBindings(method *protogen.Method) ([]*model.HTTPRule, error) {
    // Handle additional_bindings for multiple routes
    opts := method.Desc.Options()
    if opts == nil {
        return nil, nil
    }

    httpRule := proto.GetExtension(opts, annotations.E_Http)
    if httpRule == nil {
        return nil, nil
    }

    rule := httpRule.(*annotations.HttpRule)
    if len(rule.AdditionalBindings) == 0 {
        return nil, nil
    }

    var rules []*model.HTTPRule
    for _, binding := range rule.AdditionalBindings {
        // Parse each binding
        // Similar logic to Parse()
        // ...
    }

    return rules, nil
}
```

**Testing**:
```go
func TestHTTPParser(t *testing.T) {
    parser := NewHTTPParser()

    // Test path transformation
    path := parser.transformPath("/user/{id}/posts/{post_id}")
    assert.Equal(t, "/user/:id/posts/:post_id", path)

    // Test with actual Proto method (mock)
    // ...
}
```

#### Task 2.2: Implement Go-Zero Options Parser (5 hours)

**File**: `tools/protoc-gen-go-zero-api/generator/options_parser.go`

**Approach**:
Extract Go-Zero custom options from Proto.

**Code Example**:
```go
package generator

import (
    "strings"

    "google.golang.org/protobuf/compiler/protogen"
    "google.golang.org/protobuf/proto"

    go_zero "github.com/suyuan32/simple-admin-core/rpc/types/go_zero"
    "github.com/suyuan32/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

type OptionsParser struct{}

func NewOptionsParser() *OptionsParser {
    return &OptionsParser{}
}

// Parse service-level options
func (p *OptionsParser) ParseServiceOptions(service *protogen.Service) *model.ServerOptions {
    opts := service.Desc.Options()
    if opts == nil {
        return &model.ServerOptions{}
    }

    serverOpts := &model.ServerOptions{}

    // JWT
    if jwt := proto.GetExtension(opts, go_zero.E_Jwt); jwt != nil {
        serverOpts.JWT = jwt.(string)
    }

    // Middleware
    if mw := proto.GetExtension(opts, go_zero.E_Middleware); mw != nil {
        middleware := mw.(string)
        if middleware != "" {
            serverOpts.Middleware = strings.Split(middleware, ",")
            // Trim spaces
            for i := range serverOpts.Middleware {
                serverOpts.Middleware[i] = strings.TrimSpace(serverOpts.Middleware[i])
            }
        }
    }

    // Group
    if group := proto.GetExtension(opts, go_zero.E_Group); group != nil {
        serverOpts.Group = group.(string)
    } else {
        // Default: lowercase service name
        serverOpts.Group = strings.ToLower(string(service.Desc.Name()))
    }

    // Prefix
    if prefix := proto.GetExtension(opts, go_zero.E_Prefix); prefix != nil {
        serverOpts.Prefix = prefix.(string)
    }

    return serverOpts
}

// Parse method-level options
func (p *OptionsParser) ParseMethodOptions(method *protogen.Method) *model.MethodOptions {
    opts := method.Desc.Options()
    if opts == nil {
        return &model.MethodOptions{}
    }

    methodOpts := &model.MethodOptions{}

    // Public (no JWT)
    if public := proto.GetExtension(opts, go_zero.E_Public); public != nil {
        methodOpts.Public = public.(bool)
    }

    // Method-specific middleware
    if mw := proto.GetExtension(opts, go_zero.E_Middleware); mw != nil {
        middleware := mw.(string)
        if middleware != "" {
            methodOpts.Middleware = strings.Split(middleware, ",")
            for i := range methodOpts.Middleware {
                methodOpts.Middleware[i] = strings.TrimSpace(methodOpts.Middleware[i])
            }
        }
    }

    return methodOpts
}

// Parse file-level API info
func (p *OptionsParser) ParseAPIInfo(file *protogen.File) *go_zero.ApiInfo {
    opts := file.Desc.Options()
    if opts == nil {
        return nil
    }

    if apiInfo := proto.GetExtension(opts, go_zero.E_ApiInfo); apiInfo != nil {
        return apiInfo.(*go_zero.ApiInfo)
    }

    return nil
}

// Merge method options with service options
func (p *OptionsParser) MergeOptions(serviceOpts *model.ServerOptions, methodOpts *model.MethodOptions) *model.ServerOptions {
    merged := &model.ServerOptions{
        JWT:        serviceOpts.JWT,
        Middleware: make([]string, len(serviceOpts.Middleware)),
        Group:      serviceOpts.Group,
        Prefix:     serviceOpts.Prefix,
    }
    copy(merged.Middleware, serviceOpts.Middleware)

    // Method overrides
    if methodOpts.Public {
        merged.JWT = "" // Remove JWT requirement
    }

    if len(methodOpts.Middleware) > 0 {
        merged.Middleware = methodOpts.Middleware
    }

    return merged
}
```

**Testing**:
```go
func TestOptionsParser(t *testing.T) {
    parser := NewOptionsParser()

    // Test merge
    serviceOpts := &model.ServerOptions{
        JWT:        "Auth",
        Middleware: []string{"Authority"},
        Group:      "user",
    }

    methodOpts := &model.MethodOptions{
        Public: true,
    }

    merged := parser.MergeOptions(serviceOpts, methodOpts)
    assert.Equal(t, "", merged.JWT) // JWT removed for public endpoint
    assert.Equal(t, "user", merged.Group)
}
```

#### Task 2.3: Implement Type Converter (5 hours)

**File**: `tools/protoc-gen-go-zero-api/generator/type_converter.go`

**Approach**:
Convert Proto messages to Go-Zero type definitions.

**Code Example**:
```go
package generator

import (
    "fmt"
    "strings"

    "google.golang.org/protobuf/compiler/protogen"
    "google.golang.org/protobuf/reflect/protoreflect"

    "github.com/suyuan32/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

type TypeConverter struct {
    convertedTypes map[string]bool // Track converted types to avoid duplicates
}

func NewTypeConverter() *TypeConverter {
    return &TypeConverter{
        convertedTypes: make(map[string]bool),
    }
}

func (c *TypeConverter) ConvertMessage(msg *protogen.Message) *model.Message {
    if c.convertedTypes[string(msg.Desc.FullName())] {
        return nil // Already converted
    }
    c.convertedTypes[string(msg.Desc.FullName())] = true

    message := &model.Message{
        Name:   string(msg.Desc.Name()),
        Fields: []*model.Field{},
    }

    for _, field := range msg.Fields {
        message.Fields = append(message.Fields, c.convertField(field))
    }

    return message
}

func (c *TypeConverter) convertField(field *protogen.Field) *model.Field {
    f := &model.Field{
        Name:      string(field.Desc.Name()),
        JSONTag:   field.Desc.JSONName(),
        Optional:  field.Desc.HasOptionalKeyword(),
        Repeated:  field.Desc.Cardinality() == protoreflect.Repeated,
    }

    // Convert type
    f.Type, f.ProtoType = c.convertType(field)

    return f
}

func (c *TypeConverter) convertType(field *protogen.Field) (goType, protoType string) {
    kind := field.Desc.Kind()

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
        return "float32", "float32"
    case protoreflect.DoubleKind:
        return "float64", "float64"
    case protoreflect.StringKind:
        return "string", "string"
    case protoreflect.BytesKind:
        return "[]byte", "bytes"
    case protoreflect.EnumKind:
        // Enum -> string in Go-Zero
        return "string", "enum"
    case protoreflect.MessageKind:
        // Custom message type
        msgName := string(field.Message.Desc.Name())
        return msgName, "message"
    default:
        return "interface{}", "unknown"
    }
}

func (c *TypeConverter) GenerateTypeDefinition(msg *model.Message) string {
    var b strings.Builder

    b.WriteString(fmt.Sprintf("    %s {\n", msg.Name))

    for _, field := range msg.Fields {
        // Generate field line
        fieldType := field.Type
        if field.Repeated {
            fieldType = "[]" + fieldType
        }

        // Optional fields use pointer
        if field.Optional {
            fieldType = "*" + fieldType
        }

        tags := fmt.Sprintf("`json:\"%s", field.JSONTag)
        if field.Optional {
            tags += ",optional"
        }
        tags += "\"`"

        b.WriteString(fmt.Sprintf("        %s %s %s\n",
            capitalize(field.Name),
            fieldType,
            tags))
    }

    b.WriteString("    }\n")

    return b.String()
}

func capitalize(s string) string {
    if len(s) == 0 {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

**Testing**:
```go
func TestTypeConverter(t *testing.T) {
    converter := NewTypeConverter()

    // Test basic type conversion
    goType, protoType := converter.convertType(mockStringField())
    assert.Equal(t, "string", goType)
    assert.Equal(t, "string", protoType)

    // Test optional field
    field := &model.Field{
        Name:     "username",
        Type:     "string",
        JSONTag:  "username",
        Optional: true,
    }
    // Should generate: Username *string `json:"username,optional"`
}
```

### Phase 3: Service Grouping and Generation (12-16 hours)

#### Task 3.1: Implement Service Grouper (5 hours)

**File**: `tools/protoc-gen-go-zero-api/generator/grouper.go`

**Approach**:
Group methods by their @server configuration to generate separate service blocks.

**Code Example**:
```go
package generator

import (
    "github.com/suyuan32/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

type ServiceGrouper struct{}

func NewServiceGrouper() *ServiceGrouper {
    return &ServiceGrouper{}
}

// ServiceGroup represents a group of methods with same @server config
type ServiceGroup struct {
    ServerOptions *model.ServerOptions
    Methods       []*model.Method
}

func (g *ServiceGrouper) GroupMethods(service *model.Service) []*ServiceGroup {
    // Group methods by their effective @server options
    groups := make(map[string]*ServiceGroup)

    for _, method := range service.Methods {
        // Get effective options for this method
        opts := method.Options
        signature := opts.Signature()

        if group, exists := groups[signature]; exists {
            group.Methods = append(group.Methods, method)
        } else {
            groups[signature] = &ServiceGroup{
                ServerOptions: opts,
                Methods:       []*model.Method{method},
            }
        }
    }

    // Convert map to slice
    var result []*ServiceGroup
    for _, group := range groups {
        result = append(result, group)
    }

    // Sort groups: protected endpoints first, then public
    // This ensures consistent output
    sortGroups(result)

    return result
}

func sortGroups(groups []*ServiceGroup) {
    // Sort by: JWT present > middleware count > group name
    sort.Slice(groups, func(i, j int) bool {
        a, b := groups[i].ServerOptions, groups[j].ServerOptions

        // JWT endpoints first
        if (a.JWT != "") != (b.JWT != "") {
            return a.JWT != ""
        }

        // More middleware first
        if len(a.Middleware) != len(b.Middleware) {
            return len(a.Middleware) > len(b.Middleware)
        }

        // Alphabetical by group
        return a.Group < b.Group
    })
}
```

**Testing**:
```go
func TestServiceGrouper(t *testing.T) {
    grouper := NewServiceGrouper()

    service := &model.Service{
        Name: "User",
        Methods: []*model.Method{
            {
                Name: "CreateUser",
                Options: &model.ServerOptions{
                    JWT: "Auth",
                    Middleware: []string{"Authority"},
                    Group: "user",
                },
            },
            {
                Name: "Login",
                Options: &model.ServerOptions{
                    Group: "user", // No JWT
                },
            },
            {
                Name: "GetUser",
                Options: &model.ServerOptions{
                    JWT: "Auth",
                    Middleware: []string{"Authority"},
                    Group: "user",
                },
            },
        },
    }

    groups := grouper.GroupMethods(service)

    // Should have 2 groups:
    // 1. JWT + Authority (CreateUser, GetUser)
    // 2. No JWT (Login)
    assert.Equal(t, 2, len(groups))
    assert.Equal(t, 2, len(groups[0].Methods)) // Protected group
    assert.Equal(t, 1, len(groups[1].Methods)) // Public group
}
```

#### Task 3.2: Implement .api Template Generator (7 hours)

**File**: `tools/protoc-gen-go-zero-api/generator/template.go`

**Approach**:
Use Go text templates to generate .api file format.

**Code Example**:
```go
package generator

import (
    "bytes"
    "strings"
    "text/template"

    "github.com/suyuan32/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

const apiTemplate = `syntax = "v1"

{{- if .APIInfo }}

info(
    title: "{{ .APIInfo.Title }}"
    desc: "{{ .APIInfo.Desc }}"
    author: "{{ .APIInfo.Author }}"
    email: "{{ .APIInfo.Email }}"
    version: "{{ .APIInfo.Version }}"
)
{{- end }}

import "../base.api"

{{- if .Types }}

type (
{{- range .Types }}
    {{ .Definition }}
{{- end }}
)
{{- end }}

{{- range .ServiceGroups }}

{{- if .ServerOptions }}
@server(
{{- if .ServerOptions.JWT }}
    jwt: {{ .ServerOptions.JWT }}
{{- end }}
{{- if .ServerOptions.Group }}
    group: {{ .ServerOptions.Group }}
{{- end }}
{{- if .ServerOptions.Middleware }}
    middleware: {{ join .ServerOptions.Middleware "," }}
{{- end }}
)
{{- end }}

service Core {
{{- range .Methods }}
    @handler {{ .HandlerName }}
    {{ .HTTPRule.Method }} {{ .HTTPRule.Path }} ({{ .RequestType }}) returns ({{ .ResponseType }})
{{- end }}
}
{{- end }}
`

type TemplateData struct {
    APIInfo       *go_zero.ApiInfo
    Types         []*TypeDefinition
    ServiceGroups []*ServiceGroup
}

type TypeDefinition struct {
    Name       string
    Definition string
}

type TemplateGenerator struct {
    tmpl *template.Template
}

func NewTemplateGenerator() *TemplateGenerator {
    funcMap := template.FuncMap{
        "join": strings.Join,
    }

    tmpl := template.Must(template.New("api").Funcs(funcMap).Parse(apiTemplate))

    return &TemplateGenerator{
        tmpl: tmpl,
    }
}

func (g *TemplateGenerator) Generate(data *TemplateData) (string, error) {
    var buf bytes.Buffer

    if err := g.tmpl.Execute(&buf, data); err != nil {
        return "", err
    }

    return buf.String(), nil
}
```

**Testing**:
```go
func TestTemplateGenerator(t *testing.T) {
    gen := NewTemplateGenerator()

    data := &TemplateData{
        APIInfo: &go_zero.ApiInfo{
            Title:   "User API",
            Desc:    "User management",
            Author:  "Test",
            Email:   "test@example.com",
            Version: "v1.0",
        },
        Types: []*TypeDefinition{
            {
                Name: "CreateUserReq",
                Definition: "    CreateUserReq {\n        Username string `json:\"username\"`\n    }",
            },
        },
        ServiceGroups: []*ServiceGroup{
            {
                ServerOptions: &model.ServerOptions{
                    JWT:   "Auth",
                    Group: "user",
                },
                Methods: []*model.Method{
                    {
                        Name:         "CreateUser",
                        RequestType:  "CreateUserReq",
                        ResponseType: "BaseIDResp",
                        HTTPRule: &model.HTTPRule{
                            Method: "post",
                            Path:   "/user/create",
                        },
                    },
                },
            },
        },
    }

    output, err := gen.Generate(data)
    assert.NoError(t, err)
    assert.Contains(t, output, "info(")
    assert.Contains(t, output, "@server(")
    assert.Contains(t, output, "jwt: Auth")
}
```

### Phase 4: Integration and Testing (12-16 hours)

#### Task 4.1: Integrate Parser and Generator (4 hours)

**File**: `tools/protoc-gen-go-zero-api/generator/generator.go`

**Code Example**:
```go
package generator

import (
    "fmt"

    "google.golang.org/protobuf/compiler/protogen"

    "github.com/suyuan32/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

type Generator struct {
    file            *protogen.File
    httpParser      *HTTPParser
    optionsParser   *OptionsParser
    typeConverter   *TypeConverter
    grouper         *ServiceGrouper
    templateGen     *TemplateGenerator
}

func NewGenerator(file *protogen.File) *Generator {
    return &Generator{
        file:          file,
        httpParser:    NewHTTPParser(),
        optionsParser: NewOptionsParser(),
        typeConverter: NewTypeConverter(),
        grouper:       NewServiceGrouper(),
        templateGen:   NewTemplateGenerator(),
    }
}

func (g *Generator) Generate() (string, error) {
    // 1. Parse API info
    apiInfo := g.optionsParser.ParseAPIInfo(g.file)

    // 2. Collect all messages (types)
    var messages []*model.Message
    for _, msg := range g.file.Messages {
        if converted := g.typeConverter.ConvertMessage(msg); converted != nil {
            messages = append(messages, converted)
        }
    }

    // 3. Parse services
    var allServiceGroups []*ServiceGroup

    for _, service := range g.file.Services {
        serviceModel := g.parseService(service)
        groups := g.grouper.GroupMethods(serviceModel)
        allServiceGroups = append(allServiceGroups, groups...)
    }

    // 4. Generate type definitions
    var typeDefs []*TypeDefinition
    for _, msg := range messages {
        typeDefs = append(typeDefs, &TypeDefinition{
            Name:       msg.Name,
            Definition: g.typeConverter.GenerateTypeDefinition(msg),
        })
    }

    // 5. Prepare template data
    data := &TemplateData{
        APIInfo:       apiInfo,
        Types:         typeDefs,
        ServiceGroups: allServiceGroups,
    }

    // 6. Generate .api content
    return g.templateGen.Generate(data)
}

func (g *Generator) parseService(service *protogen.Service) *model.Service {
    serviceOpts := g.optionsParser.ParseServiceOptions(service)

    serviceModel := &model.Service{
        Name:    string(service.Desc.Name()),
        Methods: []*model.Method{},
        Options: serviceOpts,
    }

    for _, method := range service.Methods {
        methodModel := g.parseMethod(method, serviceOpts)
        if methodModel != nil {
            serviceModel.Methods = append(serviceModel.Methods, methodModel)
        }
    }

    return serviceModel
}

func (g *Generator) parseMethod(method *protogen.Method, serviceOpts *model.ServerOptions) *model.Method {
    // Parse HTTP annotation
    httpRule, err := g.httpParser.Parse(method)
    if err != nil {
        return nil // Skip methods without valid HTTP annotations
    }
    if httpRule == nil {
        return nil // No HTTP annotation
    }

    // Parse method options
    methodOpts := g.optionsParser.ParseMethodOptions(method)

    // Merge with service options
    effectiveOpts := g.optionsParser.MergeOptions(serviceOpts, methodOpts)

    return &model.Method{
        Name:         string(method.Desc.Name()),
        RequestType:  string(method.Input.Desc.Name()),
        ResponseType: string(method.Output.Desc.Name()),
        HTTPRule:     httpRule,
        Options:      effectiveOpts,
    }
}
```

#### Task 4.2: Update Makefile Integration (2 hours)

**File**: `Makefile`

**Code Example**:
```makefile
# Makefile additions

# Build protoc-gen-go-zero-api plugin
.PHONY: build-proto-plugin
build-proto-plugin:
	@echo "Building protoc-gen-go-zero-api plugin..."
	cd tools/protoc-gen-go-zero-api && go build -o ../../bin/protoc-gen-go-zero-api
	@echo "Plugin built successfully at bin/protoc-gen-go-zero-api"

# Generate Go-Zero custom options
.PHONY: gen-go-zero-options
gen-go-zero-options:
	@echo "Generating Go-Zero custom options..."
	protoc --go_out=. --go_opt=paths=source_relative \
	       rpc/desc/go_zero/options.proto

# Generate .api files from Proto
.PHONY: gen-proto-api
gen-proto-api: build-proto-plugin
	@echo "Generating .api files from Proto..."
	protoc --plugin=protoc-gen-go-zero-api=./bin/protoc-gen-go-zero-api \
	       --go-zero-api_out=api/desc \
	       --go-zero-api_opt=paths=source_relative \
	       --proto_path=. \
	       --proto_path=third_party \
	       rpc/desc/**/*.proto
	@echo ".api files generated successfully"

# Generate API code from .api files
.PHONY: gen-api-code
gen-api-code:
	@echo "Generating API code from .api files..."
	goctl api go -api api/desc/all.api -dir api/ --style=$(API_STYLE)

# Combined: Proto → .api → Go code
.PHONY: gen-api-all
gen-api-all: gen-proto-api gen-api-code
	@echo "API generation complete"

# Update gen-all to include new workflow
.PHONY: gen-all
gen-all: gen-ent gen-rpc gen-api-all
	@echo "All code generation completed!"

# Validate generated .api files compile
.PHONY: validate-api
validate-api: gen-proto-api
	@echo "Validating generated .api files..."
	@for file in api/desc/**/*.api; do \
		echo "Checking $$file..."; \
		goctl api validate -api $$file || exit 1; \
	done
	@echo "All .api files are valid"
```

**Testing**:
```bash
# Test the new workflow
make gen-proto-api
make validate-api
make gen-api-all
```

#### Task 4.3: Create Integration Tests (6 hours)

**File**: `tools/protoc-gen-go-zero-api/test/integration_test.go`

**Code Example**:
```go
package test

import (
    "os"
    "os/exec"
    "path/filepath"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestProtoToAPIGeneration(t *testing.T) {
    // Setup
    tmpDir := t.TempDir()

    // Copy test proto file
    testProto := `
syntax = "proto3";

package test.v1;

import "google/api/annotations.proto";
import "go_zero/options.proto";

option go_package = "test/v1;v1";

option (go_zero.api_info) = {
  title: "Test API"
  desc: "Test description"
  author: "Test"
  version: "v1.0"
};

service TestService {
  option (go_zero.jwt) = "Auth";
  option (go_zero.middleware) = "Authority";
  option (go_zero.group) = "test";

  rpc CreateItem(CreateItemReq) returns (CreateItemResp) {
    option (google.api.http) = {
      post: "/item/create"
      body: "*"
    };
  }

  rpc GetItem(GetItemReq) returns (GetItemResp) {
    option (google.api.http) = {
      post: "/item"
    };
  }
}

message CreateItemReq {
  string name = 1;
  int32 value = 2;
}

message CreateItemResp {
  uint64 id = 1;
}

message GetItemReq {
  uint64 id = 1;
}

message GetItemResp {
  uint64 id = 1;
  string name = 2;
  int32 value = 3;
}
`

    protoFile := filepath.Join(tmpDir, "test.proto")
    err := os.WriteFile(protoFile, []byte(testProto), 0644)
    require.NoError(t, err)

    // Run protoc with plugin
    cmd := exec.Command("protoc",
        "--plugin=protoc-gen-go-zero-api=../../bin/protoc-gen-go-zero-api",
        "--go-zero-api_out="+tmpDir,
        "--proto_path=.",
        protoFile,
    )

    output, err := cmd.CombinedOutput()
    require.NoError(t, err, "protoc failed: %s", output)

    // Read generated .api file
    apiFile := filepath.Join(tmpDir, "test.api")
    content, err := os.ReadFile(apiFile)
    require.NoError(t, err)

    apiContent := string(content)

    // Assertions
    assert.Contains(t, apiContent, `syntax = "v1"`)
    assert.Contains(t, apiContent, `title: "Test API"`)
    assert.Contains(t, apiContent, `@server(`)
    assert.Contains(t, apiContent, `jwt: Auth`)
    assert.Contains(t, apiContent, `middleware: Authority`)
    assert.Contains(t, apiContent, `group: test`)
    assert.Contains(t, apiContent, `@handler createItem`)
    assert.Contains(t, apiContent, `post /item/create (CreateItemReq) returns (CreateItemResp)`)

    // Validate types generated
    assert.Contains(t, apiContent, `CreateItemReq {`)
    assert.Contains(t, apiContent, `Name string`)
    assert.Contains(t, apiContent, `Value int32`)
}

func TestPublicEndpoint(t *testing.T) {
    // Test that public endpoints generate without JWT
    // Similar structure to above
}

func TestMultipleServiceGroups(t *testing.T) {
    // Test that methods with different middleware generate separate service blocks
}

func TestGoldenFiles(t *testing.T) {
    // Golden file testing
    fixtures := []string{"user.proto", "role.proto", "menu.proto"}

    for _, fixture := range fixtures {
        t.Run(fixture, func(t *testing.T) {
            // Generate from proto
            generated := generateFromProto(t, fixture)

            // Compare with expected .api file
            goldenFile := filepath.Join("fixtures", "expected", fixture+".api")
            expected, err := os.ReadFile(goldenFile)
            require.NoError(t, err)

            assert.Equal(t, string(expected), generated, "Generated output doesn't match golden file")
        })
    }
}
```

### Phase 5: Pilot Migration and Documentation (12-16 hours)

#### Task 5.1: Migrate User Module (Pilot) (4 hours)

**Steps**:

1. **Add Go-Zero options to user.proto**:
```protobuf
// rpc/desc/user.proto
syntax = "proto3";

package core.v1;

import "google/api/annotations.proto";
import "go_zero/options.proto";

option go_package = "github.com/suyuan32/simple-admin-core/rpc/types/core";

option (go_zero.api_info) = {
  title: "User Management"
  desc: "User management and authentication"
  author: "Ryan Su"
  email: "yuansu.china.work@gmail.com"
  version: "v1.0"
};

service User {
  option (go_zero.jwt) = "Auth";
  option (go_zero.middleware) = "Authority";
  option (go_zero.group) = "user";

  rpc CreateUser(CreateUserReq) returns (BaseIDResp) {
    option (google.api.http) = {
      post: "/user/create"
      body: "*"
    };
  }

  rpc Login(LoginReq) returns (LoginResp) {
    option (google.api.http) = {
      post: "/user/login"
      body: "*"
    };
    option (go_zero.public) = true;  // Public endpoint
  }

  // ... other methods
}
```

2. **Generate .api file**:
```bash
make gen-proto-api
```

3. **Compare with existing .api**:
```bash
diff api/desc/core/user.api api/desc/core/user.api.generated
```

4. **Backup and replace**:
```bash
cp api/desc/core/user.api api/desc/core/user.api.backup
cp api/desc/core/user.api.generated api/desc/core/user.api
```

5. **Regenerate API code and test**:
```bash
make gen-api-code
go build ./api/...
go test ./api/internal/logic/user/...
```

#### Task 5.2: Create Migration Guide (4 hours)

**File**: `docs/proto-first-migration-guide.md`

**Content**:
```markdown
# Proto-First API Migration Guide

## Overview

This guide helps migrate existing modules from dual API definitions (Proto + .api) to Proto-First approach.

## Prerequisites

- Install protoc-gen-go-zero-api plugin: `make build-proto-plugin`
- Ensure Go-Zero options are generated: `make gen-go-zero-options`

## Migration Steps

### Step 1: Add Go-Zero Options to Proto

Add service-level options:
```protobuf
service YourService {
  option (go_zero.jwt) = "Auth";           // If JWT required
  option (go_zero.middleware) = "Authority"; // Middleware
  option (go_zero.group) = "your_group";   // Route group

  // ... methods
}
```

Add method-level overrides:
```protobuf
rpc PublicMethod(...) returns (...) {
  option (google.api.http) = {...};
  option (go_zero.public) = true;  // Override JWT
}
```

### Step 2: Generate .api File

```bash
make gen-proto-api
```

### Step 3: Validate

```bash
# Validate generated .api
make validate-api

# Compare with existing
diff api/desc/core/your_module.api api/desc/core/your_module.api.generated
```

### Step 4: Backup and Replace

```bash
cp api/desc/core/your_module.api api/desc/core/your_module.api.backup
mv api/desc/core/your_module.api.generated api/desc/core/your_module.api
```

### Step 5: Test

```bash
make gen-api-code
go build ./api/...
go test ./api/...
```

## Troubleshooting

### Generated .api doesn't compile

- Check Proto message field names match expected Go-Zero format
- Verify all HTTP annotations are valid

### Missing @server annotations

- Ensure Go-Zero options are set at service level
- Check plugin is latest version

## Rollback

```bash
cp api/desc/core/your_module.api.backup api/desc/core/your_module.api
make gen-api-code
```
```

#### Task 5.3: Update Documentation (4 hours)

**Update CLAUDE.md**:
```markdown
## Proto-First API Generation (New)

This project now supports Proto-First API generation, eliminating the need to maintain separate `.api` files.

### Quick Start

1. **Define API in Proto with Go-Zero options**:
```protobuf
service User {
  option (go_zero.jwt) = "Auth";
  option (go_zero.group) = "user";

  rpc CreateUser(CreateUserReq) returns (BaseIDResp) {
    option (google.api.http) = {
      post: "/user/create"
      body: "*"
    };
  }
}
```

2. **Generate .api and code**:
```bash
make gen-api-all  # Proto → .api → Go code
```

### Available Commands

- `make gen-proto-api` - Generate .api from Proto
- `make gen-api-all` - Complete API generation pipeline
- `make validate-api` - Validate generated .api files

### Go-Zero Custom Options

See `rpc/desc/go_zero/options.proto` for available options:

- `(go_zero.jwt)` - JWT config name
- `(go_zero.middleware)` - Middleware list
- `(go_zero.group)` - Route group
- `(go_zero.public)` - Mark method as public (no JWT)

### Migration

See [Migration Guide](../specs/003-proto-first-api-generation/migration-guide.md)
```

## Performance Considerations

### Bundle Size Impact

- Plugin binary: ~8-10 MB (statically linked Go binary)
- No runtime impact on services
- Generated .api files are text, minimal size increase

### Runtime Performance

- Code generation only (no runtime component)
- No impact on API service performance
- Slightly faster development cycle (less manual work)

### Optimization Strategies

1. **Parallel Proto Processing**:
```go
// Process multiple proto files in parallel
var wg sync.WaitGroup
for _, file := range protoFiles {
    wg.Add(1)
    go func(f *protogen.File) {
        defer wg.Done()
        generateFile(gen, f)
    }(file)
}
wg.Wait()
```

2. **Caching**:
```go
// Cache converted types to avoid duplicate processing
var typeCache sync.Map

func (c *TypeConverter) ConvertMessage(msg *protogen.Message) *model.Message {
    if cached, ok := typeCache.Load(msg.Desc.FullName()); ok {
        return cached.(*model.Message)
    }

    converted := c.doConvert(msg)
    typeCache.Store(msg.Desc.FullName(), converted)
    return converted
}
```

## Deployment Strategy

### Rollout Plan

**Week 1: Plugin Development & Testing**
- Days 1-3: Core plugin implementation
- Days 4-5: Unit tests and integration tests
- Deliverable: Working plugin with 80%+ test coverage

**Week 2: Pilot Migration**
- Days 1-2: Migrate User module
- Day 3: Migrate Role module
- Days 4-5: Fix issues, refine plugin
- Deliverable: 2 modules successfully migrated

**Week 3: Full Migration & Documentation**
- Days 1-3: Migrate remaining 13 modules
- Day 4: Documentation and training materials
- Day 5: Team training session
- Deliverable: All modules migrated, team trained

### Feature Flag

No feature flag needed (build-time generation, not runtime).

### Monitoring

Track in CI/CD:
- Proto → .api generation time
- .api file validation success rate
- API service build success rate

## Rollback Plan

1. **Keep original .api files as backup**:
```bash
git checkout develop -- api/desc/core/*.api
```

2. **Disable plugin in Makefile**:
```makefile
# Comment out gen-proto-api target
# gen-all: gen-ent gen-rpc gen-api-code  # Skip gen-proto-api
```

3. **Revert to manual .api maintenance**:
```bash
# Use original workflow
vim api/desc/core/user.api
make gen-api-code
```

## Success Metrics

**Technical Metrics**:
- ✅ Plugin generates valid .api for 100% of Proto files
- ✅ Generated .api files compile without errors
- ✅ API services pass all existing tests
- ✅ Code generation time < 5 seconds per Proto file

**Business Metrics**:
- ✅ Development time reduced by 50% (10 min → 5 min per endpoint)
- ✅ Zero API definition inconsistencies reported
- ✅ Team adoption rate > 80% within 2 weeks

**User Satisfaction**:
- ✅ Developer survey: 80%+ prefer Proto-First
- ✅ Zero regression bugs in production

## Timeline

| Phase | Duration | Deliverables |
|-------|----------|--------------|
| Phase 1: Setup | 8-12h | Plugin structure, Proto options |
| Phase 2: Parsing | 16-20h | HTTP parser, options parser, type converter |
| Phase 3: Generation | 12-16h | Service grouper, template generator |
| Phase 4: Integration | 12-16h | Makefile, integration tests |
| Phase 5: Migration | 12-16h | Pilot modules, documentation |
| **Total** | **60-80h** | **Fully functional Proto-First system** |

## Team Assignment

| Role | Responsibility | Hours |
|------|----------------|-------|
| Backend Developer | Plugin implementation | 50h |
| Backend Developer | Testing and validation | 15h |
| DevOps Engineer | CI/CD integration | 4h |
| Tech Lead | Code review, architecture | 8h |
| Documentation | Migration guide, training | 8h |
| **Total** | | **85h** |

## Dependencies

### External Libraries

```go
// go.mod additions
require (
    google.golang.org/protobuf v1.31.0
    google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5
    github.com/stretchr/testify v1.8.4
)
```

### Build Tools

- protoc v3.19.0+
- Go 1.21+
- goctl v1.6.0+

## Risk Mitigation

### Risk: Generated .api doesn't compile

**Mitigation**:
- Comprehensive integration tests with golden files
- CI/CD validation step before deployment
- Gradual module migration (1-2 at a time)

### Risk: Plugin performance degrades

**Mitigation**:
- Performance benchmarks in CI
- Parallel processing for multiple files
- Caching for type conversions

### Risk: Team resistance

**Mitigation**:
- Hands-on training session
- Migration guide with examples
- Support channel for questions

## Future Enhancements (Out of Scope)

- [ ] Bidirectional sync (Proto ← .api)
- [ ] Custom code generation templates
- [ ] Auto-migration tool for existing .api files
- [ ] IDE plugin for Proto-First development
