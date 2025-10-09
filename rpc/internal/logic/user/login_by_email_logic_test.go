package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginByEmailLogic_LoginByEmail_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with email
	testEmail := "test@example.com"
	userInfo := createTestUser(t, svcCtx, "testuser", "password123")
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetEmail(testEmail).
		Save(context.Background())
	require.NoError(t, err)

	// Test email login
	logic := NewLoginByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.LoginByEmail(&core.LoginByEmailReq{
		Email:   testEmail,
		Captcha: "12345",
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Equal(t, "login.loginSuccessTitle", resp.Msg)
	assert.NotNil(t, resp.Data)
	assert.NotEmpty(t, resp.Data.UserId)
}

func TestLoginByEmailLogic_LoginByEmail_EmailNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test login with non-existent email
	logic := NewLoginByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.LoginByEmail(&core.LoginByEmailReq{
		Email:   "nonexistent@example.com",
		Captcha: "12345",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.wrongUsernameOrPassword")
}

func TestLoginByEmailLogic_LoginByEmail_InactiveUser(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create inactive user
	testEmail := "inactive@example.com"
	userInfo := createTestUser(t, svcCtx, "inactiveuser", "password123")
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetEmail(testEmail).
		SetStatus(0). // Inactive
		Save(context.Background())
	require.NoError(t, err)

	// Test login with inactive user
	logic := NewLoginByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.LoginByEmail(&core.LoginByEmailReq{
		Email:   testEmail,
		Captcha: "12345",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.userBanned")
}
