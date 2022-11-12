package authority

import (
	"context"
	"net/http"
	"strconv"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetApiAuthorityLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetApiAuthorityLogic {
	return &GetApiAuthorityLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetApiAuthorityLogic) GetApiAuthority(req *types.IDReq) (resp *types.ApiAuthorityListResp, err error) {
	roleId := strconv.Itoa(int(req.Id))
	data := l.svcCtx.Casbin.GetFilteredPolicy(0, roleId)
	resp = &types.ApiAuthorityListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = uint64(len(data))
	for _, v := range data {
		resp.Data.Data = append(resp.Data.Data, types.ApiAuthorityInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return resp, nil
}
