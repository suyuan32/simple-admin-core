package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"

	_ "github.com/mattn/go-sqlite3"
)

func TestResetPasswordByEmailLogic_ResetPasswordByEmail_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with email
	testEmail := "reset@example.com"
	oldPassword := "oldpassword123"
	userInfo := createTestUser(t, svcCtx, "resetuser", oldPassword)
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetEmail(testEmail).
		Save(context.Background())
	require.NoError(t, err)

	// Test password reset
	newPassword := "newpassword456"
	logic := NewResetPasswordByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordByEmail(&core.ResetPasswordByEmailReq{
		Email:       testEmail,
		Captcha:     "123456", // Mock email verification code
		NewPassword: newPassword,
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Equal(t, "common.updateSuccess", resp.Msg)

	// Verify password was updated
	updatedUser, err := svcCtx.DB.User.Query().
		Where(user.Email(testEmail)).
		Only(context.Background())
	require.NoError(t, err)

	// Verify new password works (old password shouldn't)
	assert.True(t, encrypt.BcryptCheck(newPassword, updatedUser.Password))
	assert.False(t, encrypt.BcryptCheck(oldPassword, updatedUser.Password))
}

func TestResetPasswordByEmailLogic_ResetPasswordByEmail_EmailNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test password reset with non-existent email
	logic := NewResetPasswordByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordByEmail(&core.ResetPasswordByEmailReq{
		Email:       "nonexistent@example.com",
		Captcha:     "123456",
		NewPassword: "newpassword456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.wrongEmail")
}

func TestResetPasswordByEmailLogic_ResetPasswordByEmail_WeakPassword(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with email
	testEmail := "reset@example.com"
	userInfo := createTestUser(t, svcCtx, "resetuser", "oldpassword123")
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetEmail(testEmail).
		Save(context.Background())
	require.NoError(t, err)

	// Test password reset with weak password
	logic := NewResetPasswordByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordByEmail(&core.ResetPasswordByEmailReq{
		Email:       testEmail,
		Captcha:     "123456",
		NewPassword: "123", // Too weak
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user.passwordTooWeak")
}

func TestResetPasswordByEmailLogic_ResetPasswordByEmail_InvalidCaptcha(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with email
	testEmail := "reset@example.com"
	userInfo := createTestUser(t, svcCtx, "resetuser", "oldpassword123")
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetEmail(testEmail).
		Save(context.Background())
	require.NoError(t, err)

	// Test password reset with invalid captcha
	logic := NewResetPasswordByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordByEmail(&core.ResetPasswordByEmailReq{
		Email:       testEmail,
		Captcha:     "", // Empty captcha
		NewPassword: "newpassword456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}
