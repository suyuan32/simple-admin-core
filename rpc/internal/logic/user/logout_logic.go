package user

import (
	"context"

	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Logout handles user logout
// Note: Since JWT is stateless, actual token invalidation happens in API layer
// This method can be used for logging, cleanup, or adding token to blacklist
func (l *LogoutLogic) Logout(in *core.Empty) (*core.BaseResp, error) {
	// Get user ID from context (if available)
	userId, ok := l.ctx.Value("userId").(string)
	if ok && userId != "" {
		// Log the logout event
		l.Logger.Infof("User %s logged out", userId)

		// Optional: Add current token to Redis blacklist
		// This would require the token to be passed in the context
		// token := l.ctx.Value("token").(string)
		// l.svcCtx.Redis.Set(ctx, "blacklist:"+token, "1", tokenExpiry)
	}

	return &core.BaseResp{
		Msg: "login.logoutSuccessTitle",
	}, nil
}
