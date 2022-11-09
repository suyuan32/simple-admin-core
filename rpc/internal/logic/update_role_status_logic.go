package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleStatusLogic {
	return &UpdateRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleStatusLogic) UpdateRoleStatus(in *core.StatusCodeReq) (*core.BaseResp, error) {
	role, err := l.svcCtx.DB.Role.UpdateOneID(in.Id).SetStatus(uint8(in.Status)).Save(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}
	}

	// update redis

	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", role.ID), role.Name)
	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", role.ID), role.Value)
	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", role.ID), strconv.Itoa(int(role.Status)))
	if err != nil {
		return nil, statuserr.NewInternalError(errorx.RedisError)
	}

	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
