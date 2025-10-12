package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(req *types.DictionaryInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateDictionary(l.ctx,
		&core.DictionaryInfo{
			Id:       req.Id,
			Title:    req.Title,
			Name:     req.Name,
			Status:   req.Status,
			Desc:     req.Desc,
			IsPublic: req.IsPublic,
		})
	if err != nil {
		return nil, err
	}

	dict, err := NewGetDictionaryByIdLogic(l.ctx, l.svcCtx).GetDictionaryById(&types.IDReq{Id: *req.Id})
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.Redis.Del(l.ctx, "DICTIONARY:PUBLIC_STATE:"+*dict.Data.Name).Err(); err != nil {
		logx.Errorw("failed to delete dictionary data in redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
