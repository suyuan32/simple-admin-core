package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuParamLogic {
	return &CreateOrUpdateMenuParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateMenuParamLogic) CreateOrUpdateMenuParam(req *types.CreateOrUpdateMenuParamReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateMenuParam(l.ctx, &core.CreateOrUpdateMenuParamReq{
		Id:     req.Id,
		MenuId: req.MenuId,
		Type:   req.DataType,
		Key:    req.Key,
		Value:  req.Value,
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
