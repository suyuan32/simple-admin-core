package model

import (
	"fmt"
	"strings"
)

// Service represents a Proto service
type Service struct {
	Name        string
	Methods     []*Method
	Options     *ServerOptions
	PackageName string
}

// ServerOptions represents @server block options in Go-Zero .api file
type ServerOptions struct {
	JWT        string   // JWT config name (e.g., "Auth")
	Middleware []string // Middleware list (e.g., ["Authority", "RateLimit"])
	Group      string   // Route group (e.g., "user")
	Prefix     string   // Route prefix (e.g., "/api/v1")
}

// Signature generates a unique key for grouping methods by @server config
func (s *ServerOptions) Signature() string {
	return fmt.Sprintf("jwt:%s|mw:%s|group:%s|prefix:%s",
		s.JWT,
		strings.Join(s.Middleware, ","),
		s.Group,
		s.Prefix)
}

// IsEmpty checks if all options are empty
func (s *ServerOptions) IsEmpty() bool {
	return s.JWT == "" && len(s.Middleware) == 0 && s.Group == "" && s.Prefix == ""
}

// HasJWT checks if JWT authentication is required
func (s *ServerOptions) HasJWT() bool {
	return s.JWT != ""
}

// MiddlewareString returns comma-separated middleware list
func (s *ServerOptions) MiddlewareString() string {
	return strings.Join(s.Middleware, ",")
}
