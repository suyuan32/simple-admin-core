package generator

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// Generator generates Go-Zero .api files from Proto definitions
type Generator struct {
	file *protogen.File
	// Parsers will be added in later tasks
	// httpParser    *HTTPParser
	// optionsParser *OptionsParser
	// typeConverter *TypeConverter
	// grouper       *ServiceGrouper
	// templateGen   *TemplateGenerator
}

// NewGenerator creates a new Generator instance
func NewGenerator(file *protogen.File) *Generator {
	return &Generator{
		file: file,
		// Initialize parsers here in future tasks
	}
}

// Generate generates .api file content from proto file
func (g *Generator) Generate() (string, error) {
	// TODO: This is a stub implementation
	// Full implementation will be done in [PF-011] task

	// For now, return a basic template to allow compilation
	return g.generateStub()
}

// generateStub generates a minimal .api file for testing
func (g *Generator) generateStub() (string, error) {
	// Get package name
	packageName := string(g.file.Desc.Package())

	// Count services
	serviceCount := len(g.file.Services)

	content := fmt.Sprintf(`syntax = "v1"

// Auto-generated from %s
// Package: %s
// Services: %d
// TODO: Full implementation in progress

info(
    title: "API (Generated)"
    desc: "Auto-generated API from Proto"
    author: "protoc-gen-go-zero-api"
    version: "v1.0"
)

// TODO: Import base types
// import "../base.api"

// TODO: Type definitions will be generated here

// TODO: Service definitions will be generated here
`,
		g.file.Desc.Path(),
		packageName,
		serviceCount,
	)

	return content, nil
}

// parseService parses a Proto service (stub for now)
func (g *Generator) parseService(service *protogen.Service) *model.Service {
	// TODO: Implement in [PF-011]
	return &model.Service{
		Name:    string(service.Desc.Name()),
		Methods: []*model.Method{},
		Options: &model.ServerOptions{},
	}
}

// parseMethod parses a Proto method (stub for now)
func (g *Generator) parseMethod(method *protogen.Method, serviceOpts *model.ServerOptions) *model.Method {
	// TODO: Implement in [PF-011]
	return &model.Method{
		Name:         string(method.Desc.Name()),
		RequestType:  string(method.Input.Desc.Name()),
		ResponseType: string(method.Output.Desc.Name()),
		HTTPRule:     nil, // Will be parsed by HTTPParser
		Options:      serviceOpts,
	}
}
