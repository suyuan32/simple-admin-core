package generator

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	go_zero "github.com/chimerakang/simple-admin-core/rpc/types/go_zero"
	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// OptionsParser parses Go-Zero custom options from Proto definitions
type OptionsParser struct{}

// NewOptionsParser creates a new OptionsParser instance
func NewOptionsParser() *OptionsParser {
	return &OptionsParser{}
}

// ParseServiceOptions extracts Go-Zero options from a Proto service
func (p *OptionsParser) ParseServiceOptions(service *protogen.Service) *model.ServerOptions {
	opts := service.Desc.Options()
	if opts == nil {
		return p.getDefaultServerOptions(service)
	}

	serverOpts := &model.ServerOptions{}

	// Extract JWT option
	if ext := proto.GetExtension(opts, go_zero.E_Jwt); ext != nil {
		if jwt, ok := ext.(string); ok {
			serverOpts.JWT = jwt
		}
	}

	// Extract Middleware option
	if ext := proto.GetExtension(opts, go_zero.E_Middleware); ext != nil {
		if middleware, ok := ext.(string); ok && middleware != "" {
			serverOpts.Middleware = p.parseMiddlewareList(middleware)
		}
	}

	// Extract Group option
	if ext := proto.GetExtension(opts, go_zero.E_Group); ext != nil {
		if group, ok := ext.(string); ok {
			serverOpts.Group = group
		}
	}

	// Extract Prefix option
	if ext := proto.GetExtension(opts, go_zero.E_Prefix); ext != nil {
		if prefix, ok := ext.(string); ok {
			serverOpts.Prefix = prefix
		}
	}

	// Set default group if not specified
	if serverOpts.Group == "" {
		serverOpts.Group = strings.ToLower(string(service.Desc.Name()))
	}

	return serverOpts
}

// ParseMethodOptions extracts Go-Zero options from a Proto method
func (p *OptionsParser) ParseMethodOptions(method *protogen.Method) *model.MethodOptions {
	opts := method.Desc.Options()
	if opts == nil {
		return &model.MethodOptions{}
	}

	methodOpts := &model.MethodOptions{}

	// Extract Public option
	if ext := proto.GetExtension(opts, go_zero.E_Public); ext != nil {
		if public, ok := ext.(bool); ok {
			methodOpts.Public = public
		}
	}

	// Extract Method-specific Middleware option
	if ext := proto.GetExtension(opts, go_zero.E_MethodMiddleware); ext != nil {
		if middleware, ok := ext.(string); ok && middleware != "" {
			methodOpts.Middleware = p.parseMiddlewareList(middleware)
		}
	}

	return methodOpts
}

// ParseAPIInfo extracts API metadata from file-level options
func (p *OptionsParser) ParseAPIInfo(file *protogen.File) *go_zero.ApiInfo {
	opts := file.Desc.Options()
	if opts == nil {
		return nil
	}

	if ext := proto.GetExtension(opts, go_zero.E_ApiInfo); ext != nil {
		if apiInfo, ok := ext.(*go_zero.ApiInfo); ok {
			return apiInfo
		}
	}

	return nil
}

// MergeOptions merges service-level and method-level options
// Method-level options take precedence over service-level options
func (p *OptionsParser) MergeOptions(serviceOpts *model.ServerOptions, methodOpts *model.MethodOptions) *model.ServerOptions {
	if serviceOpts == nil {
		serviceOpts = &model.ServerOptions{}
	}
	if methodOpts == nil {
		methodOpts = &model.MethodOptions{}
	}

	// Create a copy of service options
	merged := &model.ServerOptions{
		JWT:        serviceOpts.JWT,
		Middleware: make([]string, len(serviceOpts.Middleware)),
		Group:      serviceOpts.Group,
		Prefix:     serviceOpts.Prefix,
	}
	copy(merged.Middleware, serviceOpts.Middleware)

	// Apply method-level overrides

	// If method is public, remove JWT requirement
	if methodOpts.Public {
		merged.JWT = ""
	}

	// If method has specific middleware, override service middleware
	if len(methodOpts.Middleware) > 0 {
		merged.Middleware = make([]string, len(methodOpts.Middleware))
		copy(merged.Middleware, methodOpts.Middleware)
	}

	return merged
}

// parseMiddlewareList parses comma-separated middleware string into a slice
// Trims whitespace from each middleware name
func (p *OptionsParser) parseMiddlewareList(middleware string) []string {
	if middleware == "" {
		return nil
	}

	parts := strings.Split(middleware, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

// getDefaultServerOptions returns default server options when none are specified
func (p *OptionsParser) getDefaultServerOptions(service *protogen.Service) *model.ServerOptions {
	return &model.ServerOptions{
		Group: strings.ToLower(string(service.Desc.Name())),
	}
}

// GetEffectiveOptions computes the effective options for a method
// This is a convenience method that combines ParseServiceOptions, ParseMethodOptions, and MergeOptions
func (p *OptionsParser) GetEffectiveOptions(service *protogen.Service, method *protogen.Method) *model.ServerOptions {
	serviceOpts := p.ParseServiceOptions(service)
	methodOpts := p.ParseMethodOptions(method)
	return p.MergeOptions(serviceOpts, methodOpts)
}

// HasJWT checks if the effective options require JWT authentication
func (p *OptionsParser) HasJWT(serviceOpts *model.ServerOptions, methodOpts *model.MethodOptions) bool {
	// If method is explicitly public, no JWT required
	if methodOpts != nil && methodOpts.Public {
		return false
	}

	// Otherwise, check service-level JWT requirement
	return serviceOpts != nil && serviceOpts.JWT != ""
}

// GetMiddleware returns the effective middleware list for a method
func (p *OptionsParser) GetMiddleware(serviceOpts *model.ServerOptions, methodOpts *model.MethodOptions) []string {
	// Method-level middleware overrides service-level
	if methodOpts != nil && len(methodOpts.Middleware) > 0 {
		return methodOpts.Middleware
	}

	// Fall back to service-level middleware
	if serviceOpts != nil {
		return serviceOpts.Middleware
	}

	return nil
}
