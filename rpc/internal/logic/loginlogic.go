package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/suyuan32/simple-admin-core/api/common/errorx"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  user service
func (l *LoginLogic) Login(in *core.LoginReq) (*core.LoginResp, error) {
	var u model.User
	result := l.svcCtx.DB.Where(&model.User{Username: in.Username}).First(&u)
	if result.Error != nil {
		l.Logger.Error("login logic: database error ", result.Error)
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, errorx.UserNotExists)
	}

	if ok := util.BcryptCheck(in.Password, u.Password); !ok {
		return nil, status.Error(codes.InvalidArgument, errorx.WrongUsernameOrPassword)
	}

	// get role data from redis
	var roleName, value string
	if s, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleId)); err != nil {
		var roleData []model.Role
		res := l.svcCtx.DB.Find(&roleData)
		if res.RowsAffected == 0 {
			return nil, status.Error(codes.NotFound, "role not found")
		}
		for _, v := range roleData {
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", v.ID), v.Name)
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
			if err != nil {
				return nil, status.Error(codes.Internal, errorx.RedisError)
			}
			if v.ID == uint(u.RoleId) {
				roleName = v.Name
				value = v.Value
			}
		}
	} else {
		roleName = s
		value, err = l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d_value", u.RoleId))
		if err != nil {
			return nil, errorx.NewRpcError(codes.NotFound, errorx.TargetNotExist)
		}
	}

	return &core.LoginResp{
		Id:        u.UUID,
		RoleValue: value,
		RoleName:  roleName,
		RoleId:    u.RoleId,
	}, nil

}
