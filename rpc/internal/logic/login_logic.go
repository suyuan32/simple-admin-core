package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
	result, err := l.svcCtx.DB.User.Query().Where(user.UsernameEQ(in.Username)).First(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("user not found", logx.Field("username", in.Username))
			return nil, status.Error(codes.InvalidArgument, "login.userNotExist")
		}
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	if ok := utils.BcryptCheck(in.Password, result.Password); !ok {
		logx.Errorw("wrong password", logx.Field("detail", in))
		return nil, status.Error(codes.InvalidArgument, "login.wrongUsernameOrPassword")
	}

	roleName, value, err := getRoleInfo(result.RoleID, l.svcCtx.Redis, l.svcCtx.DB, l.ctx)
	if err != nil {
		return nil, err
	}

	logx.Infow("log in successfully", logx.Field("UUID", result.UUID))
	return &core.LoginResp{
		Id:        result.UUID,
		RoleValue: value,
		RoleName:  roleName,
		RoleId:    result.RoleID,
	}, nil
}

func getRoleInfo(roleId uint64, rds *redis.Redis, db *ent.Client, ctx context.Context) (roleName, roleValue string, err error) {
	if s, err := rds.Hget("roleData", strconv.Itoa(int(roleId))); err != nil {
		roleData, err := db.Role.Query().All(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				logx.Error("fail to find any roles")
				return "", "", status.Error(codes.NotFound, errorx.TargetNotFound)
			}
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return "", "", status.Error(codes.NotFound, errorx.TargetNotFound)
		}

		for _, v := range roleData {
			err = rds.Hset("roleData", strconv.Itoa(int(v.ID)), v.Name)
			err = rds.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
			err = rds.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return "", "", statuserr.NewInternalError(errorx.RedisError)
			}
			if v.ID == roleId {
				roleName = v.Name
				roleValue = v.Value
			}
		}
	} else {
		roleName = s
		roleValue, err = rds.Hget("roleData", fmt.Sprintf("%d_value", roleId))
		if err != nil {
			logx.Error("fail to find the role data")
			return "", "", status.Error(codes.NotFound, errorx.TargetNotFound)
		}
	}
	return roleName, roleValue, nil
}
