package dictionary

import (
	"context"

	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/entx"

	"github.com/chimerakang/simple-admin-core/rpc/ent"
	"github.com/chimerakang/simple-admin-core/rpc/ent/dictionary"
	"github.com/chimerakang/simple-admin-core/rpc/ent/dictionarydetail"

	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/chimerakang/simple-admin-common/i18n"
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
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		_, txErr := tx.DictionaryDetail.Delete().Where(dictionarydetail.HasDictionariesWith(dictionary.IDIn(in.Ids...))).Exec(l.ctx)
		if txErr != nil {
			return txErr
		}

		_, txErr = tx.Dictionary.Delete().Where(dictionary.IDIn(in.Ids...)).Exec(l.ctx)
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
