package position

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdatePositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdatePositionLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdatePositionLogic {
	return &CreateOrUpdatePositionLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdatePositionLogic) CreateOrUpdatePosition(req *types.CreateOrUpdatePositionReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdatePosition(l.ctx,
		&core.PositionInfo{
			Id:     req.Id,
			Status: req.Status,
			Sort:   req.Sort,
			Name:   req.Name,
			Code:   req.Code,
			Remark: req.Remark,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
