package dictionarydetail

import (
	"context"
	"encoding/json"
	"time"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailByDictionaryNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailByDictionaryNameLogic {
	return &GetDictionaryDetailByDictionaryNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryDetailByDictionaryNameLogic) GetDictionaryDetailByDictionaryName(req *types.DictionaryNameReq) (resp *types.DictionaryDetailListResp, err error) {
	if val, err := l.svcCtx.Redis.Get(l.ctx, "DICTIONARY:"+*req.Name).Result(); err == nil && val != "" {
		resp = &types.DictionaryDetailListResp{}
		resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
		err = json.Unmarshal([]byte(val), &resp.Data)
		if err != nil {
			logx.Errorw("failed to unmarshal the dictionary data from redis",
				logx.Field("detail", err), logx.Field("dictionaryName", req.Name))
			return nil, errorx.NewCodeInternalError(i18n.Failed)
		}

		for _, v := range resp.Data.Data {
			v.Trans = l.svcCtx.Trans.Trans(l.ctx, *v.Title)
		}
		return resp, nil
	}

	data, err := l.svcCtx.CoreRpc.GetDictionaryDetailByDictionaryName(l.ctx, &core.BaseMsg{Msg: *req.Name})
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
				Trans:        l.svcCtx.Trans.Trans(l.ctx, *v.Title),
			})
	}

	storeData, err := json.Marshal(&resp.Data)
	err = l.svcCtx.Redis.Set(l.ctx, "DICTIONARY:"+*req.Name, string(storeData), 24*time.Hour).Err()
	if err != nil {
		logx.Errorw("failed to set dictionary detail data to redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	return resp, nil
}
