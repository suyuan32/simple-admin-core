package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserProfileLogic_UpdateUserProfile_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	userInfo := createTestUser(t, svcCtx, "testuser", "password123")

	// Test update profile
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewUpdateUserProfileLogic(ctx, svcCtx)

	newNickname := "Updated Nickname"
	newMobile := "+886912345678"
	resp, err := logic.UpdateUserProfile(&core.UpdateProfileReq{
		Nickname: &newNickname,
		Mobile:   &newMobile,
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "common.updateSuccess", resp.Msg)

	// Verify update persisted
	updatedUser, err := svcCtx.DB.User.Get(context.Background(), userInfo.ID)
	require.NoError(t, err)
	assert.Equal(t, newNickname, updatedUser.Nickname)
	assert.Equal(t, newMobile, updatedUser.Mobile)
}

func TestUpdateUserProfileLogic_UpdateUserProfile_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test without userId in context
	logic := NewUpdateUserProfileLogic(context.Background(), svcCtx)
	nickname := "Test"
	resp, err := logic.UpdateUserProfile(&core.UpdateProfileReq{
		Nickname: &nickname,
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestUpdateUserProfileLogic_UpdateUserProfile_EmptyUpdate(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	userInfo := createTestUser(t, svcCtx, "testuser", "password123")

	// Test update with no fields
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewUpdateUserProfileLogic(ctx, svcCtx)

	resp, err := logic.UpdateUserProfile(&core.UpdateProfileReq{})

	// Should succeed even with no updates
	require.NoError(t, err)
	assert.NotNil(t, resp)
}
