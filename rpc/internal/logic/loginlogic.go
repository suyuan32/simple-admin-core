package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/common/message"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
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

// user service
func (l *LoginLogic) Login(in *core.LoginReq) (*core.LoginResp, error) {
	var u model.User
	result := l.svcCtx.DB.Where(&model.User{Username: in.Username}).First(&u)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	if result.RowsAffected == 0 {
		logx.Errorw("User does not find", logx.Field("Username", in.Username))
		return nil, status.Error(codes.InvalidArgument, message.UserNotExists)
	}

	if ok := util.BcryptCheck(in.Password, u.Password); !ok {
		logx.Errorw("Wrong password", logx.Field("Detail", in))
		return nil, status.Error(codes.InvalidArgument, message.WrongUsernameOrPassword)
	}

	// get role data from redis
	var roleName, value string
	if s, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleId)); err != nil {
		var roleData []model.Role
		res := l.svcCtx.DB.Find(&roleData)
		if res.RowsAffected == 0 {
			logx.Error("Fail to find any role")
			return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
		}
		for _, v := range roleData {
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", v.ID), v.Name)
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
			if err != nil {
				logx.Errorw(logmessage.RedisError, logx.Field("Detail", err.Error()))
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
			logx.Error("Fail to find the role data")
			return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
		}
	}

	logx.Infow("Log in successfully", logx.Field("UUID", u.UUID))
	return &core.LoginResp{
		Id:        u.UUID,
		RoleValue: value,
		RoleName:  roleName,
		RoleId:    u.RoleId,
	}, nil

}
