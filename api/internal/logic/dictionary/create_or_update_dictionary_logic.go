package dictionary

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateDictionaryLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryLogic {
	return &CreateOrUpdateDictionaryLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateDictionaryLogic) CreateOrUpdateDictionary(req *types.CreateOrUpdateDictionaryReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateDictionary(l.ctx, &core.DictionaryInfo{
		Id:     req.Id,
		Title:  req.Title,
		Name:   req.Name,
		Status: req.Status,
		Desc:   req.Description,
	})

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
