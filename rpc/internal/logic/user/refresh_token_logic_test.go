package user

import (
	"context"
	"testing"
	"time"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefreshTokenLogic_RefreshToken_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user
	userInfo := createTestUser(t, svcCtx, "testuser", "password123")

	// Create active token that will expire soon
	expiredAt := time.Now().Add(5 * time.Minute).UnixMilli()
	_, err := svcCtx.DB.Token.Create().
		SetUUID(userInfo.ID).
		SetToken("old-token-123").
		SetSource("core_user").
		SetStatus(1).
		SetUsername(userInfo.Username).
		SetExpiredAt(expiredAt).
		Save(context.Background())
	require.NoError(t, err)

	// Test refresh token
	ctx := context.WithValue(context.Background(), "userId", userInfo.ID.String())
	logic := NewRefreshTokenLogic(ctx, svcCtx)

	resp, err := logic.RefreshToken(&core.Empty{})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotEmpty(t, resp.Data.Token)
	assert.Greater(t, resp.Data.ExpiredAt, uint64(time.Now().UnixMilli()))
}

func TestRefreshTokenLogic_RefreshToken_MissingUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test without userId in context
	logic := NewRefreshTokenLogic(context.Background(), svcCtx)
	resp, err := logic.RefreshToken(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "common.unauthorized")
}

func TestRefreshTokenLogic_RefreshToken_InvalidUserId(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test with invalid userId
	ctx := context.WithValue(context.Background(), "userId", "invalid-uuid")
	logic := NewRefreshTokenLogic(ctx, svcCtx)

	resp, err := logic.RefreshToken(&core.Empty{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}
