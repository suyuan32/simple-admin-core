package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryDetailLogic {
	return &CreateOrUpdateDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateDictionaryDetailLogic) CreateOrUpdateDictionaryDetail(req *types.CreateOrUpdateDictionaryDetailReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateDictionaryDetail(l.ctx, &core.DictionaryDetail{
		Id:           req.ID,
		Title:        req.Title,
		Key:          req.Key,
		Value:        req.Value,
		Status:       req.Status,
		DictionaryId: req.ParentID,
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
