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

func TestRegisterBySmsLogic_RegisterBySms_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test SMS registration
	logic := NewRegisterBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterBySms(&core.RegisterBySmsReq{
		Username:    "smsreguser",
		PhoneNumber: "+1234567890",
		Password:    "SecurePass123!",
		Captcha:     "123456", // Mock SMS verification code
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.NotNil(t, resp.Data)
	assert.NotEmpty(t, resp.Data.UserId)

	// Verify user was created in database
	createdUser, err := svcCtx.DB.User.Query().
		Where(user.Mobile("+1234567890")).
		Only(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "smsreguser", createdUser.Username)
	assert.Equal(t, "+1234567890", createdUser.Mobile)
	assert.Equal(t, uint8(1), createdUser.Status) // Active by default
}

func TestRegisterBySmsLogic_RegisterBySms_DuplicatePhone(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create existing user with phone number
	testMobile := "+1234567890"
	userInfo := createTestUser(t, svcCtx, "existinguser", "password123")
	_, err := svcCtx.DB.User.UpdateOne(userInfo).
		SetMobile(testMobile).
		Save(context.Background())
	require.NoError(t, err)

	// Test registration with duplicate phone
	logic := NewRegisterBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterBySms(&core.RegisterBySmsReq{
		Username:    "newuser",
		PhoneNumber: testMobile,
		Password:    "SecurePass123!",
		Captcha:     "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "register.phoneAlreadyExists")
}

func TestRegisterBySmsLogic_RegisterBySms_InvalidPhoneFormat(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test registration with invalid phone format
	logic := NewRegisterBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterBySms(&core.RegisterBySmsReq{
		Username:    "newuser",
		PhoneNumber: "invalid-phone",
		Password:    "SecurePass123!",
		Captcha:     "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestRegisterBySmsLogic_RegisterBySms_EmptyPhone(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test registration with empty phone number
	logic := NewRegisterBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.RegisterBySms(&core.RegisterBySmsReq{
		Username:    "newuser",
		PhoneNumber: "",
		Password:    "SecurePass123!",
		Captcha:     "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}
