package menuparam

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuParamLogic {
	return &CreateMenuParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuParamLogic) CreateMenuParam(req *types.MenuParamInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateMenuParam(l.ctx,
		&core.MenuParamInfo{
			Id:    req.Id,
			Type:  req.Type,
			Key:   req.Key,
			Value: req.Value,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
