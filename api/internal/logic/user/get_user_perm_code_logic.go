package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermCodeLogic) GetUserPermCode() (resp *types.PermCodeResp, err error) {
	roleId := l.ctx.Value("roleId").(string)
	if roleId == "" {
		return nil, &errorx.ApiError{
			Code: http.StatusUnauthorized,
			Msg:  "login.requireLogin",
		}
	}

	return &types.PermCodeResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data:         strings.Split(roleId, ","),
	}, nil
}

func setRoleInfoToRedis(roleId uint64, rds *redis.Redis, roleInfos []*core.RoleInfo) (err error) {
	if _, err := rds.Hget("roleData", strconv.Itoa(int(roleId))); err != nil {
		for _, v := range roleInfos {
			err = rds.Hset("roleData", strconv.Itoa(int(v.Id)), v.Name)
			err = rds.Hset("roleData", fmt.Sprintf("%d_code", v.Id), v.Code)
			err = rds.Hset("roleData", fmt.Sprintf("%d_status", v.Id), strconv.Itoa(int(v.Status)))
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return errorx.NewCodeInternalError(i18n.RedisError)
			}
		}
	}
	return nil
}
