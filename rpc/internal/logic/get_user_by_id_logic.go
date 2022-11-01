package logic

import (
	"context"
	"fmt"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var u model.User
	result := l.svcCtx.DB.Where("uuid = ?", in.UUID).First(&u)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		logx.Errorw("user dose not find, please check the UUID", logx.Field("UUID", in.UUID))
		return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
	}
	roleName, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleId))
	roleValue, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d_value", u.RoleId))
	if err != nil {
		return nil, err
	}
	return &core.UserInfoResp{
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		RoleId:    u.RoleId,
		RoleName:  roleName,
		RoleValue: roleValue,
		Mobile:    u.Mobile,
		Email:     u.Email,
		Status:    u.Status,
		Id:        uint64(u.ID),
		Username:  u.Username,
		UUID:      u.UUID,
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}, nil
}
