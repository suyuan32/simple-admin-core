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

type ResetPasswordByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordByEmailLogic {
	return &ResetPasswordByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ResetPasswordByEmail resets user password via email verification
func (l *ResetPasswordByEmailLogic) ResetPasswordByEmail(in *core.ResetPasswordByEmailReq) (*core.BaseResp, error) {
	// Find user by email
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.EmailEQ(in.Email)).
		Only(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("login.emailNotFound")
	}

	// Note: Email captcha verification happens in API layer

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
