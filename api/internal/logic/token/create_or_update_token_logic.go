package token

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateTokenLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateTokenLogic {
	return &CreateOrUpdateTokenLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateTokenLogic) CreateOrUpdateToken(req *types.CreateOrUpdateTokenReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.UpdateRoleStatus(l.ctx, &core.StatusCodeReq{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
