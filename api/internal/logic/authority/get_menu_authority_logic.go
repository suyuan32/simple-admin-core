package authority

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMenuAuthorityLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(req *types.IDReq) (resp *types.MenuAuthorityInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuAuthority(l.ctx, &core.IDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.MenuAuthorityInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Msg: l.svcCtx.Trans.Trans(l.lang, i18n.Success),
		},
		Data: types.MenuAuthorityInfoReq{
			RoleId:  req.Id,
			MenuIds: data.MenuId,
		},
	}

	return resp, nil
}
