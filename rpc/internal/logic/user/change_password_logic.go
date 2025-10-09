package user

import (
	"context"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ChangePassword allows authenticated user to change their password
func (l *ChangePasswordLogic) ChangePassword(in *core.ChangePasswordReq) (*core.BaseResp, error) {
	// Get user ID from context (set by API layer JWT middleware)
	userId, ok := l.ctx.Value("userId").(string)
	if !ok || userId == "" {
		return nil, errorx.NewCodeUnauthenticatedError("common.unauthorized")
	}

	// Parse UUID
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("common.invalidUserId")
	}

	// Get user from database
	userInfo, err := l.svcCtx.DB.User.Get(l.ctx, userUUID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// Verify old password
	if !encrypt.BcryptCheck(in.OldPassword, userInfo.Password) {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongPassword")
	}

	// Hash new password
	hashedPassword, err := encrypt.BcryptEncrypt(in.NewPassword)
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
		Msg: "login.changePasswordSuccessTitle",
	}, nil
}
