package user

import (
	"context"

	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserProfile updates user's profile information (nickname, avatar, mobile, email, locale)
func (l *UpdateUserProfileLogic) UpdateUserProfile(in *core.ProfileInfo) (*core.BaseResp, error) {
	// Get user ID from context
	userId, ok := l.ctx.Value("userId").(string)
	if !ok || userId == "" {
		return nil, errorx.NewCodeUnauthenticatedError("common.unauthorized")
	}

	// Parse UUID
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("common.invalidUserId")
	}

	// Get user
	userInfo, err := l.svcCtx.DB.User.Get(l.ctx, userUUID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// Build update query
	updateBuilder := l.svcCtx.DB.User.UpdateOne(userInfo)

	if in.Nickname != nil {
		updateBuilder = updateBuilder.SetNickname(*in.Nickname)
	}
	if in.Avatar != nil {
		updateBuilder = updateBuilder.SetAvatar(*in.Avatar)
	}
	if in.Mobile != nil {
		updateBuilder = updateBuilder.SetMobile(*in.Mobile)
	}
	if in.Email != nil {
		updateBuilder = updateBuilder.SetEmail(*in.Email)
	}
	if in.Locale != nil {
		updateBuilder = updateBuilder.SetLocale(*in.Locale)
	}

	// Execute update
	err = updateBuilder.Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{
		Msg: "common.updateSuccess",
	}, nil
}
