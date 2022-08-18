package logic

import (
	"context"
	"fmt"
	"github.com/suyuan32/simple-admin-core/common/message"
	"github.com/zeromicro/go-zero/core/errorx"
	"strconv"
	"time"

	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CreateOrUpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateRoleLogic {
	return &CreateOrUpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  role service
func (l *CreateOrUpdateRoleLogic) CreateOrUpdateRole(in *core.RoleInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		data := &model.Role{
			Model:         gorm.Model{},
			Name:          in.Name,
			Value:         in.Value,
			DefaultRouter: in.DefaultRouter,
			Status:        in.Status,
			Remark:        in.Remark,
			OrderNo:       in.OrderNo,
		}
		result := l.svcCtx.DB.Create(&data)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, message.DuplicateRoleValue)
		}
		err := l.UpdateRoleInfoInRedis()
		if err != nil {
			return nil, err
		}
		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin *model.Role
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if check.Error != nil {
			return nil, status.Error(codes.Internal, check.Error.Error())
		}
		if check.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
		data := &model.Role{
			Model:         gorm.Model{ID: origin.ID, CreatedAt: origin.CreatedAt, UpdatedAt: time.Now()},
			Name:          in.Name,
			Value:         in.Value,
			DefaultRouter: in.DefaultRouter,
			Status:        in.Status,
			Remark:        in.Remark,
			OrderNo:       in.OrderNo,
		}
		result := l.svcCtx.DB.Save(&data)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
		err := l.UpdateRoleInfoInRedis()
		if err != nil {
			return nil, err
		}
		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}

func (l *CreateOrUpdateRoleLogic) UpdateRoleInfoInRedis() error {
	var roleData []model.Role
	res := l.svcCtx.DB.Find(&roleData)
	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, errorx.TargetNotExist)
	}
	for _, v := range roleData {
		err := l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", v.ID), v.Name)
		err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
		err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
		if err != nil {
			return status.Error(codes.Internal, errorx.RedisError)
		}
	}
	return nil
}
