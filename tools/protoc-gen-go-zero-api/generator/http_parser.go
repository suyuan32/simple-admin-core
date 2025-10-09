package generator

import (
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// HTTPParser parses google.api.http annotations from Proto methods
type HTTPParser struct{}

// NewHTTPParser creates a new HTTPParser instance
func NewHTTPParser() *HTTPParser {
	return &HTTPParser{}
}

// Parse extracts HTTP rule from a Proto method
// Returns nil if no google.api.http annotation found
func (p *HTTPParser) Parse(method *protogen.Method) (*model.HTTPRule, error) {
	// Get method options
	opts := method.Desc.Options()
	if opts == nil {
		return nil, nil
	}

	// Extract google.api.http extension
	httpRule := proto.GetExtension(opts, annotations.E_Http)
	if httpRule == nil {
		return nil, nil
	}

	rule := httpRule.(*annotations.HttpRule)

	// Parse HTTP method and path from the rule pattern
	httpMethod, path, body, err := p.parsePattern(rule)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP pattern for method %s: %w", method.Desc.Name(), err)
	}

	// Transform path from Proto format to Go-Zero format
	// Example: /user/{id} -> /user/:id
	path = p.transformPath(path)

	return &model.HTTPRule{
		Method: httpMethod,
		Path:   path,
		Body:   body,
	}, nil
}

// parsePattern extracts HTTP method, path, and body from HttpRule
func (p *HTTPParser) parsePattern(rule *annotations.HttpRule) (httpMethod, path, body string, err error) {
	// Get body mapping (defaults to "*" if not specified)
	body = rule.Body
	if body == "" {
		body = "*"
	}

	// Extract HTTP method and path from the pattern union
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return "get", pattern.Get, body, nil
	case *annotations.HttpRule_Post:
		return "post", pattern.Post, body, nil
	case *annotations.HttpRule_Put:
		return "put", pattern.Put, body, nil
	case *annotations.HttpRule_Delete:
		return "delete", pattern.Delete, body, nil
	case *annotations.HttpRule_Patch:
		return "patch", pattern.Patch, body, nil
	case *annotations.HttpRule_Custom:
		// Custom HTTP methods not commonly supported in Go-Zero
		return "", "", "", fmt.Errorf("custom HTTP method not supported: %s", pattern.Custom.Kind)
	default:
		return "", "", "", fmt.Errorf("unknown HTTP pattern type")
	}
}

// transformPath converts Proto path parameters to Go-Zero format
// Example: /user/{id}/posts/{post_id} -> /user/:id/posts/:post_id
func (p *HTTPParser) transformPath(path string) string {
	result := path

	// Replace all {param} with :param
	for {
		start := strings.Index(result, "{")
		if start == -1 {
			break
		}

		end := strings.Index(result[start:], "}")
		if end == -1 {
			// Malformed path, return as-is
			break
		}
		end += start

		// Extract parameter name (including any field path like user.id)
		paramName := result[start+1 : end]

		// Handle field paths in path parameters (e.g., {user.id} -> :user_id)
		// Go-Zero uses underscores instead of dots
		paramName = strings.ReplaceAll(paramName, ".", "_")

		// Replace {param} with :param
		result = result[:start] + ":" + paramName + result[end+1:]
	}

	return result
}

// ParseAdditionalBindings extracts additional HTTP bindings for the method
// Additional bindings allow a single RPC method to have multiple HTTP routes
func (p *HTTPParser) ParseAdditionalBindings(method *protogen.Method) ([]*model.HTTPRule, error) {
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

	// Parse each additional binding
	for _, binding := range rule.AdditionalBindings {
		httpMethod, path, body, err := p.parsePattern(binding)
		if err != nil {
			return nil, fmt.Errorf("failed to parse additional binding: %w", err)
		}

		path = p.transformPath(path)

		rules = append(rules, &model.HTTPRule{
			Method: httpMethod,
			Path:   path,
			Body:   body,
		})
	}

	return rules, nil
}

// ValidatePathParams validates that all path parameters exist in the request message
// This helps catch errors early before code generation
func (p *HTTPParser) ValidatePathParams(httpRule *model.HTTPRule, requestMsg *protogen.Message) error {
	if httpRule == nil || requestMsg == nil {
		return nil
	}

	// Extract path parameters from the path
	pathParams := p.extractPathParams(httpRule.Path)
	if len(pathParams) == 0 {
		return nil
	}

	// Build a map of available fields in the request message
	availableFields := make(map[string]bool)
	for _, field := range requestMsg.Fields {
		fieldName := string(field.Desc.Name())
		availableFields[fieldName] = true

		// Also handle snake_case to camelCase conversion
		// Go-Zero might use different naming conventions
		camelCase := p.snakeToCamel(fieldName)
		if camelCase != fieldName {
			availableFields[camelCase] = true
		}
	}

	// Check each path parameter
	for _, param := range pathParams {
		// Handle nested field paths (e.g., user_id from user.id)
		baseParam := strings.Split(param, "_")[0]

		if !availableFields[param] && !availableFields[baseParam] {
			return fmt.Errorf("path parameter '%s' not found in request message %s", param, requestMsg.Desc.Name())
		}
	}

	return nil
}

// extractPathParams extracts parameter names from a path
// Example: /user/:id/posts/:post_id -> ["id", "post_id"]
func (p *HTTPParser) extractPathParams(path string) []string {
	var params []string
	parts := strings.Split(path, "/")

	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			params = append(params, part[1:])
		}
	}

	return params
}

// snakeToCamel converts snake_case to camelCase
func (p *HTTPParser) snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 1 {
		return s
	}

	result := parts[0]
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			result += strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}

	return result
}
