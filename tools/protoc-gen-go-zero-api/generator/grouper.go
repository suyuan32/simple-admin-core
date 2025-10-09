package generator

import (
	"sort"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// ServiceGrouper groups methods by their @server configuration
// Methods with the same @server options are grouped together
type ServiceGrouper struct{}

// NewServiceGrouper creates a new ServiceGrouper instance
func NewServiceGrouper() *ServiceGrouper {
	return &ServiceGrouper{}
}

// ServiceGroup represents a group of methods with the same @server config
type ServiceGroup struct {
	ServerOptions *model.ServerOptions
	Methods       []*model.Method
}

// GroupMethods groups methods by their effective @server options
// Returns a slice of ServiceGroups sorted by priority (protected methods first, then public)
func (g *ServiceGrouper) GroupMethods(methods []*model.Method) []*ServiceGroup {
	if len(methods) == 0 {
		return nil
	}

	// Group methods by their ServerOptions signature
	groupMap := make(map[string]*ServiceGroup)

	for _, method := range methods {
		if method.Options == nil {
			continue
		}

		signature := method.Options.Signature()

		if group, exists := groupMap[signature]; exists {
			group.Methods = append(group.Methods, method)
		} else {
			groupMap[signature] = &ServiceGroup{
				ServerOptions: method.Options,
				Methods:       []*model.Method{method},
			}
		}
	}

	// Convert map to slice
	groups := make([]*ServiceGroup, 0, len(groupMap))
	for _, group := range groupMap {
		groups = append(groups, group)
	}

	// Sort groups for consistent output
	g.sortGroups(groups)

	return groups
}

// sortGroups sorts service groups by priority
// Priority order:
// 1. JWT-protected endpoints (highest priority)
// 2. Number of middleware (more middleware = higher priority)
// 3. Group name (alphabetical)
func (g *ServiceGrouper) sortGroups(groups []*ServiceGroup) {
	sort.Slice(groups, func(i, j int) bool {
		a, b := groups[i].ServerOptions, groups[j].ServerOptions

		// JWT endpoints come first
		aHasJWT := a.JWT != ""
		bHasJWT := b.JWT != ""
		if aHasJWT != bHasJWT {
			return aHasJWT // true (JWT) comes before false (no JWT)
		}

		// More middleware comes first
		if len(a.Middleware) != len(b.Middleware) {
			return len(a.Middleware) > len(b.Middleware)
		}

		// Alphabetical by group name
		return a.Group < b.Group
	})
}

// SplitByServerConfig splits methods into separate groups based on @server configuration
// This is useful when methods within a service need different JWT or middleware settings
func (g *ServiceGrouper) SplitByServerConfig(service *model.Service) []*ServiceGroup {
	return g.GroupMethods(service.Methods)
}

// MergeGroups merges multiple service groups if they have the same @server configuration
// This is useful when processing multiple Proto services that should be combined
func (g *ServiceGrouper) MergeGroups(groups []*ServiceGroup) []*ServiceGroup {
	if len(groups) <= 1 {
		return groups
	}

	groupMap := make(map[string]*ServiceGroup)

	for _, group := range groups {
		signature := group.ServerOptions.Signature()

		if existing, exists := groupMap[signature]; exists {
			existing.Methods = append(existing.Methods, group.Methods...)
		} else {
			// Create a copy to avoid modifying the original
			newGroup := &ServiceGroup{
				ServerOptions: group.ServerOptions,
				Methods:       make([]*model.Method, len(group.Methods)),
			}
			copy(newGroup.Methods, group.Methods)
			groupMap[signature] = newGroup
		}
	}

	// Convert back to slice
	merged := make([]*ServiceGroup, 0, len(groupMap))
	for _, group := range groupMap {
		merged = append(merged, group)
	}

	// Sort for consistent output
	g.sortGroups(merged)

	return merged
}

// GetPublicGroups returns all service groups that don't require JWT
func (g *ServiceGrouper) GetPublicGroups(groups []*ServiceGroup) []*ServiceGroup {
	var public []*ServiceGroup

	for _, group := range groups {
		if !group.ServerOptions.HasJWT() {
			public = append(public, group)
		}
	}

	return public
}

// GetProtectedGroups returns all service groups that require JWT
func (g *ServiceGrouper) GetProtectedGroups(groups []*ServiceGroup) []*ServiceGroup {
	var protected []*ServiceGroup

	for _, group := range groups {
		if group.ServerOptions.HasJWT() {
			protected = append(protected, group)
		}
	}

	return protected
}

// HasMultipleGroups checks if methods need to be split into multiple @server blocks
func (g *ServiceGrouper) HasMultipleGroups(methods []*model.Method) bool {
	groups := g.GroupMethods(methods)
	return len(groups) > 1
}

// GetGroupCount returns the number of unique @server configurations
func (g *ServiceGrouper) GetGroupCount(methods []*model.Method) int {
	groups := g.GroupMethods(methods)
	return len(groups)
}
