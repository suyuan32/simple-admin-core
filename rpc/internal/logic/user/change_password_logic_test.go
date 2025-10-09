package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangePasswordLogic_ChangePassword_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	testUsername := "testuser"
	oldPassword := "oldpassword123"
	userInfo := createTestUser(t, svcCtx, testUsername, oldPassword)

	// Test password change
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewChangePasswordLogic(ctx, svcCtx)

	newPassword := "newpassword456"
	resp, err := logic.ChangePassword(&core.ChangePasswordReq{
		OldPassword: oldPassword,
		NewPassword: newPassword,
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "user.passwordChangeSuccess", resp.Msg)

	// Verify password was actually changed
	updatedUser, err := svcCtx.DB.User.Get(context.Background(), userInfo.ID)
	require.NoError(t, err)
	assert.True(t, encrypt.BcryptCheck(newPassword, updatedUser.Password))
}

func TestChangePasswordLogic_ChangePassword_WrongOldPassword(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	testUsername := "testuser"
	oldPassword := "oldpassword123"
	userInfo := createTestUser(t, svcCtx, testUsername, oldPassword)

	// Test password change with wrong old password
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewChangePasswordLogic(ctx, svcCtx)

	resp, err := logic.ChangePassword(&core.ChangePasswordReq{
		OldPassword: "wrongoldpassword",
		NewPassword: "newpassword456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user.wrongPassword")
}

func TestChangePasswordLogic_ChangePassword_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test password change without userId in context
	logic := NewChangePasswordLogic(context.Background(), svcCtx)

	resp, err := logic.ChangePassword(&core.ChangePasswordReq{
		OldPassword: "oldpassword",
		NewPassword: "newpassword",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestChangePasswordLogic_ChangePassword_SamePassword(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	testUsername := "testuser"
	password := "password123"
	userInfo := createTestUser(t, svcCtx, testUsername, password)

	// Test changing to same password
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewChangePasswordLogic(ctx, svcCtx)

	resp, err := logic.ChangePassword(&core.ChangePasswordReq{
		OldPassword: password,
		NewPassword: password,
	})

	// Assertions
	// Should succeed (changing to same password is allowed)
	require.NoError(t, err)
	assert.NotNil(t, resp)
}
