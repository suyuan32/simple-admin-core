package token

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type GetTokenListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetTokenListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetTokenListLogic) GetTokenList(req *types.TokenListReq) (resp *types.TokenListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetTokenList(l.ctx,
		&core.TokenListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Username: req.Username,
			Nickname: req.Nickname,
			Email:    req.Email,
			Uuid:     req.Uuid,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.TokenListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.TokenInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:    v.Status,
				Uuid:      v.Uuid,
				Token:     v.Token,
				Source:    v.Source,
				ExpiredAt: v.ExpiredAt,
			})
	}
	return resp, nil
}
