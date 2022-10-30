package logic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/common/msg"
	"github.com/suyuan32/simple-admin-core/rpc/model"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryLogic {
	return &CreateOrUpdateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// dictionary management service
func (l *CreateOrUpdateDictionaryLogic) CreateOrUpdateDictionary(in *core.DictionaryInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		result := l.svcCtx.DB.Create(&model.Dictionary{
			Model:  gorm.Model{},
			Title:  in.Title,
			Name:   in.Name,
			Status: in.Status,
			Desc:   in.Desc,
			Detail: nil,
		})
		if result.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("dictionary already exists", logx.Field("detail", in))
			return nil, status.Error(codes.InvalidArgument, msg.DictionaryAlreadyExists)
		}

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin model.Dictionary
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if check.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}

		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			logx.Errorw(logmsg.TargetNotFound, logx.Field("detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.TargetNotExist)
		}

		origin.Title = in.Title
		origin.Name = in.Name
		origin.Status = in.Status
		origin.Desc = in.Desc

		result := l.svcCtx.DB.Save(&origin)

		if result.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			logx.Errorw(logmsg.UpdateFailed, logx.Field("detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}

		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}
