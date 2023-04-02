package dictionarydetail

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetDictionaryDetailListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailListLogic {
	return &GetDictionaryDetailListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryDetailListLogic) GetDictionaryDetailList(req *types.DictionaryDetailListReq) (resp *types.DictionaryDetailListResp, err error) {
	if val, err := l.svcCtx.Redis.GetCtx(l.ctx, fmt.Sprintf("dict_%d", req.DictionaryId)); err == nil && val != "" {
		resp = &types.DictionaryDetailListResp{}
		resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
		err = json.Unmarshal([]byte(val), &resp.Data)
		if err != nil {
			logx.Errorw("failed to unmarshal the dictionary data from redis",
				logx.Field("detail", err), logx.Field("dictionaryId", req.DictionaryId))
			return nil, errorx.NewCodeInternalError(i18n.Failed)
		}
		return resp, nil
	}

	data, err := l.svcCtx.CoreRpc.GetDictionaryDetailList(l.ctx,
		&core.DictionaryDetailListReq{
			Page:         req.Page,
			PageSize:     req.PageSize,
			DictionaryId: req.DictionaryId,
			Key:          req.Key,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.DictionaryDetailListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.DictionaryDetailInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:       v.Status,
				Title:        v.Title,
				Key:          v.Key,
				Value:        v.Value,
				DictionaryId: v.DictionaryId,
				Sort:         v.Sort,
			})
	}

	storeData, err := json.Marshal(&resp.Data)
	err = l.svcCtx.Redis.SetexCtx(l.ctx, fmt.Sprintf("dict_%d", req.DictionaryId), string(storeData), 3600)
	if err != nil {
		logx.Errorw("failed to set dictionary detail data to redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	return resp, nil
}
