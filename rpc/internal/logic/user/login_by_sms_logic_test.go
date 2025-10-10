package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"

	_ "github.com/mattn/go-sqlite3"
)

func TestLoginBySmsLogic_LoginBySms_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with mobile number
	testPassword := "password123"
	hashedPassword := encrypt.BcryptEncrypt(testPassword)
	testMobile := "+1234567890"

	userInfo, err := svcCtx.DB.User.Create().
		SetUsername("smsuser").
		SetPassword(hashedPassword).
		SetNickname("SMS User").
		SetStatus(1).
		SetMobile(testMobile).
		SetEmail("sms@example.com").
		Save(context.Background())
	require.NoError(t, err)

	// Test SMS login
	logic := NewLoginBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.LoginBySms(&core.LoginBySmsReq{
		PhoneNumber: testMobile,
		Captcha:     "123456", // Mock SMS code
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.NotNil(t, resp.Data)
	assert.Equal(t, userInfo.ID.String(), resp.Data.UserId)
}

func TestLoginBySmsLogic_LoginBySms_PhoneNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test login with non-existent phone number
	logic := NewLoginBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.LoginBySms(&core.LoginBySmsReq{
		PhoneNumber: "+9999999999",
		Captcha:     "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.wrongPhoneNumber")
}

func TestLoginBySmsLogic_LoginBySms_InactiveUser(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create inactive user with mobile
	testMobile := "+1234567890"
	hashedPassword := encrypt.BcryptEncrypt("password123")

	_, err := svcCtx.DB.User.Create().
		SetUsername("inactivesmsuser").
		SetPassword(hashedPassword).
		SetNickname("Inactive SMS User").
		SetStatus(0). // Inactive
		SetMobile(testMobile).
		SetEmail("inactive@example.com").
		Save(context.Background())
	require.NoError(t, err)

	// Test login with inactive user
	logic := NewLoginBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.LoginBySms(&core.LoginBySmsReq{
		PhoneNumber: testMobile,
		Captcha:     "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.userBanned")
}

func TestLoginBySmsLogic_LoginBySms_EmptyPhoneNumber(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test login with empty phone number
	logic := NewLoginBySmsLogic(context.Background(), svcCtx)
	resp, err := logic.LoginBySms(&core.LoginBySmsReq{
		PhoneNumber: "",
		Captcha:     "123456",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}
