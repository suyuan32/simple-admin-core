package menuparam

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuParamLogic {
	return &UpdateMenuParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuParamLogic) UpdateMenuParam(req *types.MenuParamInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateMenuParam(l.ctx,
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
