package user

import (
	"context"

	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserProfile returns user's editable profile information
func (l *GetUserProfileLogic) GetUserProfile(in *core.Empty) (*core.ProfileResp, error) {
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

	// Query user
	userInfo, err := l.svcCtx.DB.User.Get(l.ctx, userUUID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.ProfileResp{
		Code: 0,
		Msg:  "common.success",
		Data: &core.ProfileInfo{
			Nickname: &userInfo.Nickname,
			Avatar:   &userInfo.Avatar,
			Mobile:   &userInfo.Mobile,
			Email:    &userInfo.Email,
			Locale:   &userInfo.Locale,
		},
	}, nil
}
