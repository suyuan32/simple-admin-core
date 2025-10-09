package generator

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// Generator generates Go-Zero .api files from Proto definitions
type Generator struct {
	file          *protogen.File
	httpParser    *HTTPParser
	optionsParser *OptionsParser
	typeConverter *TypeConverter
	grouper       *ServiceGrouper
	templateGen   *TemplateGenerator
}

// NewGenerator creates a new Generator instance
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

// Generate generates .api file content from proto file
func (g *Generator) Generate() (string, error) {
	// 1. Parse API info from file-level options
	apiInfo := g.optionsParser.ParseAPIInfo(g.file)

	// 2. Convert all message types (request/response)
	var types []*TypeDefinition
	for _, msg := range g.file.Messages {
		converted := g.typeConverter.ConvertMessage(msg)
		if converted != nil {
			typeDef := &TypeDefinition{
				Name:       converted.Name,
				Definition: g.typeConverter.GenerateTypeDefinition(converted),
			}
			types = append(types, typeDef)
		}
	}

	// 3. Parse all services and methods
	var allMethods []*model.Method
	for _, service := range g.file.Services {
		methods, err := g.parseService(service)
		if err != nil {
			return "", fmt.Errorf("failed to parse service %s: %w", service.Desc.Name(), err)
		}
		allMethods = append(allMethods, methods...)
	}

	// 4. Group methods by @server configuration
	serviceGroups := g.grouper.GroupMethods(allMethods)

	// 5. Prepare template data
	data := &TemplateData{
		APIInfo:       apiInfo,
		Types:         types,
		ServiceGroups: serviceGroups,
	}

	// 6. Validate template data
	if err := g.templateGen.ValidateData(data); err != nil {
		return "", fmt.Errorf("template validation failed: %w", err)
	}

	// 7. Generate .api content
	content, err := g.templateGen.Generate(data)
	if err != nil {
		return "", fmt.Errorf("template generation failed: %w", err)
	}

	return content, nil
}


// parseService parses a Proto service and returns all its methods
func (g *Generator) parseService(service *protogen.Service) ([]*model.Method, error) {
	// Parse service-level options
	serviceOpts := g.optionsParser.ParseServiceOptions(service)

	var methods []*model.Method

	// Parse each method
	for _, method := range service.Methods {
		parsedMethod, err := g.parseMethod(method, serviceOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to parse method %s: %w", method.Desc.Name(), err)
		}
		methods = append(methods, parsedMethod)
	}

	return methods, nil
}

// parseMethod parses a Proto method with HTTP and options annotations
func (g *Generator) parseMethod(method *protogen.Method, serviceOpts *model.ServerOptions) (*model.Method, error) {
	// 1. Parse HTTP rule (google.api.http)
	httpRule, err := g.httpParser.Parse(method)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP rule: %w", err)
	}

	// Skip methods without HTTP annotations
	if httpRule == nil {
		return nil, fmt.Errorf("method %s missing google.api.http annotation", method.Desc.Name())
	}

	// 2. Validate path parameters
	if err := g.httpParser.ValidatePathParams(httpRule, method.Input); err != nil {
		return nil, fmt.Errorf("invalid path parameters: %w", err)
	}

	// 3. Parse method-level options
	methodOpts := g.optionsParser.ParseMethodOptions(method)

	// 4. Merge service and method options
	effectiveOpts := g.optionsParser.MergeOptions(serviceOpts, methodOpts)

	// 5. Create method model
	return &model.Method{
		Name:         string(method.Desc.Name()),
		RequestType:  string(method.Input.Desc.Name()),
		ResponseType: string(method.Output.Desc.Name()),
		HTTPRule:     httpRule,
		Options:      effectiveOpts,
	}, nil
}
