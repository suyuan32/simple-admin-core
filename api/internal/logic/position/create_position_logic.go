package position

import (
	"context"

	"github.com/chimerakang/simple-admin-core/api/internal/svc"
	"github.com/chimerakang/simple-admin-core/api/internal/types"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePositionLogic {
	return &CreatePositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePositionLogic) CreatePosition(req *types.PositionInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreatePosition(l.ctx,
		&core.PositionInfo{
			Status: req.Status,
			Sort:   req.Sort,
			Name:   req.Name,
			Code:   req.Code,
			Remark: req.Remark,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
