package user

import (
	"context"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"
	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordBySmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordBySmsLogic {
	return &ResetPasswordBySmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ResetPasswordBySms resets user password via SMS verification
func (l *ResetPasswordBySmsLogic) ResetPasswordBySms(in *core.ResetPasswordBySmsReq) (*core.BaseResp, error) {
	// Find user by phone number
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.MobileEQ(in.PhoneNumber)).
		Only(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("login.phoneNotFound")
	}

	// Note: SMS captcha verification happens in API layer

	// Hash new password
	hashedPassword, err := encrypt.BcryptEncrypt(in.Password)
	if err != nil {
		return nil, errorx.NewInternalError("common.encryptionFailed")
	}

	// Update password
	err = l.svcCtx.DB.User.UpdateOne(userInfo).
		SetPassword(hashedPassword).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{
		Msg: "login.resetPasswordSuccessTitle",
	}, nil
}
