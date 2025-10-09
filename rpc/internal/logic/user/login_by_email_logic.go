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

type LoginByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginByEmail validates user by email and captcha code
func (l *LoginByEmailLogic) LoginByEmail(in *core.LoginByEmailReq) (*core.LoginResp, error) {
	// Query user by email
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.EmailEQ(in.Email)).
		Only(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongEmailOrCaptcha")
	}

	// Check if user is active
	if userInfo.Status != 1 {
		return nil, errorx.NewCodeAbortedError("login.userBanned")
	}

	// Note: Captcha verification happens in API layer
	// This RPC method assumes captcha is already validated

	return &core.LoginResp{
		Code: 0,
		Msg:  "login.loginSuccessTitle",
		Data: &core.LoginInfo{
			UserId: userInfo.ID.String(),
		},
	}, nil
}
