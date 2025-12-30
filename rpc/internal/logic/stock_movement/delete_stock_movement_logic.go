package stock_movement

import (
	"context"

    "github.com/suyuan32/simple-admin-core/rpc/ent/stockmovement"
    "github.com/suyuan32/simple-admin-core/rpc/internal/svc"
    "github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteStockMovementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStockMovementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStockMovementLogic {
	return &DeleteStockMovementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStockMovementLogic) DeleteStockMovement(in *core.UUIDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.StockMovement.Delete().Where(stockmovement.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
