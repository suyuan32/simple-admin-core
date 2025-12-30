package warehouse

import (
	"context"

    "github.com/suyuan32/simple-admin-core/rpc/ent/warehouse"
    "github.com/suyuan32/simple-admin-core/rpc/internal/svc"
    "github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteWarehouseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteWarehouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWarehouseLogic {
	return &DeleteWarehouseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteWarehouseLogic) DeleteWarehouse(in *core.UUIDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Warehouse.Delete().Where(warehouse.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
