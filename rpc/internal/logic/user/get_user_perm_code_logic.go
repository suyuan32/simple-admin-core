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

type GetUserPermCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserPermCode returns all permission codes for the authenticated user
func (l *GetUserPermCodeLogic) GetUserPermCode(in *core.Empty) (*core.PermCodeResp, error) {
	// Get user ID from context
	userId, ok := l.ctx.Value("userId").(string)
	if !ok || userId == "" {
		return nil, errorx.NewCodeUnauthenticatedError("common.unauthorized")
	}

	// Parse UUID
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("common.invalidUserId")
	}

	// Query user with roles and menus
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.IDEQ(userUUID)).
		WithRoles(func(q *ent.RoleQuery) {
			q.WithMenus() // Load menus associated with roles
		}).
		Only(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// Collect all permission codes from menus
	permCodesMap := make(map[string]bool)
	for _, role := range userInfo.Edges.Roles {
		if role.Edges.Menus != nil {
			for _, menu := range role.Edges.Menus {
				if menu.Permission != "" {
					permCodesMap[menu.Permission] = true
				}
			}
		}
	}

	// Convert map to slice
	permCodes := make([]string, 0, len(permCodesMap))
	for code := range permCodesMap {
		permCodes = append(permCodes, code)
	}

	return &core.PermCodeResp{
		Code: 0,
		Msg:  "common.success",
		Data: permCodes,
	}, nil
}
