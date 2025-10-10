package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetUserProfileLogic_GetUserProfile_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with full profile
	userInfo := createTestUser(t, svcCtx, "profileuser", "password123")
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetNickname("Profile User").
		SetAvatar("https://example.com/avatar.jpg").
		SetMobile("+1234567890").
		SetEmail("profile@example.com").
		SetLocale("en_US").
		Save(context.Background())
	require.NoError(t, err)

	// Test getting user profile
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewGetUserProfileLogic(ctx, svcCtx)
	resp, err := logic.GetUserProfile(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.NotNil(t, resp.Data)
	assert.Equal(t, "Profile User", *resp.Data.Nickname)
	assert.Equal(t, "https://example.com/avatar.jpg", *resp.Data.Avatar)
	assert.Equal(t, "+1234567890", *resp.Data.Mobile)
	assert.Equal(t, "profile@example.com", *resp.Data.Email)
	assert.Equal(t, "en_US", *resp.Data.Locale)
}

func TestGetUserProfileLogic_GetUserProfile_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test without userId in context
	logic := NewGetUserProfileLogic(context.Background(), svcCtx)
	resp, err := logic.GetUserProfile(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestGetUserProfileLogic_GetUserProfile_InvalidUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with invalid UUID format
	ctx := context.WithValue(context.Background(), "userId", "invalid-uuid")
	logic := NewGetUserProfileLogic(ctx, svcCtx)
	resp, err := logic.GetUserProfile(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.invalidUserId")
}

func TestGetUserProfileLogic_GetUserProfile_UserNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with non-existent user ID
	nonExistentUUID := "00000000-0000-0000-0000-000000000000"
	ctx := context.WithValue(context.Background(), "userId", nonExistentUUID)
	logic := NewGetUserProfileLogic(ctx, svcCtx)
	resp, err := logic.GetUserProfile(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestGetUserProfileLogic_GetUserProfile_PartialData(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with minimal profile data
	userInfo := createTestUser(t, svcCtx, "minimaluser", "password123")

	// Test getting user profile with minimal data
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewGetUserProfileLogic(ctx, svcCtx)
	resp, err := logic.GetUserProfile(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.NotNil(t, resp.Data)
	// Some fields may be empty or default values
	assert.NotNil(t, resp.Data.Nickname)
	assert.NotNil(t, resp.Data.Email)
}
