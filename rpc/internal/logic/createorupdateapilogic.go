package logic

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/common/message"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CreateOrUpdateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiLogic {
	return &CreateOrUpdateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// api management service
func (l *CreateOrUpdateApiLogic) CreateOrUpdateApi(in *core.ApiInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		data := &model.Api{
			Model:       gorm.Model{},
			Path:        in.Path,
			Description: in.Description,
			ApiGroup:    in.Group,
			Method:      in.Method,
		}
		result := l.svcCtx.DB.Create(&data)

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			logx.Errorw(message.ApiAlreadyExists, logx.Field("Detail", data))
			return nil, status.Error(codes.InvalidArgument, message.ApiAlreadyExists)
		}

		logx.Infow(errorx.CreateSuccess, logx.Field("path", in.Path), logx.Field("desc", in.Description),
			logx.Field("group", in.Group), logx.Field("method", in.Method))

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin *model.Api
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if check.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}
		if check.RowsAffected == 0 {
			logx.Errorw(errorx.TargetNotExist, logx.Field("id", in.Id))
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}

		data := &model.Api{
			Model:       gorm.Model{ID: origin.ID, CreatedAt: origin.CreatedAt, UpdatedAt: time.Now()},
			Path:        in.Path,
			Description: in.Description,
			ApiGroup:    in.Group,
			Method:      in.Method,
		}
		result := l.svcCtx.DB.Save(&data)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw(errorx.UpdateFailed)
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}

		logx.Infow(errorx.UpdateSuccess, logx.Field("Detail", data))
		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}
