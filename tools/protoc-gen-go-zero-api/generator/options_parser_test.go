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

// Test Suite: Middleware Parsing

func TestParseMiddlewareList_SingleMiddleware(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("Authority")

	assert.Equal(t, []string{"Authority"}, result)
}

func TestParseMiddlewareList_MultipleMiddleware(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("Authority,RateLimit,Audit")

	assert.Equal(t, []string{"Authority", "RateLimit", "Audit"}, result)
}

func TestParseMiddlewareList_WithSpaces(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("Authority , RateLimit , Audit")

	assert.Equal(t, []string{"Authority", "RateLimit", "Audit"}, result)
}

func TestParseMiddlewareList_WithExtraCommas(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("Authority,,RateLimit,,,Audit")

	assert.Equal(t, []string{"Authority", "RateLimit", "Audit"}, result)
}

func TestParseMiddlewareList_EmptyString(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("")

	assert.Nil(t, result)
}

func TestParseMiddlewareList_OnlySpacesAndCommas(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("  ,  ,  ")

	assert.Empty(t, result)
}

func TestParseMiddlewareList_TrailingComma(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("Authority,RateLimit,")

	assert.Equal(t, []string{"Authority", "RateLimit"}, result)
}

func TestParseMiddlewareList_LeadingComma(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList(",Authority,RateLimit")

	assert.Equal(t, []string{"Authority", "RateLimit"}, result)
}

func TestParseMiddlewareList_OnlyCommas(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList(",,,")

	assert.Empty(t, result)
}

func TestParseMiddlewareList_SingleWithSpaces(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("  Authority  ")

	assert.Equal(t, []string{"Authority"}, result)
}

func TestParseMiddlewareList_MixedWhitespace(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("Authority,  ,RateLimit,  Audit  , ,Log")

	assert.Equal(t, []string{"Authority", "RateLimit", "Audit", "Log"}, result)
}

// Test Suite: Options Merging Logic

func TestMergeOptions_ServiceOnly(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority", "RateLimit"},
		Group:      "user",
		Prefix:     "/api/v1",
	}
	methodOpts := &model.MethodOptions{}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Equal(t, "Auth", result.JWT)
	assert.Equal(t, []string{"Authority", "RateLimit"}, result.Middleware)
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestMergeOptions_PublicMethodOverridesJWT(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority"},
		Group:      "user",
		Prefix:     "/api/v1",
	}
	methodOpts := &model.MethodOptions{
		Public: true,
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT, "Public method should remove JWT requirement")
	assert.Equal(t, []string{"Authority"}, result.Middleware)
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestMergeOptions_MethodMiddlewareOverridesService(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority", "RateLimit"},
		Group:      "user",
		Prefix:     "/api/v1",
	}
	methodOpts := &model.MethodOptions{
		Middleware: []string{"Audit", "Log"},
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Equal(t, "Auth", result.JWT)
	assert.Equal(t, []string{"Audit", "Log"}, result.Middleware, "Method middleware should override service middleware")
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestMergeOptions_PublicWithCustomMiddleware(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority"},
		Group:      "user",
		Prefix:     "/api/v1",
	}
	methodOpts := &model.MethodOptions{
		Public:     true,
		Middleware: []string{"RateLimit", "Captcha"},
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT, "Public method removes JWT")
	assert.Equal(t, []string{"RateLimit", "Captcha"}, result.Middleware, "Method middleware overrides service")
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestMergeOptions_NilServiceOptions(t *testing.T) {
	parser := NewOptionsParser()
	methodOpts := &model.MethodOptions{
		Public:     true,
		Middleware: []string{"RateLimit"},
	}

	result := parser.MergeOptions(nil, methodOpts)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT)
	assert.Equal(t, []string{"RateLimit"}, result.Middleware)
	assert.Empty(t, result.Group)
	assert.Empty(t, result.Prefix)
}

func TestMergeOptions_NilMethodOptions(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority"},
		Group:      "user",
		Prefix:     "/api/v1",
	}

	result := parser.MergeOptions(serviceOpts, nil)

	require.NotNil(t, result)
	assert.Equal(t, "Auth", result.JWT)
	assert.Equal(t, []string{"Authority"}, result.Middleware)
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestMergeOptions_BothNil(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.MergeOptions(nil, nil)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT)
	assert.Empty(t, result.Middleware)
	assert.Empty(t, result.Group)
	assert.Empty(t, result.Prefix)
}

func TestMergeOptions_EmptyMiddlewareSlices(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{},
		Group:      "user",
		Prefix:     "/api/v1",
	}
	methodOpts := &model.MethodOptions{
		Middleware: []string{},
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Equal(t, "Auth", result.JWT)
	assert.Empty(t, result.Middleware, "Empty method middleware should not override service middleware")
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestMergeOptions_DoesNotModifyOriginal(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority", "RateLimit"},
		Group:      "user",
		Prefix:     "/api/v1",
	}
	originalJWT := serviceOpts.JWT
	originalMiddleware := serviceOpts.Middleware
	originalGroup := serviceOpts.Group
	originalPrefix := serviceOpts.Prefix

	methodOpts := &model.MethodOptions{
		Public:     true,
		Middleware: []string{"Audit"},
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	// Verify original service options are not modified
	assert.Equal(t, originalJWT, serviceOpts.JWT, "Original service JWT should not change")
	assert.Equal(t, originalMiddleware, serviceOpts.Middleware, "Original service middleware should not change")
	assert.Equal(t, originalGroup, serviceOpts.Group, "Original service Group should not change")
	assert.Equal(t, originalPrefix, serviceOpts.Prefix, "Original service Prefix should not change")

	// Verify result is different
	assert.Empty(t, result.JWT, "Result should have empty JWT")
	assert.Equal(t, []string{"Audit"}, result.Middleware, "Result should have method middleware")
}

func TestMergeOptions_MiddlewareArrayIsolation(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		Middleware: []string{"Authority", "RateLimit"},
	}
	methodOpts := &model.MethodOptions{}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	// Modify the result's middleware
	result.Middleware[0] = "Modified"

	// Verify original service middleware is not affected
	assert.Equal(t, "Authority", serviceOpts.Middleware[0], "Original middleware should not be affected")
}

// Test Suite: Convenience Methods

func TestHasJWT_ServiceRequiresJWT(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT: "Auth",
	}
	methodOpts := &model.MethodOptions{
		Public: false,
	}

	result := parser.HasJWT(serviceOpts, methodOpts)

	assert.True(t, result)
}

func TestHasJWT_PublicMethodNoJWT(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT: "Auth",
	}
	methodOpts := &model.MethodOptions{
		Public: true,
	}

	result := parser.HasJWT(serviceOpts, methodOpts)

	assert.False(t, result, "Public method should not require JWT")
}

func TestHasJWT_NoServiceJWT(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT: "",
	}
	methodOpts := &model.MethodOptions{
		Public: false,
	}

	result := parser.HasJWT(serviceOpts, methodOpts)

	assert.False(t, result)
}

func TestHasJWT_NilOptions(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.HasJWT(nil, nil)

	assert.False(t, result)
}

func TestHasJWT_NilServiceOptions(t *testing.T) {
	parser := NewOptionsParser()
	methodOpts := &model.MethodOptions{
		Public: false,
	}

	result := parser.HasJWT(nil, methodOpts)

	assert.False(t, result)
}

func TestHasJWT_NilMethodOptions_ServiceHasJWT(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT: "Auth",
	}

	result := parser.HasJWT(serviceOpts, nil)

	assert.True(t, result)
}

func TestGetMiddleware_ServiceLevelOnly(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		Middleware: []string{"Authority", "RateLimit"},
	}
	methodOpts := &model.MethodOptions{}

	result := parser.GetMiddleware(serviceOpts, methodOpts)

	assert.Equal(t, []string{"Authority", "RateLimit"}, result)
}

func TestGetMiddleware_MethodOverridesService(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		Middleware: []string{"Authority", "RateLimit"},
	}
	methodOpts := &model.MethodOptions{
		Middleware: []string{"Audit"},
	}

	result := parser.GetMiddleware(serviceOpts, methodOpts)

	assert.Equal(t, []string{"Audit"}, result, "Method middleware should override service")
}

func TestGetMiddleware_NoMiddleware(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{}
	methodOpts := &model.MethodOptions{}

	result := parser.GetMiddleware(serviceOpts, methodOpts)

	assert.Nil(t, result)
}

func TestGetMiddleware_NilOptions(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.GetMiddleware(nil, nil)

	assert.Nil(t, result)
}

func TestGetMiddleware_NilServiceOptions(t *testing.T) {
	parser := NewOptionsParser()
	methodOpts := &model.MethodOptions{
		Middleware: []string{"RateLimit"},
	}

	result := parser.GetMiddleware(nil, methodOpts)

	assert.Equal(t, []string{"RateLimit"}, result)
}

func TestGetMiddleware_NilMethodOptions(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		Middleware: []string{"Authority"},
	}

	result := parser.GetMiddleware(serviceOpts, nil)

	assert.Equal(t, []string{"Authority"}, result)
}

func TestGetMiddleware_EmptyServiceMiddleware(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		Middleware: []string{},
	}
	methodOpts := &model.MethodOptions{}

	result := parser.GetMiddleware(serviceOpts, methodOpts)

	assert.Equal(t, []string{}, result)
}

func TestGetMiddleware_EmptyMethodMiddleware(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		Middleware: []string{"Authority"},
	}
	methodOpts := &model.MethodOptions{
		Middleware: []string{},
	}

	result := parser.GetMiddleware(serviceOpts, methodOpts)

	// Empty slice counts as method having middleware, so it should not return service middleware
	assert.Equal(t, []string{"Authority"}, result)
}

// Test Suite: Complex Scenarios

func TestComplexScenario_PublicLoginWithRateLimit(t *testing.T) {
	parser := NewOptionsParser()

	// Service requires JWT and has Authority middleware
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority"},
		Group:      "user",
		Prefix:     "/api/v1",
	}

	// Login method is public but has RateLimit middleware
	methodOpts := &model.MethodOptions{
		Public:     true,
		Middleware: []string{"RateLimit", "Captcha"},
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT, "Login should not require JWT")
	assert.Equal(t, []string{"RateLimit", "Captcha"}, result.Middleware, "Login should have its own middleware")
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestComplexScenario_ProtectedUpdateWithExtraMiddleware(t *testing.T) {
	parser := NewOptionsParser()

	// Service requires JWT and has basic middleware
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority"},
		Group:      "user",
		Prefix:     "/api/v1",
	}

	// UpdateUser method adds Audit middleware
	methodOpts := &model.MethodOptions{
		Public:     false,
		Middleware: []string{"Authority", "Audit", "Log"},
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Equal(t, "Auth", result.JWT, "UpdateUser should require JWT")
	assert.Equal(t, []string{"Authority", "Audit", "Log"}, result.Middleware, "UpdateUser should have enhanced middleware")
	assert.Equal(t, "user", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestComplexScenario_NoServiceMiddleware_MethodHasMiddleware(t *testing.T) {
	parser := NewOptionsParser()

	// Service has no middleware
	serviceOpts := &model.ServerOptions{
		JWT:    "",
		Group:  "public",
		Prefix: "/api/v1",
	}

	// Method has specific middleware
	methodOpts := &model.MethodOptions{
		Public:     false,
		Middleware: []string{"RateLimit"},
	}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT)
	assert.Equal(t, []string{"RateLimit"}, result.Middleware, "Method should have its middleware")
	assert.Equal(t, "public", result.Group)
	assert.Equal(t, "/api/v1", result.Prefix)
}

func TestComplexScenario_MultipleMiddleware(t *testing.T) {
	parser := NewOptionsParser()

	// Service with multiple middleware
	serviceOpts := &model.ServerOptions{
		JWT:        "Auth",
		Middleware: []string{"Authority", "RateLimit", "Log", "Metrics"},
		Group:      "admin",
		Prefix:     "/api/v1/admin",
	}

	methodOpts := &model.MethodOptions{}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Equal(t, "Auth", result.JWT)
	assert.Equal(t, []string{"Authority", "RateLimit", "Log", "Metrics"}, result.Middleware)
	assert.Equal(t, "admin", result.Group)
	assert.Equal(t, "/api/v1/admin", result.Prefix)
}

// Test Suite: Edge Cases and Boundary Conditions

func TestEdgeCase_VeryLongMiddlewareList(t *testing.T) {
	parser := NewOptionsParser()

	middlewareStr := "MW1,MW2,MW3,MW4,MW5,MW6,MW7,MW8,MW9,MW10"
	result := parser.parseMiddlewareList(middlewareStr)

	assert.Len(t, result, 10)
	assert.Equal(t, "MW1", result[0])
	assert.Equal(t, "MW10", result[9])
}

func TestEdgeCase_SpecialCharactersInMiddleware(t *testing.T) {
	parser := NewOptionsParser()

	result := parser.parseMiddlewareList("Auth-JWT,Rate_Limit,Log.Middleware")

	assert.Equal(t, []string{"Auth-JWT", "Rate_Limit", "Log.Middleware"}, result)
}

func TestEdgeCase_EmptyServerOptions(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{}
	methodOpts := &model.MethodOptions{}

	result := parser.MergeOptions(serviceOpts, methodOpts)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT)
	assert.Empty(t, result.Middleware)
	assert.Empty(t, result.Group)
	assert.Empty(t, result.Prefix)
}

func TestEdgeCase_PublicFalseExplicitly(t *testing.T) {
	parser := NewOptionsParser()
	serviceOpts := &model.ServerOptions{
		JWT: "Auth",
	}
	methodOpts := &model.MethodOptions{
		Public: false, // Explicitly false
	}

	result := parser.HasJWT(serviceOpts, methodOpts)

	assert.True(t, result, "Explicitly false public should still require JWT")
}

func TestEdgeCase_ZeroValueServerOptions(t *testing.T) {
	parser := NewOptionsParser()
	var serviceOpts model.ServerOptions
	var methodOpts model.MethodOptions

	result := parser.MergeOptions(&serviceOpts, &methodOpts)

	require.NotNil(t, result)
	assert.Empty(t, result.JWT)
	assert.Empty(t, result.Middleware)
	assert.Empty(t, result.Group)
	assert.Empty(t, result.Prefix)
}

// Test Suite: NewOptionsParser

func TestNewOptionsParser(t *testing.T) {
	parser := NewOptionsParser()

	require.NotNil(t, parser)
	assert.IsType(t, &OptionsParser{}, parser)
}
