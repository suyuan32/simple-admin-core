package role

import (
	"context"
	"fmt"
	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/entx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *core.RoleInfo) (*core.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		origin, err := tx.Role.Get(l.ctx, *in.Id)
		if err != nil {
			return err
		}

		err = tx.Role.UpdateOneID(*in.Id).
			SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
			SetNotNilName(in.Name).
			SetNotNilCode(in.Code).
			SetNotNilDefaultRouter(in.DefaultRouter).
			SetNotNilRemark(in.Remark).
			SetNotNilSort(in.Sort).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		if in.Code != nil && origin.Code != *in.Code {
			_, err = tx.QueryContext(l.ctx, fmt.Sprintf("update casbin_rules set v0='%s' WHERE v0='%s'", *in.Code, origin.Code))
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
