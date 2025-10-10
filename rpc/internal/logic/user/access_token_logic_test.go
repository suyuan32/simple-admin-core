package user

import (
	"context"
	"testing"
	"time"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"

	_ "github.com/mattn/go-sqlite3"
)

func TestAccessTokenLogic_AccessToken_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create active test user
	userInfo := createTestUser(t, svcCtx, "tokenuser", "password123")

	// Test access token generation
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewAccessTokenLogic(ctx, svcCtx)
	resp, err := logic.AccessToken(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
	assert.Equal(t, "common.success", resp.Msg)
	assert.NotNil(t, resp.Data)

	// Verify expiry is approximately 2 hours from now
	expectedExpiry := time.Now().Add(2 * time.Hour).Unix()
	actualExpiry := resp.Data.ExpiredAt
	timeDiff := actualExpiry - expectedExpiry
	assert.True(t, timeDiff >= -5 && timeDiff <= 5, "Expiry time should be ~2 hours from now")
}

func TestAccessTokenLogic_AccessToken_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test without userId in context
	logic := NewAccessTokenLogic(context.Background(), svcCtx)
	resp, err := logic.AccessToken(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestAccessTokenLogic_AccessToken_InvalidUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with invalid UUID format
	ctx := context.WithValue(context.Background(), "userId", "invalid-uuid")
	logic := NewAccessTokenLogic(ctx, svcCtx)
	resp, err := logic.AccessToken(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.invalidUserId")
}

func TestAccessTokenLogic_AccessToken_UserNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with non-existent user ID
	nonExistentUUID := "00000000-0000-0000-0000-000000000000"
	ctx := context.WithValue(context.Background(), "userId", nonExistentUUID)
	logic := NewAccessTokenLogic(ctx, svcCtx)
	resp, err := logic.AccessToken(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestAccessTokenLogic_AccessToken_InactiveUser(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create inactive user
	hashedPassword := encrypt.BcryptEncrypt("password123")
	userInfo, err := svcCtx.DB.User.Create().
		SetUsername("inactiveuser").
		SetPassword(hashedPassword).
		SetNickname("Inactive User").
		SetStatus(0). // Inactive
		SetEmail("inactive@example.com").
		Save(context.Background())
	require.NoError(t, err)

	// Test access token generation with inactive user
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewAccessTokenLogic(ctx, svcCtx)
	resp, err := logic.AccessToken(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "login.userBanned")
}

func TestAccessTokenLogic_AccessToken_ExpiryTime(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	userInfo := createTestUser(t, svcCtx, "expiryuser", "password123")

	// Test access token generation
	beforeTime := time.Now()
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewAccessTokenLogic(ctx, svcCtx)
	resp, err := logic.AccessToken(&core.Empty{})
	afterTime := time.Now()

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)

	// Verify expiry is between beforeTime+2h and afterTime+2h
	minExpiry := beforeTime.Add(2 * time.Hour).Unix()
	maxExpiry := afterTime.Add(2 * time.Hour).Unix()
	actualExpiry := resp.Data.ExpiredAt

	assert.True(t, actualExpiry >= minExpiry && actualExpiry <= maxExpiry,
		"Expiry time should be exactly 2 hours from token generation time")
}
