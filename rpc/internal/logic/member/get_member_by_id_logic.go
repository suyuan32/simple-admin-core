package member

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
)

type GetMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByIdLogic {
	return &GetMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByIdLogic) GetMemberById(in *core.UUIDReq) (*core.MemberInfo, error) {
	result, err := l.svcCtx.DB.Member.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
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
