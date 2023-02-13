package member

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/member"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByUsernameLogic {
	return &GetMemberByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByUsernameLogic) GetMemberByUsername(in *core.UsernameReq) (*core.MemberInfo, error) {
	result, err := l.svcCtx.DB.Member.Query().Where(member.UsernameEQ(in.Username)).First(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("user not found", logx.Field("username", in.Username))
			return nil, statuserr.NewInvalidArgumentError("login.userNotExist")
		}
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	return &core.MemberInfo{
		Id:        result.ID.String(),
		CreatedAt: result.CreatedAt.UnixMilli(),
		UpdatedAt: result.UpdatedAt.UnixMilli(),
		Status:    uint32(result.Status),
		Username:  result.Username,
		Password:  result.Password,
		Nickname:  result.Nickname,
		RankId:    result.RankID,
		Mobile:    result.Mobile,
		Email:     result.Email,
		Avatar:    result.Avatar,
	}, nil
}
