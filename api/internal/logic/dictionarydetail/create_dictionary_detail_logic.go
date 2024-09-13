package dictionarydetail

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryDetailLogic {
	return &CreateDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDictionaryDetailLogic) CreateDictionaryDetail(req *types.DictionaryDetailInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateDictionaryDetail(l.ctx,
		&core.DictionaryDetailInfo{
			Status:       req.Status,
			Title:        req.Title,
			Key:          req.Key,
			Value:        req.Value,
			DictionaryId: req.DictionaryId,
			Sort:         req.Sort,
		})
	if err != nil {
		return nil, err
	}

	dict, err := dictionary.NewGetDictionaryByIdLogic(l.ctx, l.svcCtx).GetDictionaryById(&types.IDReq{Id: *req.DictionaryId})
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.Redis.Del(l.ctx, "DICTIONARY:"+*dict.Data.Name).Err(); err != nil {
		logx.Errorw("failed to delete dictionary data in redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
