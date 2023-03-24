package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetDictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(req *types.DictionaryListReq) (resp *types.DictionaryListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetDictionaryList(l.ctx,
		&core.DictionaryListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.DictionaryListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.DictionaryInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Trans:  l.svcCtx.Trans.Trans(l.ctx, v.Title),
				Title:  v.Title,
				Name:   v.Name,
				Status: v.Status,
				Desc:   v.Desc,
			})
	}
	return resp, nil
}
