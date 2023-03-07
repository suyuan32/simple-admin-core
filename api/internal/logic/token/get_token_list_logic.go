package token

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetTokenListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
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
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
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
