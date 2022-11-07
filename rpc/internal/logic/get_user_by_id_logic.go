package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *core.UUIDReq) (*core.UserInfoResp, error) {
	u, err := l.svcCtx.DB.User.Query().Where(user.UUIDEQ(in.Uuid)).First(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
		default:
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}
	}

	roleName, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleID))
	roleValue, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d_value", u.RoleID))
	if err != nil {
		return nil, err
	}

	return &core.UserInfoResp{
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		RoleId:    u.RoleID,
		RoleName:  roleName,
		RoleValue: roleValue,
		Mobile:    u.Mobile,
		Email:     u.Email,
		Status:    uint64(u.Status),
		Id:        u.ID,
		Username:  u.Username,
		Uuid:      u.UUID,
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}, nil
}
