package user

import (
	"context"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateUserLogic {
	return &CreateOrUpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateUserLogic) CreateOrUpdateUser(req *types.CreateOrUpdateUserReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateUser(context.Background(), &core.CreateOrUpdateUserReq{
		Id:       uint64(req.Id),
		Avatar:   req.Avatar,
		RoleId:   req.RoleId,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Status:   req.Status,
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: data.Msg}, nil
}
