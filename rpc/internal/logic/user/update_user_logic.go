package user

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/entx"

	"github.com/suyuan32/simple-admin-core/rpc/ent"

	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/token"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
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
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
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
			updateQuery = updateQuery.SetNotEmptyPassword(encrypt.BcryptEncrypt(in.Password))
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

			_, err = token.NewBlockUserAllTokenLogic(l.ctx, l.svcCtx).BlockUserAllToken(&core.UUIDReq{Id: in.Id})
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
