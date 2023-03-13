package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/position"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *core.UserListReq) (*core.UserListResp, error) {
	var predicates []predicate.User

	if in.Mobile != "" {
		predicates = append(predicates, user.MobileEQ(in.Mobile))
	}

	if in.Username != "" {
		predicates = append(predicates, user.UsernameContains(in.Username))
	}

	if in.Email != "" {
		predicates = append(predicates, user.EmailEQ(in.Email))
	}

	if in.Nickname != "" {
		predicates = append(predicates, user.NicknameContains(in.Nickname))
	}

	if in.RoleIds != nil {
		predicates = append(predicates, user.HasRolesWith(role.IDIn(in.RoleIds...)))
	}

	if in.DepartmentId != 0 {
		predicates = append(predicates, user.DepartmentIDEQ(in.DepartmentId))
	}

	if in.PositionIds != nil {
		predicates = append(predicates, user.HasPositionsWith(position.IDIn(in.PositionIds...)))
	}

	users, err := l.svcCtx.DB.Debug().User.Query().Where(predicates...).WithRoles().WithPositions().Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.UserListResp{}
	resp.Total = users.PageDetails.Total

	for _, v := range users.List {
		resp.Data = append(resp.Data, &core.UserInfo{
			Id:           v.ID.String(),
			Avatar:       v.Avatar,
			RoleIds:      GetRoleIds(v.Edges.Roles),
			Mobile:       v.Mobile,
			Email:        v.Email,
			Status:       uint32(v.Status),
			Username:     v.Username,
			Nickname:     v.Nickname,
			HomePath:     v.HomePath,
			Description:  v.Description,
			DepartmentId: v.DepartmentID,
			PositionIds:  GetPositionIds(v.Edges.Positions),
			CreatedAt:    v.CreatedAt.UnixMilli(),
			UpdatedAt:    v.UpdatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
