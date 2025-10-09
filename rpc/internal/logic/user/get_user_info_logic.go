package user

import (
	"context"

	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserInfo returns basic user information with roles and department
func (l *GetUserInfoLogic) GetUserInfo(in *core.Empty) (*core.UserBaseIDInfoResp, error) {
	// Get user ID from context (set by API layer JWT middleware)
	userId, ok := l.ctx.Value("userId").(string)
	if !ok || userId == "" {
		return nil, errorx.NewCodeUnauthenticatedError("common.unauthorized")
	}

	// Parse UUID
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("common.invalidUserId")
	}

	// Query user with eager loading of roles and department
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.IDEQ(userUUID)).
		WithRoles().
		WithDepartment().
		Only(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// Build role names
	roleNames := make([]string, 0, len(userInfo.Edges.Roles))
	for _, role := range userInfo.Edges.Roles {
		roleNames = append(roleNames, role.Name)
	}

	// Build department name
	var deptName string
	if userInfo.Edges.Department != nil {
		deptName = userInfo.Edges.Department.Name
	}

	uuidStr := userInfo.ID.String()
	return &core.UserBaseIDInfoResp{
		Code: 0,
		Msg:  "common.success",
		Data: &core.UserBaseIDInfo{
			Uuid:           &uuidStr,
			Username:       &userInfo.Username,
			Nickname:       &userInfo.Nickname,
			Avatar:         &userInfo.Avatar,
			HomePath:       &userInfo.HomePath,
			Description:    &userInfo.Description,
			RoleName:       roleNames,
			DepartmentName: deptName,
			Locale:         &userInfo.Locale,
		},
	}, nil
}
