package logic

import (
	"context"

	"github.com/google/uuid"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
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
	if in.Id == 0 {
		err := l.svcCtx.DB.User.Create().
			SetUUID(uuid.NewString()).
			SetUsername(in.Username).
			SetPassword(utils.BcryptEncrypt(in.Password)).
			SetNickname(in.Email).
			SetEmail(in.Email).
			SetMobile(in.Mobile).
			SetAvatar(in.Avatar).
			SetRoleID(in.RoleId).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.CreateFailed)
			default:
				logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: errorx.Success}, nil
	} else {
		err := l.svcCtx.DB.User.UpdateOneID(in.Id).
			SetUUID(uuid.NewString()).
			SetUsername(in.Username).
			SetPassword(utils.BcryptEncrypt(in.Password)).
			SetNickname(in.Email).
			SetEmail(in.Email).
			SetMobile(in.Mobile).
			SetAvatar(in.Avatar).
			SetRoleID(in.RoleId).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.UpdateFailed)
			default:
				logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		return &core.BaseResp{
			Msg: errorx.Success,
		}, nil
	}
}
