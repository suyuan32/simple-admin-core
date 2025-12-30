package stock_movement

import (
	"context"
	"errors"
	"fmt"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/ent/inventory"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStockMovementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStockMovementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStockMovementLogic {
	return &CreateStockMovementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStockMovementLogic) CreateStockMovement(in *core.StockMovementInfo) (*core.BaseUUIDResp, error) {
	if in.Quantity == nil || *in.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}
	qty := *in.Quantity

	if in.ProductId == nil {
		return nil, errors.New("product_id is required")
	}
	prodID := uuidx.ParseUUIDString(*in.ProductId)

	if in.Reference == nil || *in.Reference == "" {
		return nil, errors.New("reference is required")
	}

	tx, err := l.svcCtx.DB.Tx(l.ctx)
	if err != nil {
		return nil, err
	}

	var moveID string

	err = func() error {
		movementType := "IN"
		if in.MovementType != nil {
			movementType = *in.MovementType
		}

		switch movementType {
		case "IN":
			if in.ToWarehouseId == nil {
				return errors.New("to_warehouse_id is required for IN")
			}
			warehouseID := uuidx.ParseUUIDString(*in.ToWarehouseId)

			// Check if inventory exists
			inv, err := tx.Inventory.Query().
				Where(inventory.ProductIDEQ(prodID), inventory.WarehouseIDEQ(warehouseID)).
				First(l.ctx)
			if ent.IsNotFound(err) {
				_, err = tx.Inventory.Create().
					SetProductID(prodID).
					SetWarehouseID(warehouseID).
					SetQuantity(qty).
					Save(l.ctx)
			} else if err != nil {
				return err
			} else {
				_, err = inv.Update().AddQuantity(qty).Save(l.ctx)
			}
			if err != nil {
				return err
			}

		case "OUT":
			if in.FromWarehouseId == nil {
				return errors.New("from_warehouse_id is required for OUT")
			}
			warehouseID := uuidx.ParseUUIDString(*in.FromWarehouseId)

			inv, err := tx.Inventory.Query().
				Where(inventory.ProductIDEQ(prodID), inventory.WarehouseIDEQ(warehouseID)).
				First(l.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					return errors.New("insufficient stock: inventory record not found")
				}
				return err
			}
			if inv.Quantity < qty {
				return errors.New("insufficient stock")
			}
			_, err = inv.Update().AddQuantity(-qty).Save(l.ctx)
			if err != nil {
				return err
			}

		case "MOVE":
			if in.FromWarehouseId == nil || in.ToWarehouseId == nil {
				return errors.New("from_warehouse_id and to_warehouse_id are required for MOVE")
			}
			fromID := uuidx.ParseUUIDString(*in.FromWarehouseId)
			toID := uuidx.ParseUUIDString(*in.ToWarehouseId)

			// Out from source
			invFrom, err := tx.Inventory.Query().
				Where(inventory.ProductIDEQ(prodID), inventory.WarehouseIDEQ(fromID)).
				First(l.ctx)
			if err != nil {
				return err
			}
			if invFrom.Quantity < qty {
				return errors.New("insufficient stock")
			}
			_, err = invFrom.Update().AddQuantity(-qty).Save(l.ctx)
			if err != nil {
				return err
			}

			// In to dest
			invTo, err := tx.Inventory.Query().
				Where(inventory.ProductIDEQ(prodID), inventory.WarehouseIDEQ(toID)).
				First(l.ctx)
			if ent.IsNotFound(err) {
				_, err = tx.Inventory.Create().
					SetProductID(prodID).
					SetWarehouseID(toID).
					SetQuantity(qty).
					Save(l.ctx)
			} else if err != nil {
				return err
			} else {
				_, err = invTo.Update().AddQuantity(qty).Save(l.ctx)
			}
			if err != nil {
				return err
			}

		default:
			return errors.New("invalid movement type")
		}

		// Create Stock Movement
		result, err := tx.StockMovement.Create().
			SetProductID(prodID).
			SetNillableFromWarehouseID(uuidx.ParseUUIDStringToPointer(in.FromWarehouseId)).
			SetNillableToWarehouseID(uuidx.ParseUUIDStringToPointer(in.ToWarehouseId)).
			SetQuantity(qty).
			SetMovementType(movementType).
			SetReference(*in.Reference).
			SetNillableDetails(in.Details).
			Save(l.ctx)

		if err != nil {
			return err
		}
		moveID = result.ID.String()
		return nil
	}()

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if err := tx.Commit(); err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: moveID, Msg: i18n.CreateSuccess}, nil
}
