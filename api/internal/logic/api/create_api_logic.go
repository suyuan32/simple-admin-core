package api

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiLogic) CreateApi(req *types.ApiInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateApi(l.ctx,
		&core.ApiInfo{
			Path:        req.Path,
			Description: req.Description,
			ApiGroup:    req.Group,
			Method:      req.Method,
			IsRequired:  req.IsRequired,
			ServiceName: req.ServiceName,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
