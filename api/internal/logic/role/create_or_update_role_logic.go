package role

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateRoleLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateRoleLogic {
	return &CreateOrUpdateRoleLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateRoleLogic) CreateOrUpdateRole(req *types.RoleInfo) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateRole(l.ctx, &core.RoleInfo{
		Id:            req.Id,
		Name:          req.Name,
		Value:         req.Value,
		DefaultRouter: req.DefaultRouter,
		Status:        uint64(req.Status),
		Remark:        req.Remark,
		OrderNo:       req.OrderNo,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
