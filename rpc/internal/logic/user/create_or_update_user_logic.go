package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateUserLogic {
	return &CreateOrUpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateUserLogic) CreateOrUpdateUser(in *core.CreateOrUpdateUserReq) (*core.BaseResp, error) {
	if in.Id == "" {
		err := l.svcCtx.DB.User.Create().
			SetUsername(in.Username).
			SetPassword(utils.BcryptEncrypt(in.Password)).
			SetNickname(in.Nickname).
			SetEmail(in.Email).
			SetMobile(in.Mobile).
			SetAvatar(in.Avatar).
			SetRoleID(in.RoleId).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.CreateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: i18n.Success}, nil
	} else {
		var err error
		if in.Password != "" {
			err = l.svcCtx.DB.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
				SetUsername(in.Username).
				SetPassword(utils.BcryptEncrypt(in.Password)).
				SetNickname(in.Nickname).
				SetEmail(in.Email).
				SetMobile(in.Mobile).
				SetAvatar(in.Avatar).
				SetRoleID(in.RoleId).
				Exec(l.ctx)
		} else {
			err = l.svcCtx.DB.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
				SetUsername(in.Username).
				SetNickname(in.Nickname).
				SetEmail(in.Email).
				SetMobile(in.Mobile).
				SetAvatar(in.Avatar).
				SetRoleID(in.RoleId).
				Exec(l.ctx)
		}

		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		return &core.BaseResp{
			Msg: i18n.Success,
		}, nil
	}
}
