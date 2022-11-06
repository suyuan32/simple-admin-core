package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
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

func (l *GetUserListLogic) GetUserList(in *core.GetUserListReq) (*core.UserListResp, error) {
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

	if in.RoleId != 0 {
		predicates = append(predicates, user.RoleIDEQ(in.RoleId))
	}

	users, err := l.svcCtx.DB.User.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(errorx.DatabaseError)
	}

	resp := &core.UserListResp{}
	resp.Total = users.PageDetails.Total

	for _, v := range users.List {
		resp.Data = append(resp.Data, &core.UserInfoResp{
			Id:        v.ID,
			Avatar:    v.Avatar,
			RoleId:    v.RoleID,
			Mobile:    v.Mobile,
			Email:     v.Email,
			Status:    uint64(v.Status),
			Username:  v.Username,
			Uuid:      v.UUID,
			Nickname:  v.Nickname,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
