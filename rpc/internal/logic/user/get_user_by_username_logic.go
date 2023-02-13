package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *core.UsernameReq) (*core.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Query().Where(user.UsernameEQ(in.Username)).WithRoles().First(l.ctx)
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

	if err != nil {
		return nil, err
	}

	return &core.UserInfo{
		Nickname:     result.Nickname,
		Avatar:       result.Avatar,
		Password:     result.Password,
		RoleIds:      GetRoleIds(result.Edges.Roles),
		RoleCodes:    GetRoleCodes(result.Edges.Roles),
		Mobile:       result.Mobile,
		Email:        result.Email,
		Status:       uint32(result.Status),
		Id:           result.ID.String(),
		Username:     result.Username,
		HomePath:     result.HomePath,
		Description:  result.Description,
		DepartmentId: result.DepartmentID,
		CreatedAt:    result.CreatedAt.Unix(),
		UpdatedAt:    result.UpdatedAt.Unix(),
	}, nil
}
