package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.ProfileReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.UpdateProfile(l.ctx, &core.UpdateProfileReq{
		Uuid:     l.ctx.Value("userId").(string),
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
