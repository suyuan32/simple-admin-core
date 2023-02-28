package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type DeleteDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Dictionary.Delete().Where(dictionary.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
