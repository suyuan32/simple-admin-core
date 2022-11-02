package logic

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/suyuan32/simple-admin-core/pkg/msg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
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

// role service
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
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("role value had been used", logx.Field("detail", data))
			return nil, status.Error(codes.InvalidArgument, i18n.DuplicateRoleValue)
		}

		err := l.UpdateRoleInfoInRedis()
		if err != nil {
			logx.Errorw("fail to update the role info in Redis", logx.Field("detail", err.Error()))
			return nil, err
		}

		logx.Infow("create role successfully", logx.Field("detail", data))
		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin *model.Role
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.InvalidArgument, errorx.TargetNotExist)
		}
		if check.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
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
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
		err := l.UpdateRoleInfoInRedis()
		if err != nil {
			logx.Errorw("fail to update the role info in Redis", logx.Field("detail", err.Error()))
			return nil, err
		}

		logx.Infow("update role successfully", logx.Field("detail", data))
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
