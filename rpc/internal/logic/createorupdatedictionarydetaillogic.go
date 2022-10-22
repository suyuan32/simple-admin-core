package logic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/common/message"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

type CreateOrUpdateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryDetailLogic {
	return &CreateOrUpdateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateDictionaryDetailLogic) CreateOrUpdateDictionaryDetail(in *core.DictionaryDetail) (*core.BaseResp, error) {
	var parent model.Dictionary
	check := l.svcCtx.DB.Where("id = ?", in.ParentId).First(&parent)

	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		logx.Errorw(message.ParentNotExist, logx.Field("detail", in))
		return nil, status.Error(codes.InvalidArgument, message.ParentNotExist)
	}

	if check.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", check.Error.Error()))
		return nil, status.Error(codes.Internal, check.Error.Error())
	}
	if in.Id == 0 {
		result := l.svcCtx.DB.Create(&model.DictionaryDetail{
			Model:        gorm.Model{},
			Title:        in.Title,
			Key:          in.Key,
			Value:        in.Value,
			Status:       in.Status,
			DictionaryID: uint(in.ParentId),
		})

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		logx.Infow(logmessage.CreateSuccess, logx.Field("detail", in))

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin model.DictionaryDetail
		checkOrigin := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)

		if errors.Is(checkOrigin.Error, gorm.ErrRecordNotFound) {
			logx.Errorw(logmessage.TargetNotFound, logx.Field("detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.TargetNotExist)
		}

		if checkOrigin.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("detail", checkOrigin.Error.Error()))
			return nil, status.Error(codes.Internal, checkOrigin.Error.Error())
		}

		origin.Title = in.Title
		origin.Key = in.Key
		origin.Value = in.Value
		origin.Status = in.Status

		result := l.svcCtx.DB.Save(&origin)

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			logx.Errorw(logmessage.UpdateFailed, logx.Field("detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}

		logx.Infow(logmessage.UpdateSuccess, logx.Field("detail", in))

		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}
