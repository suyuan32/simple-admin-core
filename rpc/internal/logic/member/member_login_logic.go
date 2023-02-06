package member

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/member"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogic {
	return &MemberLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogic) MemberLogin(in *core.LoginReq) (*core.MemberLoginResp, error) {
	result, err := l.svcCtx.DB.Member.Query().Where(member.UsernameEQ(in.Username)).First(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("user not found", logx.Field("username", in.Username))
			return nil, statuserr.NewInvalidArgumentError("login.userNotExist")
		}
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	if ok := utils.BcryptCheck(in.Password, result.Password); !ok {
		logx.Errorw("wrong password", logx.Field("detail", in))
		return nil, statuserr.NewInvalidArgumentError("login.wrongUsernameOrPassword")
	}

	return &core.MemberLoginResp{
		Id:       result.ID.String(),
		Nickname: result.Nickname,
		Avatar:   result.Avatar,
		RankId:   result.RankID,
	}, nil
}
