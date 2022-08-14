package logic

import (
	"context"
	"github.com/suyuan32/simple-admin-core/api/common/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *core.IDReq) (*core.BaseResp, error) {
	var users []model.User
	check := l.svcCtx.DB.Model(&model.User{}).Where("role_id = ?", in.ID).Find(&users).RowsAffected
	if check != 0 {
		return nil, status.Error(codes.InvalidArgument, errorx.UserExists)
	}
	result := l.svcCtx.DB.Delete(&model.Role{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, errorx.DeleteFailed)
	}

	return &core.BaseResp{Msg: errorx.DeleteSuccess}, nil
}
