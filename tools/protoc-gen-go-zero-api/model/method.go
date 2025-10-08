package model

import (
	"strings"
)

// Method represents a Proto RPC method
type Method struct {
	Name         string
	RequestType  string
	ResponseType string
	HTTPRule     *HTTPRule
	Options      *ServerOptions // Effective options (merged service + method)
}

// HTTPRule represents google.api.http annotation
type HTTPRule struct {
	Method string // HTTP method: GET, POST, PUT, DELETE, PATCH
	Path   string // URL path (e.g., "/user/:id")
	Body   string // Body mapping (e.g., "*" or "user")
}

// MethodOptions represents method-level options
type MethodOptions struct {
	Public     bool     // No JWT required (override service-level JWT)
	Middleware []string // Method-specific middleware (override service-level)
}

// HandlerName converts method name to Go-Zero handler name
// Example: CreateUser -> createUser
func (m *Method) HandlerName() string {
	if len(m.Name) == 0 {
		return ""
	}
	// Convert first letter to lowercase
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

// IsValid checks if method has valid HTTP rule
func (m *Method) IsValid() bool {
	return m.HTTPRule != nil && m.HTTPRule.Method != "" && m.HTTPRule.Path != ""
}
