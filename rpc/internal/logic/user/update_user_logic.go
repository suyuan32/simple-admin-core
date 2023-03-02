package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *core.UserInfo) (*core.BaseResp, error) {
	err := utils.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		updateQuery := tx.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
			SetNotEmptyUsername(in.Username).
			SetNotEmptyNickname(in.Nickname).
			SetNotEmptyEmail(in.Email).
			SetNotEmptyMobile(in.Mobile).
			SetNotEmptyAvatar(in.Avatar).
			SetNotEmptyHomePath(in.HomePath).
			SetNotEmptyDescription(in.Description).
			SetNotEmptyDepartmentID(in.DepartmentId)

		if in.Password != "" {
			updateQuery = updateQuery.SetNotEmptyPassword(utils.BcryptEncrypt(in.Password))
		}

		if in.RoleIds != nil {
			err := l.svcCtx.DB.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).ClearRoles().Exec(l.ctx)
			if err != nil {
				return err
			}

			updateQuery = updateQuery.AddRoleIDs(in.RoleIds...)
		}

		if in.PositionIds != nil {
			err := l.svcCtx.DB.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).ClearPositions().Exec(l.ctx)
			if err != nil {
				return err
			}

			updateQuery = updateQuery.AddPositionIDs(in.PositionIds...)
		}

		return updateQuery.Exec(l.ctx)
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{
		Msg: i18n.Success,
	}, nil
}
