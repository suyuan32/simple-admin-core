package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuByIdLogic {
	return &GetMenuByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuByIdLogic) GetMenuById(req *types.IDReq) (resp *types.MenuInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
