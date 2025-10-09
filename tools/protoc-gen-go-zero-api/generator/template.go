package generator

import (
	"bytes"
	"strings"
	"text/template"

	go_zero "github.com/chimerakang/simple-admin-core/rpc/types/go_zero"
)

// APITemplate is the template for generating .api files
const apiTemplate = `syntax = "v1"
{{- if .APIInfo }}

info(
    title: "{{.APIInfo.Title}}"
    desc: "{{.APIInfo.Desc}}"
    author: "{{.APIInfo.Author}}"
    {{- if .APIInfo.Email }}
    email: "{{.APIInfo.Email}}"
    {{- end }}
    version: "{{.APIInfo.Version}}"
)
{{- end }}

import "../base.api"
{{- if .Types }}

type (
{{- range .Types }}
{{.Definition}}
{{- end }}
)
{{- end }}
{{- range .ServiceGroups }}

{{- if not .ServerOptions.IsEmpty }}
@server(
    {{- if .ServerOptions.JWT }}
    jwt: {{.ServerOptions.JWT}}
    {{- end }}
    {{- if .ServerOptions.Group }}
    group: {{.ServerOptions.Group}}
    {{- end }}
    {{- if .ServerOptions.Middleware }}
    middleware: {{.ServerOptions.MiddlewareString}}
    {{- end }}
    {{- if .ServerOptions.Prefix }}
    prefix: {{.ServerOptions.Prefix}}
    {{- end }}
)
{{- end }}
service Core {
{{- range .Methods }}
    @handler {{.HandlerName}}
    {{.HTTPRule.Method}} {{.HTTPRule.Path}} ({{.RequestType}}) returns ({{.ResponseType}})
{{- end }}
}
{{- end }}
`

// TemplateData contains all data needed to generate .api file
type TemplateData struct {
	APIInfo       *go_zero.ApiInfo
	Types         []*TypeDefinition
	ServiceGroups []*ServiceGroup
}

// TypeDefinition represents a type definition in .api file
type TypeDefinition struct {
	Name       string
	Definition string
}

// TemplateGenerator generates .api file content from template
type TemplateGenerator struct {
	tmpl *template.Template
}

// NewTemplateGenerator creates a new TemplateGenerator instance
func NewTemplateGenerator() *TemplateGenerator {
	// Create template with custom functions
	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tmpl := template.Must(template.New("api").Funcs(funcMap).Parse(apiTemplate))

	return &TemplateGenerator{
		tmpl: tmpl,
	}
}

// Generate generates .api file content from template data
func (g *TemplateGenerator) Generate(data *TemplateData) (string, error) {
	var buf bytes.Buffer

	if err := g.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	// Clean up the output
	content := buf.String()
	content = g.cleanupWhitespace(content)

	return content, nil
}

// cleanupWhitespace removes excessive blank lines and trailing whitespace
func (g *TemplateGenerator) cleanupWhitespace(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	prevBlank := false

	for _, line := range lines {
		// Trim trailing whitespace
		line = strings.TrimRight(line, " \t")

		// Skip consecutive blank lines
		if line == "" {
			if !prevBlank {
				result = append(result, line)
				prevBlank = true
			}
		} else {
			result = append(result, line)
			prevBlank = false
		}
	}

	// Remove leading and trailing blank lines
	for len(result) > 0 && result[0] == "" {
		result = result[1:]
	}
	for len(result) > 0 && result[len(result)-1] == "" {
		result = result[:len(result)-1]
	}

	return strings.Join(result, "\n") + "\n"
}

// GenerateWithDefaults generates .api file with default values if data is missing
func (g *TemplateGenerator) GenerateWithDefaults(data *TemplateData) (string, error) {
	// Provide defaults if missing
	if data == nil {
		data = &TemplateData{}
	}

	if data.APIInfo == nil {
		data.APIInfo = &go_zero.ApiInfo{
			Title:   "API",
			Desc:    "Auto-generated API from Proto",
			Author:  "protoc-gen-go-zero-api",
			Version: "v1.0",
		}
	}

	return g.Generate(data)
}

// ValidateData validates template data before generation
func (g *TemplateGenerator) ValidateData(data *TemplateData) error {
	if data == nil {
		return nil
	}

	// Check for duplicate type names
	typeNames := make(map[string]bool)
	for _, typeDef := range data.Types {
		if typeNames[typeDef.Name] {
			return &ValidationError{
				Message: "duplicate type name: " + typeDef.Name,
			}
		}
		typeNames[typeDef.Name] = true
	}

	// Check for duplicate handler names across all groups
	handlerNames := make(map[string]bool)
	for _, group := range data.ServiceGroups {
		for _, method := range group.Methods {
			handlerName := method.HandlerName()
			if handlerNames[handlerName] {
				return &ValidationError{
					Message: "duplicate handler name: " + handlerName,
				}
			}
			handlerNames[handlerName] = true
		}
	}

	return nil
}

// ValidationError represents a template data validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return "template validation error: " + e.Message
}
