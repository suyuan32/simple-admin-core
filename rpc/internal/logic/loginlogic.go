package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"

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

	roleName, value, err := getRoleInfo(u.RoleId, l.svcCtx.Redis, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}

	logx.Infow("Log in successfully", logx.Field("UUID", u.UUID))
	return &core.LoginResp{
		Id:        u.UUID,
		RoleValue: value,
		RoleName:  roleName,
		RoleId:    u.RoleId,
	}, nil

}

func getRoleInfo(roleId uint32, rds *redis.Redis, db *gorm.DB) (roleName, roleValue string, err error) {
	if s, err := rds.Hget("roleData", strconv.Itoa(int(roleId))); err != nil {
		var roleData []model.Role
		res := db.Find(&roleData)
		if res.RowsAffected == 0 {
			logx.Error("Fail to find any role")
			return "", "", status.Error(codes.NotFound, errorx.TargetNotExist)
		}
		for _, v := range roleData {
			err = rds.Hset("roleData", strconv.Itoa(int(v.ID)), v.Name)
			err = rds.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
			err = rds.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
			if err != nil {
				logx.Errorw(logmessage.RedisError, logx.Field("Detail", err.Error()))
				return "", "", status.Error(codes.Internal, errorx.RedisError)
			}
			if v.ID == uint(roleId) {
				roleName = v.Name
				roleValue = v.Value
			}
		}
	} else {
		roleName = s
		roleValue, err = rds.Hget("roleData", fmt.Sprintf("%d_value", roleId))
		if err != nil {
			logx.Error("Fail to find the role data")
			return "", "", status.Error(codes.NotFound, errorx.TargetNotExist)
		}
	}
	return roleName, roleValue, nil
}
