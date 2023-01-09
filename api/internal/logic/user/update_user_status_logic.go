package user

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewUpdateUserStatusLogic(r *http.Request, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(req *types.StatusCodeUUIDReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.UpdateUserStatus(l.ctx, &core.StatusCodeUUIDReq{
		Id:     req.Id,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
