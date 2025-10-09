package user

import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogoutLogic_Logout_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	userInfo := createTestUser(t, svcCtx, "testuser", "password123")

	// Create active token
	tokenInfo, err := svcCtx.DB.Token.Create().
		SetUUID(userInfo.ID).
		SetToken("test-token-123").
		SetSource("core_user").
		SetStatus(1).
		SetUsername(userInfo.Username).
		SetExpiredAt(9999999999999).
		Save(context.Background())
	require.NoError(t, err)

	// Test logout
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewLogoutLogic(ctx, svcCtx)

	resp, err := logic.Logout(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "login.logoutSuccessTitle", resp.Msg)

	// Verify token was invalidated (status set to 0)
	updatedToken, err := svcCtx.DB.Token.Get(context.Background(), tokenInfo.ID)
	require.NoError(t, err)
	assert.Equal(t, uint8(0), updatedToken.Status)
}

func TestLogoutLogic_Logout_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test without userId in context
	logic := NewLogoutLogic(context.Background(), svcCtx)
	resp, err := logic.Logout(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestLogoutLogic_Logout_NoActiveTokens(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user but no tokens
	userInfo := createTestUser(t, svcCtx, "testuser", "password123")

	// Test logout without active tokens
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewLogoutLogic(ctx, svcCtx)

	resp, err := logic.Logout(&core.Empty{})

	// Should succeed even without tokens
	require.NoError(t, err)
	assert.NotNil(t, resp)
}
