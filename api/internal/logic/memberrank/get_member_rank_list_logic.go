package memberrank

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type GetMemberRankListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMemberRankListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMemberRankListLogic {
	return &GetMemberRankListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMemberRankListLogic) GetMemberRankList(req *types.MemberRankListReq) (resp *types.MemberRankListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMemberRankList(l.ctx,
		&core.MemberRankListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			Name:        req.Name,
			Description: req.Description,
			Remark:      req.Remark,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MemberRankListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MemberRankInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Trans:       l.svcCtx.Trans.Trans(l.lang, v.Name),
				Name:        v.Name,
				Description: v.Description,
				Remark:      v.Remark,
			})
	}
	return resp, nil
}
