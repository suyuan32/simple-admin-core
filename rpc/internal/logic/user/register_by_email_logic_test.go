package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestRegisterByEmailLogic_RegisterByEmail_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test email registration
	logic := NewRegisterByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterByEmail(&core.RegisterByEmailReq{
		Username: "emailuser",
		Email:    "newuser@example.com",
		Password: "SecurePass123!",
		Captcha:  "123456", // Mock email verification code
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.NotNil(t, resp.Data)
	assert.NotEmpty(t, resp.Data.UserId)

	// Verify user was created in database
	createdUser, err := svcCtx.DB.User.Query().
		Where(user.Email("newuser@example.com")).
		Only(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "emailuser", createdUser.Username)
	assert.Equal(t, "newuser@example.com", createdUser.Email)
	assert.Equal(t, uint8(1), createdUser.Status) // Active by default
}

func TestRegisterByEmailLogic_RegisterByEmail_DuplicateEmail(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create existing user with email
	existingEmail := "existing@example.com"
	createTestUser(t, svcCtx, "existinguser", "password123")
	_, err := svcCtx.DB.User.Query().
		Where(user.Username("existinguser")).
		Only(context.Background())
	require.NoError(t, err)

	// Update user with email
	_, err = svcCtx.DB.User.Update().
		Where(user.Username("existinguser")).
		SetEmail(existingEmail).
		Save(context.Background())
	require.NoError(t, err)

	// Test registration with duplicate email
	logic := NewRegisterByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterByEmail(&core.RegisterByEmailReq{
		Username: "newuser",
		Email:    existingEmail,
		Password: "SecurePass123!",
		Captcha:  "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "register.emailAlreadyExists")
}

func TestRegisterByEmailLogic_RegisterByEmail_InvalidEmail(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test registration with invalid email format
	logic := NewRegisterByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterByEmail(&core.RegisterByEmailReq{
		Username: "newuser",
		Email:    "invalid-email-format",
		Password: "SecurePass123!",
		Captcha:  "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestRegisterByEmailLogic_RegisterByEmail_WeakPassword(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test registration with weak password
	logic := NewRegisterByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterByEmail(&core.RegisterByEmailReq{
		Username: "newuser",
		Email:    "newuser@example.com",
		Password: "123", // Too weak
		Captcha:  "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "register.passwordTooWeak")
}
