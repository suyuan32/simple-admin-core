package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserInfoLogic_GetUserInfo_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	testUsername := "testuser"
	userInfo := createTestUser(t, svcCtx, testUsername, "password123")

	// Test get user info with userId in context
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewGetUserInfoLogic(ctx, svcCtx)
	resp, err := logic.GetUserInfo(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.NotNil(t, resp.Data)
	assert.Equal(t, testUsername, *resp.Data.Username)
}

func TestGetUserInfoLogic_GetUserInfo_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test without userId in context
	logic := NewGetUserInfoLogic(context.Background(), svcCtx)
	resp, err := logic.GetUserInfo(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestGetUserInfoLogic_GetUserInfo_InvalidUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with invalid userId
	ctx := context.WithValue(context.Background(), "userId", "invalid-uuid")
	logic := NewGetUserInfoLogic(ctx, svcCtx)
	resp, err := logic.GetUserInfo(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestGetUserInfoLogic_GetUserInfo_UserNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with non-existent userId
	ctx := context.WithValue(context.Background(), "userId", "00000000-0000-0000-0000-000000000000")
	logic := NewGetUserInfoLogic(ctx, svcCtx)
	resp, err := logic.GetUserInfo(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}
