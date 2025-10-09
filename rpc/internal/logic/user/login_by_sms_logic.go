package user

import (
	"context"

	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginBySmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginBySmsLogic {
	return &LoginBySmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginBySms validates user by phone number and SMS captcha
func (l *LoginBySmsLogic) LoginBySms(in *core.LoginBySmsReq) (*core.LoginResp, error) {
	// Query user by mobile phone
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.MobileEQ(in.PhoneNumber)).
		Only(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongPhoneOrCaptcha")
	}

	// Check if user is active
	if userInfo.Status != 1 {
		return nil, errorx.NewCodeAbortedError("login.userBanned")
	}

	// Note: SMS captcha verification happens in API layer

	return &core.LoginResp{
		Code: 0,
		Msg:  "login.loginSuccessTitle",
		Data: &core.LoginInfo{
			UserId: userInfo.ID.String(),
		},
	}, nil
}
