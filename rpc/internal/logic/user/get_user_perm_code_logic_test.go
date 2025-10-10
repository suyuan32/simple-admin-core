package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetUserPermCodeLogic_GetUserPermCode_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	userInfo := createTestUser(t, svcCtx, "permuser", "password123")

	// Create test role
	role, err := svcCtx.DB.Role.Create().
		SetName("admin").
		SetCode("admin_role").
		SetDefaultRouter("dashboard").
		SetStatus(1).
		SetSort(1).
		Save(context.Background())
	require.NoError(t, err)

	// Create test menus with permissions
	menu1, err := svcCtx.DB.Menu.Create().
		SetMenuType(1).
		SetTitle("User Management").
		SetName("user_management").
		SetPath("/users").
		SetComponent("UserList").
		SetPermission("user:read").
		SetSort(1).
		Save(context.Background())
	require.NoError(t, err)

	menu2, err := svcCtx.DB.Menu.Create().
		SetMenuType(1).
		SetTitle("User Create").
		SetName("user_create").
		SetPath("/users/create").
		SetComponent("UserCreate").
		SetPermission("user:create").
		SetSort(2).
		Save(context.Background())
	require.NoError(t, err)

	// Associate role with menus
	_, err = svcCtx.DB.Role.UpdateOne(role).
		AddMenus(menu1, menu2).
		Save(context.Background())
	require.NoError(t, err)

	// Associate user with role
	_, err = svcCtx.DB.User.UpdateOne(userInfo).
		AddRoles(role).
		Save(context.Background())
	require.NoError(t, err)

	// Test getting permission codes
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewGetUserPermCodeLogic(ctx, svcCtx)
	resp, err := logic.GetUserPermCode(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Len(t, resp.Data, 2)
	assert.Contains(t, resp.Data, "user:read")
	assert.Contains(t, resp.Data, "user:create")
}

func TestGetUserPermCodeLogic_GetUserPermCode_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test without userId in context
	logic := NewGetUserPermCodeLogic(context.Background(), svcCtx)
	resp, err := logic.GetUserPermCode(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestGetUserPermCodeLogic_GetUserPermCode_InvalidUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with invalid UUID format
	ctx := context.WithValue(context.Background(), "userId", "invalid-uuid")
	logic := NewGetUserPermCodeLogic(ctx, svcCtx)
	resp, err := logic.GetUserPermCode(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.invalidUserId")
}

func TestGetUserPermCodeLogic_GetUserPermCode_NoRoles(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user without roles
	userInfo := createTestUser(t, svcCtx, "noroleuser", "password123")

	// Test getting permission codes
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewGetUserPermCodeLogic(ctx, svcCtx)
	resp, err := logic.GetUserPermCode(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Empty(t, resp.Data) // No permissions
}

func TestGetUserPermCodeLogic_GetUserPermCode_MultipleRoles(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	userInfo := createTestUser(t, svcCtx, "multiuser", "password123")

	// Create first role with permissions
	role1, err := svcCtx.DB.Role.Create().
		SetName("editor").
		SetCode("editor_role").
		SetDefaultRouter("dashboard").
		SetStatus(1).
		SetSort(1).
		Save(context.Background())
	require.NoError(t, err)

	menu1, err := svcCtx.DB.Menu.Create().
		SetMenuType(1).
		SetTitle("Content Edit").
		SetName("content_edit").
		SetPath("/content/edit").
		SetComponent("ContentEdit").
		SetPermission("content:edit").
		SetSort(1).
		Save(context.Background())
	require.NoError(t, err)

	// Create second role with different permissions
	role2, err := svcCtx.DB.Role.Create().
		SetName("viewer").
		SetCode("viewer_role").
		SetDefaultRouter("dashboard").
		SetStatus(1).
		SetSort(2).
		Save(context.Background())
	require.NoError(t, err)

	menu2, err := svcCtx.DB.Menu.Create().
		SetMenuType(1).
		SetTitle("Content View").
		SetName("content_view").
		SetPath("/content/view").
		SetComponent("ContentView").
		SetPermission("content:read").
		SetSort(2).
		Save(context.Background())
	require.NoError(t, err)

	// Associate roles with menus
	_, err = svcCtx.DB.Role.UpdateOne(role1).AddMenus(menu1).Save(context.Background())
	require.NoError(t, err)
	_, err = svcCtx.DB.Role.UpdateOne(role2).AddMenus(menu2).Save(context.Background())
	require.NoError(t, err)

	// Associate user with both roles
	_, err = svcCtx.DB.User.UpdateOne(userInfo).
		AddRoles(role1, role2).
		Save(context.Background())
	require.NoError(t, err)

	// Test getting permission codes
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewGetUserPermCodeLogic(ctx, svcCtx)
	resp, err := logic.GetUserPermCode(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Len(t, resp.Data, 2) // Combined permissions from both roles
	assert.Contains(t, resp.Data, "content:edit")
	assert.Contains(t, resp.Data, "content:read")
}
