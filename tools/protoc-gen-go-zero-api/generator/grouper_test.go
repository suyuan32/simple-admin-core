// Copyright 2025 simple-admin-core. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.

package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/model"
)

// Test Suite: ServiceGrouper Initialization

func TestNewServiceGrouper(t *testing.T) {
	grouper := NewServiceGrouper()

	require.NotNil(t, grouper)
	assert.IsType(t, &ServiceGrouper{}, grouper)
}

// Test Suite: Basic Grouping Logic

func TestGroupMethods_EmptyMethods(t *testing.T) {
	grouper := NewServiceGrouper()

	result := grouper.GroupMethods(nil)

	assert.Nil(t, result)
}

func TestGroupMethods_EmptySlice(t *testing.T) {
	grouper := NewServiceGrouper()

	result := grouper.GroupMethods([]*model.Method{})

	assert.Nil(t, result)
}

func TestGroupMethods_SingleMethod(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 1)
	assert.Len(t, result[0].Methods, 1)
	assert.Equal(t, "GetUser", result[0].Methods[0].Name)
	assert.Equal(t, "Auth", result[0].ServerOptions.JWT)
}

func TestGroupMethods_IdenticalServerOptions(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "CreateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "UpdateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 1, "Methods with identical server options should be in one group")
	assert.Len(t, result[0].Methods, 3)
	assert.Equal(t, "Auth", result[0].ServerOptions.JWT)
	assert.Equal(t, []string{"Authority"}, result[0].ServerOptions.Middleware)
}

func TestGroupMethods_DifferentServerOptions(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "Login",
			Options: &model.ServerOptions{
				JWT:    "",
				Group:  "auth",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Methods with different server options should be in separate groups")
}

func TestGroupMethods_SkipsMethodsWithNilOptions(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name:    "GetUser",
			Options: nil, // Should be skipped
		},
		{
			Name: "CreateUser",
			Options: &model.ServerOptions{
				JWT: "Auth",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 1, "Only method with non-nil options should be grouped")
	assert.Equal(t, "CreateUser", result[0].Methods[0].Name)
}

// Test Suite: JWT Grouping

func TestGroupMethods_JWTVsNoJWT(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "Login",
			Options: &model.ServerOptions{
				JWT:    "",
				Group:  "auth",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "Register",
			Options: &model.ServerOptions{
				JWT:    "",
				Group:  "auth",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "GetProfile",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "UpdateProfile",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "JWT and non-JWT methods should be in separate groups")

	// First group should be JWT-protected (higher priority)
	assert.Equal(t, "Auth", result[0].ServerOptions.JWT)
	assert.Len(t, result[0].Methods, 2)

	// Second group should be public
	assert.Empty(t, result[1].ServerOptions.JWT)
	assert.Len(t, result[1].Methods, 2)
}

func TestGroupMethods_PublicMethodOverride(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "Login",
			Options: &model.ServerOptions{
				JWT:    "", // Public method (no JWT)
				Group:  "auth",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:    "Auth", // Protected method
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Public and protected methods should be separate")

	// Verify JWT-protected method comes first (due to sorting)
	hasJWTGroup := false
	noJWTGroup := false

	for _, group := range result {
		if group.ServerOptions.HasJWT() {
			hasJWTGroup = true
			assert.Equal(t, "Auth", group.ServerOptions.JWT)
		} else {
			noJWTGroup = true
			assert.Empty(t, group.ServerOptions.JWT)
		}
	}

	assert.True(t, hasJWTGroup, "Should have JWT-protected group")
	assert.True(t, noJWTGroup, "Should have public group")
}

// Test Suite: Middleware Grouping

func TestGroupMethods_DifferentMiddleware(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "Login",
			Options: &model.ServerOptions{
				JWT:        "",
				Middleware: []string{"RateLimit"},
				Group:      "auth",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Different middleware should create separate groups")
}

func TestGroupMethods_SameMiddlewareDifferentOrder(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "CreateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Audit"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "UpdateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Audit", "Authority"}, // Different order
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Middleware order matters - should be separate groups")
}

func TestGroupMethods_SameMiddlewareSameOrder(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "CreateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Audit"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "UpdateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Audit"}, // Same order
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 1, "Same middleware in same order should be in one group")
	assert.Len(t, result[0].Methods, 2)
}

func TestGroupMethods_NoMiddlewareVsWithMiddleware(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "UpdateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Methods with and without middleware should be separate")
}

func TestGroupMethods_MultipleMiddleware(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "CreateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Audit", "Log"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "UpdateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Audit", "Log"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 1, "Same multiple middleware should group together")
	assert.Len(t, result[0].Methods, 2)
	assert.Equal(t, []string{"Authority", "Audit", "Log"}, result[0].ServerOptions.Middleware)
}

// Test Suite: Group/Prefix Grouping

func TestGroupMethods_DifferentGroup(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "GetRole",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "role", // Different group
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Different groups should create separate groups")
}

func TestGroupMethods_DifferentPrefix(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "AdminGetUser",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1/admin", // Different prefix
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Different prefixes should create separate groups")
}

func TestGroupMethods_EmptyGroupVsPopulated(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "UpdateUser",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user",
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Empty group vs populated group should be separate")
}

// Test Suite: Sorting Priority

func TestSortGroups_JWTComesFirst(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "Login",
			Options: &model.ServerOptions{
				JWT:   "",
				Group: "auth",
			},
		},
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2)
	// JWT-protected should come first
	assert.NotEmpty(t, result[0].ServerOptions.JWT, "JWT group should be first")
	assert.Empty(t, result[1].ServerOptions.JWT, "Non-JWT group should be second")
}

func TestSortGroups_MoreMiddlewareFirst(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "a",
			},
		},
		{
			Name: "UpdateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Audit", "Log"},
				Group:      "b",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2)
	// More middleware should come first
	assert.Len(t, result[0].ServerOptions.Middleware, 3, "Group with more middleware should be first")
	assert.Len(t, result[1].ServerOptions.Middleware, 1, "Group with less middleware should be second")
}

func TestSortGroups_AlphabeticalByGroup(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
		},
		{
			Name: "GetRole",
			Options: &model.ServerOptions{
				JWT:   "Auth",
				Group: "role",
			},
		},
		{
			Name: "GetMenu",
			Options: &model.ServerOptions{
				JWT:   "Auth",
				Group: "menu",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 3)
	// Should be alphabetical
	assert.Equal(t, "menu", result[0].ServerOptions.Group)
	assert.Equal(t, "role", result[1].ServerOptions.Group)
	assert.Equal(t, "user", result[2].ServerOptions.Group)
}

func TestSortGroups_ComplexPriority(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		// Group 1: No JWT, no middleware
		{
			Name: "Login",
			Options: &model.ServerOptions{
				JWT:   "",
				Group: "auth",
			},
		},
		// Group 2: JWT, 3 middleware
		{
			Name: "AdminUpdateUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Admin", "Audit"},
				Group:      "admin",
			},
		},
		// Group 3: JWT, 1 middleware
		{
			Name: "GetUser",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "user",
			},
		},
		// Group 4: JWT, no middleware
		{
			Name: "GetRole",
			Options: &model.ServerOptions{
				JWT:   "Auth",
				Group: "role",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 4)

	// Priority order:
	// 1. JWT with 3 middleware (admin)
	assert.Equal(t, "Auth", result[0].ServerOptions.JWT)
	assert.Len(t, result[0].ServerOptions.Middleware, 3)

	// 2. JWT with 1 middleware (user)
	assert.Equal(t, "Auth", result[1].ServerOptions.JWT)
	assert.Len(t, result[1].ServerOptions.Middleware, 1)

	// 3. JWT with no middleware (role)
	assert.Equal(t, "Auth", result[2].ServerOptions.JWT)
	assert.Empty(t, result[2].ServerOptions.Middleware)

	// 4. No JWT (auth)
	assert.Empty(t, result[3].ServerOptions.JWT)
}

// Test Suite: SplitByServerConfig

func TestSplitByServerConfig_SingleGroup(t *testing.T) {
	grouper := NewServiceGrouper()

	service := &model.Service{
		Name: "UserService",
		Methods: []*model.Method{
			{
				Name: "GetUser",
				Options: &model.ServerOptions{
					JWT:   "Auth",
					Group: "user",
				},
			},
			{
				Name: "CreateUser",
				Options: &model.ServerOptions{
					JWT:   "Auth",
					Group: "user",
				},
			},
		},
	}

	result := grouper.SplitByServerConfig(service)

	require.Len(t, result, 1)
	assert.Len(t, result[0].Methods, 2)
}

func TestSplitByServerConfig_MultipleGroups(t *testing.T) {
	grouper := NewServiceGrouper()

	service := &model.Service{
		Name: "UserService",
		Methods: []*model.Method{
			{
				Name: "Login",
				Options: &model.ServerOptions{
					JWT:   "",
					Group: "auth",
				},
			},
			{
				Name: "GetUser",
				Options: &model.ServerOptions{
					JWT:   "Auth",
					Group: "user",
				},
			},
		},
	}

	result := grouper.SplitByServerConfig(service)

	require.Len(t, result, 2)
}

func TestSplitByServerConfig_EmptyService(t *testing.T) {
	grouper := NewServiceGrouper()

	service := &model.Service{
		Name:    "EmptyService",
		Methods: []*model.Method{},
	}

	result := grouper.SplitByServerConfig(service)

	assert.Nil(t, result)
}

// Test Suite: MergeGroups

func TestMergeGroups_EmptyGroups(t *testing.T) {
	grouper := NewServiceGrouper()

	result := grouper.MergeGroups(nil)

	assert.Nil(t, result)
}

func TestMergeGroups_SingleGroup(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
			Methods: []*model.Method{
				{Name: "GetUser"},
			},
		},
	}

	result := grouper.MergeGroups(groups)

	require.Len(t, result, 1)
	assert.Len(t, result[0].Methods, 1)
}

func TestMergeGroups_IdenticalServerOptions(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
			Methods: []*model.Method{
				{Name: "GetUser"},
			},
		},
		{
			ServerOptions: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
			Methods: []*model.Method{
				{Name: "CreateUser"},
			},
		},
	}

	result := grouper.MergeGroups(groups)

	require.Len(t, result, 1, "Groups with identical server options should merge")
	assert.Len(t, result[0].Methods, 2, "Methods should be combined")
}

func TestMergeGroups_DifferentServerOptions(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
			Methods: []*model.Method{
				{Name: "GetUser"},
			},
		},
		{
			ServerOptions: &model.ServerOptions{
				JWT:   "",
				Group: "auth",
			},
			Methods: []*model.Method{
				{Name: "Login"},
			},
		},
	}

	result := grouper.MergeGroups(groups)

	require.Len(t, result, 2, "Groups with different server options should stay separate")
}

func TestMergeGroups_DoesNotModifyOriginal(t *testing.T) {
	grouper := NewServiceGrouper()

	originalGroup1 := &ServiceGroup{
		ServerOptions: &model.ServerOptions{
			JWT:   "Auth",
			Group: "user",
		},
		Methods: []*model.Method{
			{Name: "GetUser"},
		},
	}

	originalGroup2 := &ServiceGroup{
		ServerOptions: &model.ServerOptions{
			JWT:   "Auth",
			Group: "user",
		},
		Methods: []*model.Method{
			{Name: "CreateUser"},
		},
	}

	groups := []*ServiceGroup{originalGroup1, originalGroup2}

	result := grouper.MergeGroups(groups)

	require.Len(t, result, 1)

	// Verify originals are not modified
	assert.Len(t, originalGroup1.Methods, 1, "Original group 1 should not be modified")
	assert.Len(t, originalGroup2.Methods, 1, "Original group 2 should not be modified")

	// Verify result is a copy by modifying it
	originalLength := len(result[0].Methods)
	result[0].Methods = append(result[0].Methods, &model.Method{Name: "NewMethod"})

	// Check that originals still have their original length
	assert.Len(t, originalGroup1.Methods, 1, "Modifying result should not affect original 1")
	assert.Len(t, originalGroup2.Methods, 1, "Modifying result should not affect original 2")
	assert.Equal(t, originalLength+1, len(result[0].Methods), "Result should have new method")
}

func TestMergeGroups_MultipleGroupsMerge(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth", Group: "user"},
			Methods:       []*model.Method{{Name: "GetUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth", Group: "user"},
			Methods:       []*model.Method{{Name: "CreateUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth", Group: "user"},
			Methods:       []*model.Method{{Name: "UpdateUser"}},
		},
	}

	result := grouper.MergeGroups(groups)

	require.Len(t, result, 1, "All three groups should merge")
	assert.Len(t, result[0].Methods, 3, "All methods should be combined")
}

func TestMergeGroups_PartialMerge(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth", Group: "user"},
			Methods:       []*model.Method{{Name: "GetUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth", Group: "user"},
			Methods:       []*model.Method{{Name: "CreateUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "", Group: "auth"},
			Methods:       []*model.Method{{Name: "Login"}},
		},
	}

	result := grouper.MergeGroups(groups)

	require.Len(t, result, 2, "Should merge 2 groups, keep 1 separate")

	// Count total methods
	totalMethods := 0
	for _, group := range result {
		totalMethods += len(group.Methods)
	}
	assert.Equal(t, 3, totalMethods, "All methods should be preserved")
}

// Test Suite: GetPublicGroups

func TestGetPublicGroups_NoPublicGroups(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth"},
			Methods:       []*model.Method{{Name: "GetUser"}},
		},
	}

	result := grouper.GetPublicGroups(groups)

	assert.Empty(t, result, "No public groups should be returned")
}

func TestGetPublicGroups_AllPublic(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: ""},
			Methods:       []*model.Method{{Name: "Login"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: ""},
			Methods:       []*model.Method{{Name: "Register"}},
		},
	}

	result := grouper.GetPublicGroups(groups)

	assert.Len(t, result, 2, "All groups should be public")
}

func TestGetPublicGroups_Mixed(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: ""},
			Methods:       []*model.Method{{Name: "Login"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth"},
			Methods:       []*model.Method{{Name: "GetUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: ""},
			Methods:       []*model.Method{{Name: "Register"}},
		},
	}

	result := grouper.GetPublicGroups(groups)

	require.Len(t, result, 2, "Should return only public groups")
	assert.Empty(t, result[0].ServerOptions.JWT)
	assert.Empty(t, result[1].ServerOptions.JWT)
}

// Test Suite: GetProtectedGroups

func TestGetProtectedGroups_NoProtectedGroups(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: ""},
			Methods:       []*model.Method{{Name: "Login"}},
		},
	}

	result := grouper.GetProtectedGroups(groups)

	assert.Empty(t, result, "No protected groups should be returned")
}

func TestGetProtectedGroups_AllProtected(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth"},
			Methods:       []*model.Method{{Name: "GetUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth"},
			Methods:       []*model.Method{{Name: "GetRole"}},
		},
	}

	result := grouper.GetProtectedGroups(groups)

	assert.Len(t, result, 2, "All groups should be protected")
}

func TestGetProtectedGroups_Mixed(t *testing.T) {
	grouper := NewServiceGrouper()

	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth"},
			Methods:       []*model.Method{{Name: "GetUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: ""},
			Methods:       []*model.Method{{Name: "Login"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth"},
			Methods:       []*model.Method{{Name: "UpdateUser"}},
		},
	}

	result := grouper.GetProtectedGroups(groups)

	require.Len(t, result, 2, "Should return only protected groups")
	assert.NotEmpty(t, result[0].ServerOptions.JWT)
	assert.NotEmpty(t, result[1].ServerOptions.JWT)
}

// Test Suite: HasMultipleGroups

func TestHasMultipleGroups_SingleGroup(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name:    "GetUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
		{
			Name:    "CreateUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
	}

	result := grouper.HasMultipleGroups(methods)

	assert.False(t, result, "Should return false for single group")
}

func TestHasMultipleGroups_MultipleGroups(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name:    "Login",
			Options: &model.ServerOptions{JWT: "", Group: "auth"},
		},
		{
			Name:    "GetUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
	}

	result := grouper.HasMultipleGroups(methods)

	assert.True(t, result, "Should return true for multiple groups")
}

func TestHasMultipleGroups_EmptyMethods(t *testing.T) {
	grouper := NewServiceGrouper()

	result := grouper.HasMultipleGroups(nil)

	assert.False(t, result, "Should return false for empty methods")
}

// Test Suite: GetGroupCount

func TestGetGroupCount_Zero(t *testing.T) {
	grouper := NewServiceGrouper()

	result := grouper.GetGroupCount(nil)

	assert.Equal(t, 0, result)
}

func TestGetGroupCount_One(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name:    "GetUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
	}

	result := grouper.GetGroupCount(methods)

	assert.Equal(t, 1, result)
}

func TestGetGroupCount_Multiple(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name:    "Login",
			Options: &model.ServerOptions{JWT: "", Group: "auth"},
		},
		{
			Name:    "GetUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
		{
			Name:    "GetRole",
			Options: &model.ServerOptions{JWT: "Auth", Group: "role"},
		},
	}

	result := grouper.GetGroupCount(methods)

	assert.Equal(t, 3, result)
}

// Test Suite: Complex Real-World Scenarios

func TestComplexScenario_MicroserviceAPIs(t *testing.T) {
	grouper := NewServiceGrouper()

	// Simulate a typical microservice with mixed public and protected endpoints
	methods := []*model.Method{
		// Public authentication endpoints
		{
			Name: "Login",
			Options: &model.ServerOptions{
				JWT:        "",
				Middleware: []string{"RateLimit", "Captcha"},
				Group:      "auth",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "Register",
			Options: &model.ServerOptions{
				JWT:        "",
				Middleware: []string{"RateLimit", "Captcha"},
				Group:      "auth",
				Prefix:     "/api/v1",
			},
		},
		// Protected user endpoints
		{
			Name: "GetProfile",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		{
			Name: "UpdateProfile",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Audit"},
				Group:      "user",
				Prefix:     "/api/v1",
			},
		},
		// Admin endpoints
		{
			Name: "GetAllUsers",
			Options: &model.ServerOptions{
				JWT:        "Auth",
				Middleware: []string{"Authority", "Admin"},
				Group:      "admin",
				Prefix:     "/api/v1/admin",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 4, "Should create 4 distinct groups")

	// Verify groups are properly sorted (JWT first, then by middleware count)
	assert.True(t, result[0].ServerOptions.HasJWT(), "First group should be JWT-protected")
	assert.True(t, result[1].ServerOptions.HasJWT(), "Second group should be JWT-protected")
}

func TestComplexScenario_MultiServiceMerge(t *testing.T) {
	grouper := NewServiceGrouper()

	// Simulate multiple services that can be merged
	service1Groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
			Methods: []*model.Method{
				{Name: "GetUser"},
			},
		},
	}

	service2Groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{
				JWT:   "Auth",
				Group: "user",
			},
			Methods: []*model.Method{
				{Name: "CreateUser"},
			},
		},
	}

	allGroups := append(service1Groups, service2Groups...)
	result := grouper.MergeGroups(allGroups)

	require.Len(t, result, 1, "Should merge into single group")
	assert.Len(t, result[0].Methods, 2, "Should contain all methods")
}

// Test Suite: Edge Cases

func TestEdgeCase_AllNilOptions(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{Name: "Method1", Options: nil},
		{Name: "Method2", Options: nil},
		{Name: "Method3", Options: nil},
	}

	result := grouper.GroupMethods(methods)

	assert.Empty(t, result, "All nil options should result in empty groups")
}

func TestEdgeCase_GrouperEmptyServerOptions(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name:    "Method1",
			Options: &model.ServerOptions{},
		},
		{
			Name:    "Method2",
			Options: &model.ServerOptions{},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 1, "Empty server options should group together")
	assert.Len(t, result[0].Methods, 2)
}

func TestEdgeCase_GrouperVeryLongMiddlewareList(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "ComplexMethod",
			Options: &model.ServerOptions{
				JWT: "Auth",
				Middleware: []string{
					"MW1", "MW2", "MW3", "MW4", "MW5",
					"MW6", "MW7", "MW8", "MW9", "MW10",
				},
				Group:  "complex",
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 1)
	assert.Len(t, result[0].ServerOptions.Middleware, 10)
}

func TestEdgeCase_SpecialCharactersInGroup(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name: "Method1",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user-admin",
				Prefix: "/api/v1",
			},
		},
		{
			Name: "Method2",
			Options: &model.ServerOptions{
				JWT:    "Auth",
				Group:  "user_admin",
				Prefix: "/api/v1",
			},
		},
	}

	result := grouper.GroupMethods(methods)

	require.Len(t, result, 2, "Different special characters should create different groups")
}

// Test Suite: Performance and Consistency

func TestConsistency_MultipleCalls(t *testing.T) {
	grouper := NewServiceGrouper()

	methods := []*model.Method{
		{
			Name:    "Login",
			Options: &model.ServerOptions{JWT: "", Group: "auth"},
		},
		{
			Name:    "GetUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
	}

	// Call multiple times
	result1 := grouper.GroupMethods(methods)
	result2 := grouper.GroupMethods(methods)
	result3 := grouper.GroupMethods(methods)

	// Results should be consistent
	assert.Len(t, result1, len(result2))
	assert.Len(t, result2, len(result3))

	// Group ordering should be consistent
	for i := range result1 {
		assert.Equal(t, result1[i].ServerOptions.Signature(), result2[i].ServerOptions.Signature())
		assert.Equal(t, result2[i].ServerOptions.Signature(), result3[i].ServerOptions.Signature())
	}
}

// Benchmark Tests

func BenchmarkGroupMethods_SmallSet(b *testing.B) {
	grouper := NewServiceGrouper()
	methods := []*model.Method{
		{
			Name:    "GetUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
		{
			Name:    "CreateUser",
			Options: &model.ServerOptions{JWT: "Auth", Group: "user"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grouper.GroupMethods(methods)
	}
}

func BenchmarkGroupMethods_LargeSet(b *testing.B) {
	grouper := NewServiceGrouper()
	methods := make([]*model.Method, 100)
	for i := 0; i < 100; i++ {
		methods[i] = &model.Method{
			Name:    "Method",
			Options: &model.ServerOptions{JWT: "Auth", Group: "group"},
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grouper.GroupMethods(methods)
	}
}

func BenchmarkMergeGroups(b *testing.B) {
	grouper := NewServiceGrouper()
	groups := []*ServiceGroup{
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth", Group: "user"},
			Methods:       []*model.Method{{Name: "GetUser"}},
		},
		{
			ServerOptions: &model.ServerOptions{JWT: "Auth", Group: "user"},
			Methods:       []*model.Method{{Name: "CreateUser"}},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grouper.MergeGroups(groups)
	}
}
