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

func TestResetPasswordBySmsLogic_ResetPasswordBySms_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with mobile
	testMobile := "+1234567890"
	oldPassword := "oldpassword123"
	userInfo := createTestUser(t, svcCtx, "resetuser", oldPassword)
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetMobile(testMobile).
		Save(context.Background())
	require.NoError(t, err)

	// Test password reset via SMS
	newPassword := "newpassword456"
	logic := NewResetPasswordBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordBySms(&core.ResetPasswordBySmsReq{
		PhoneNumber: testMobile,
		Captcha:     "123456", // Mock SMS verification code
		NewPassword: newPassword,
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Equal(t, "common.updateSuccess", resp.Msg)

	// Verify password was updated
	updatedUser, err := svcCtx.DB.User.Query().
		Where(user.Mobile(testMobile)).
		Only(context.Background())
	require.NoError(t, err)

	// Verify new password works (old password shouldn't)
	assert.True(t, encrypt.BcryptCheck(newPassword, updatedUser.Password))
	assert.False(t, encrypt.BcryptCheck(oldPassword, updatedUser.Password))
}

func TestResetPasswordBySmsLogic_ResetPasswordBySms_PhoneNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test password reset with non-existent phone number
	logic := NewResetPasswordBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordBySms(&core.ResetPasswordBySmsReq{
		PhoneNumber: "+9999999999",
		Captcha:     "123456",
		NewPassword: "newpassword456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.wrongPhoneNumber")
}

func TestResetPasswordBySmsLogic_ResetPasswordBySms_WeakPassword(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with mobile
	testMobile := "+1234567890"
	userInfo := createTestUser(t, svcCtx, "resetuser", "oldpassword123")
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetMobile(testMobile).
		Save(context.Background())
	require.NoError(t, err)

	// Test password reset with weak password
	logic := NewResetPasswordBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordBySms(&core.ResetPasswordBySmsReq{
		PhoneNumber: testMobile,
		Captcha:     "123456",
		NewPassword: "123", // Too weak
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user.passwordTooWeak")
}

func TestResetPasswordBySmsLogic_ResetPasswordBySms_EmptyPhone(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test password reset with empty phone number
	logic := NewResetPasswordBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.ResetPasswordBySms(&core.ResetPasswordBySmsReq{
		PhoneNumber: "",
		Captcha:     "123456",
		NewPassword: "newpassword456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}
