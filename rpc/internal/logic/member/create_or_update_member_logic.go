package member

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

type CreateOrUpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMemberLogic {
	return &CreateOrUpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMemberLogic) CreateOrUpdateMember(in *core.MemberInfo) (*core.BaseResp, error) {
	if in.Id == "" {
		err := l.svcCtx.DB.Member.Create().
			SetStatus(uint8(in.Status)).
			SetUsername(in.Username).
			SetPassword(utils.BcryptEncrypt(in.Password)).
			SetNickname(in.Nickname).
			SetRankID(in.RankId).
			SetMobile(in.Mobile).
			SetEmail(in.Email).
			SetAvatar(in.Avatar).
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

		return &core.BaseResp{Msg: i18n.CreateSuccess}, nil
	} else {
		updateQuery := l.svcCtx.DB.Member.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
			SetStatus(uint8(in.Status)).
			SetUsername(in.Username).
			SetNickname(in.Nickname).
			SetRankID(in.RankId).
			SetMobile(in.Mobile).
			SetEmail(in.Email).
			SetAvatar(in.Avatar)

		if in.Password != "" {
			updateQuery = updateQuery.SetPassword(utils.BcryptEncrypt(in.Password))
		}

		err := updateQuery.Exec(l.ctx)
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

		return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
	}
}
