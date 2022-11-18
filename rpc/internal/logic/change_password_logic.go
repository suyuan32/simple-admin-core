package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *core.ChangePasswordReq) (*core.BaseResp, error) {
	target, err := l.svcCtx.DB.User.Query().Where(user.UUIDEQ(in.Uuid)).First(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("uuid", in.Uuid))
			return nil, statuserr.NewInvalidArgumentError("login.userNotExist")
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	if ok := utils.BcryptCheck(in.OldPassword, target.Password); ok {
		newPassword := utils.BcryptEncrypt(in.NewPassword)

		err = l.svcCtx.DB.User.Update().Where(user.UUIDEQ(in.Uuid)).SetPassword(newPassword).Exec(l.ctx)
		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("uuid", in.Uuid))
				return nil, statuserr.NewInvalidArgumentError("login.userNotExist")
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}
	} else {
		logx.Errorw("old password is wrong", logx.Field("UUID", in.Uuid))
		return nil, statuserr.NewInvalidArgumentError("user.wrongPassword")
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
