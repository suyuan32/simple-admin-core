package publicapi

import (
	"context"
	"encoding/json"
	"time"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicDictionaryDetailByDictionaryNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicDictionaryDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicDictionaryDetailByDictionaryNameLogic {
	return &GetPublicDictionaryDetailByDictionaryNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicDictionaryDetailByDictionaryNameLogic) GetPublicDictionaryDetailByDictionaryName(req *types.DictionaryNameReq) (resp *types.DictionaryDetailListResp, err error) {
	if val, err := l.svcCtx.Redis.Get(l.ctx, "DICTIONARY:PUBLIC_STATE:"+*req.Name).Result(); err == nil && val == "false" {
		return nil, errorx.NewCodeUnavailableError("this dictionary cannot be viewed by public")
	} else if err != nil {
		// cache dict public state
		dictData, err1 := l.svcCtx.CoreRpc.GetDictionaryList(l.ctx, &core.DictionaryListReq{Name: req.Name, Page: 1, PageSize: 100})
		if err1 != nil {
			return nil, err1
		}

		for _, v := range dictData.Data {
			if *v.Name == *req.Name {
				publicState := "false"

				if *v.IsPublic && *v.Status == 1 {
					publicState = "true"
				}

				err1 = l.svcCtx.Redis.Set(l.ctx, "DICTIONARY:PUBLIC_STATE:"+*req.Name, publicState, 24*time.Hour).Err()
				if err1 != nil {
					logx.Errorw("failed to set dictionary public state data to redis", logx.Field("detail", err1))
					return nil, errorx.NewCodeInternalError(i18n.RedisError)
				}

				if !*v.IsPublic {
					return nil, errorx.NewCodeUnavailableError("this dictionary cannot be viewed by public")
				}
			}
		}
	}

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
