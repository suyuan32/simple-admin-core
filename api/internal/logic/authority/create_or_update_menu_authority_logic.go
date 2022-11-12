package authority

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateMenuAuthorityLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(req *types.MenuAuthorityInfoReq) (resp *types.BaseMsgResp, err error) {
	authority, err := l.svcCtx.CoreRpc.CreateOrUpdateMenuAuthority(l.ctx, &core.RoleMenuAuthorityReq{
		RoleId: req.RoleId,
		MenuId: req.MenuIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, authority.Msg)}, nil
}
