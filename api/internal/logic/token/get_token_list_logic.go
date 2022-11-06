package token

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
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
	result, err := l.svcCtx.CoreRpc.GetTokenList(l.ctx, &core.TokenListReq{
		Page: &core.PageInfoReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Uuid:     req.UUID,
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.TokenListResp{}
	resp.Total = result.Total

	for _, v := range result.Data {
		resp.Data = append(resp.Data, types.TokenInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			UUID:      v.Uuid,
			Token:     v.Token,
			Source:    v.Source,
			Status:    v.Status,
			ExpiredAt: v.ExpiredAt,
		})
	}

	return resp, nil
}
