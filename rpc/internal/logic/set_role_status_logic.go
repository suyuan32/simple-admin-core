package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SetRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleStatusLogic {
	return &SetRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRoleStatusLogic) SetRoleStatus(in *core.SetStatusReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Table("roles").Where("id = ?", in.Id).Update("status", in.Status)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("update role status failed, please check the role id", logx.Field("roleId", in.Id))
		return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
	}

	logx.Infow("update role status successfully", logx.Field("roleId", in.Id),
		logx.Field("status", in.Status))
	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
