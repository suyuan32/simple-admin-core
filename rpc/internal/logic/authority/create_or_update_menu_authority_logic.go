package authority

import (
	"context"

	"github.com/chimerakang/simple-admin-common/i18n"

	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/entx"

	"github.com/chimerakang/simple-admin-core/rpc/ent"

	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		err := tx.Role.UpdateOneID(in.RoleId).ClearMenus().Exec(l.ctx)
		if err != nil {
			return err
		}

		err = tx.Role.UpdateOneID(in.RoleId).AddMenuIDs(in.MenuIds...).Exec(l.ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
