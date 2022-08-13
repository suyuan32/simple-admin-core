package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeMenuParamListByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeMenuParamListByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeMenuParamListByIdLogic {
	return &GeMenuParamListByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeMenuParamListByIdLogic) GeMenuParamListById(in *core.IDReq) (*core.MenuParamListResp, error) {
	// todo: add your logic here and delete this line

	return &core.MenuParamListResp{}, nil
}
