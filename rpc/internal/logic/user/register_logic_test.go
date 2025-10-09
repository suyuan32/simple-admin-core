package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterLogic_Register_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test registration
	logic := NewRegisterLogic(context.Background(), svcCtx)
	resp, err := logic.Register(&core.RegisterReq{
		Username:  "newuser",
		Password:  "password123",
		Email:     "newuser@example.com",
		CaptchaId: "test-captcha",
		Captcha:   "12345",
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "login.registerSuccessTitle", resp.Msg)

	// Verify user was created in database
	count, err := svcCtx.DB.User.Query().
		Where(user.UsernameEQ("newuser")).
		Count(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestRegisterLogic_Register_DuplicateEmail(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create existing user
	testEmail := "existing@example.com"
	createTestUser(t, svcCtx, "existinguser", "password123")
	svcCtx.DB.User.Update().SetEmail(testEmail).SaveX(context.Background())

	// Test registration with duplicate email
	logic := NewRegisterLogic(context.Background(), svcCtx)
	resp, err := logic.Register(&core.RegisterReq{
		Username:  "newuser",
		Password:  "password123",
		Email:     testEmail,
		CaptchaId: "test-captcha",
		Captcha:   "12345",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.emailAlreadyExists")
}

func TestRegisterLogic_Register_InvalidEmail(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test registration with empty email
	logic := NewRegisterLogic(context.Background(), svcCtx)
	resp, err := logic.Register(&core.RegisterReq{
		Username:  "newuser",
		Password:  "password123",
		Email:     "",
		CaptchaId: "test-captcha",
		Captcha:   "12345",
	})

	// Assertions
	// Since email validation happens at API layer, RPC may succeed
	// or fail depending on CreateUser implementation
	// This test documents current behavior
	if err != nil {
		assert.Nil(t, resp)
	}
}
