package menu

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateMenuParamLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuParamLogic {
	return &CreateOrUpdateMenuParamLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateMenuParamLogic) CreateOrUpdateMenuParam(req *types.CreateOrUpdateMenuParamReq) (resp *types.BaseMsgResp, err error) {
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

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
