package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/ent/enttest"
	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"

	_ "github.com/mattn/go-sqlite3"
)

// setupTestDB creates an in-memory test database
func setupTestDB(t *testing.T) *svc.ServiceContext {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	svcCtx := &svc.ServiceContext{
		DB: client,
	}

	return svcCtx
}

// createTestUser creates a test user in the database
func createTestUser(t *testing.T, client *svc.ServiceContext, username string, password string) *ent.User {
	hashedPassword := encrypt.BcryptEncrypt(password)

	userInfo, err := client.DB.User.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		SetNickname("Test User").
		SetStatus(1).
		SetEmail("test@example.com").
		Save(context.Background())

	require.NoError(t, err)
	return userInfo
}

func TestLoginLogic_Login_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	testUsername := "testuser"
	testPassword := "password123"
	createTestUser(t, svcCtx, testUsername, testPassword)

	// Test login
	logic := NewLoginLogic(context.Background(), svcCtx)
	resp, err := logic.Login(&core.LoginReq{
		Username:  testUsername,
		Password:  testPassword,
		CaptchaId: "test-captcha-id",
		Captcha:   "12345",
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Equal(t, "login.loginSuccessTitle", resp.Msg)
	assert.NotNil(t, resp.Data)
	assert.NotEmpty(t, resp.Data.UserId)
}

func TestLoginLogic_Login_InvalidUsername(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test login with non-existent user
	logic := NewLoginLogic(context.Background(), svcCtx)
	resp, err := logic.Login(&core.LoginReq{
		Username:  "nonexistent",
		Password:  "password123",
		CaptchaId: "test-captcha-id",
		Captcha:   "12345",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.wrongUsernameOrPassword")
}

func TestLoginLogic_Login_InvalidPassword(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	testUsername := "testuser"
	testPassword := "password123"
	createTestUser(t, svcCtx, testUsername, testPassword)

	// Test login with wrong password
	logic := NewLoginLogic(context.Background(), svcCtx)
	resp, err := logic.Login(&core.LoginReq{
		Username:  testUsername,
		Password:  "wrongpassword",
		CaptchaId: "test-captcha-id",
		Captcha:   "12345",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.wrongUsernameOrPassword")
}

func TestLoginLogic_Login_InactiveUser(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create inactive user
	testUsername := "inactiveuser"
	testPassword := "password123"
	hashedPassword := encrypt.BcryptEncrypt(testPassword)

	_, err := svcCtx.DB.User.Create().
		SetUsername(testUsername).
		SetPassword(hashedPassword).
		SetNickname("Inactive User").
		SetStatus(0). // Inactive status
		SetEmail("inactive@example.com").
		Save(context.Background())
	require.NoError(t, err)

	// Test login with inactive user
	logic := NewLoginLogic(context.Background(), svcCtx)
	resp, err := logic.Login(&core.LoginReq{
		Username:  testUsername,
		Password:  testPassword,
		CaptchaId: "test-captcha-id",
		Captcha:   "12345",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.userBanned")
}
