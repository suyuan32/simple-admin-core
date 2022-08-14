package logic

import (
	"context"
	"github.com/suyuan32/simple-admin-core/api/common/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

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
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
	}

	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
