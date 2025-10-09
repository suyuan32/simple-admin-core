package user

import (
	"context"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"
	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login validates user credentials and returns user info
// Note: JWT token generation happens in API layer
func (l *LoginLogic) Login(in *core.LoginReq) (*core.LoginResp, error) {
	// Query user by username
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.UsernameEQ(in.Username)).
		Only(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
	}

	// Check if user is active
	if userInfo.Status != 1 {
		return nil, errorx.NewCodeAbortedError("login.userBanned")
	}

	// Verify password
	if !encrypt.BcryptCheck(in.Password, userInfo.Password) {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
	}

	// Return user info (API layer will generate JWT token)
	return &core.LoginResp{
		Code: 0,
		Msg:  "login.loginSuccessTitle",
		Data: &core.LoginInfo{
			UserId: userInfo.ID.String(),
			// Token and Expire will be set by API layer
		},
	}, nil
}
