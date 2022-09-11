package logic

import (
	"context"
	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuParamLogic {
	return &CreateOrUpdateMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuParamLogic) CreateOrUpdateMenuParam(in *core.CreateOrUpdateMenuParamReq) (*core.BaseResp, error) {
	if in.Id == 0 {
		result := l.svcCtx.DB.Create(&model.MenuParam{
			Model:  gorm.Model{},
			MenuId: uint(in.MenuId),
			Type:   in.Type,
			Key:    in.Key,
			Value:  in.Value,
		})

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("Create menu parameter error", logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.TargetNotExist)
		}
		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin model.MenuParam
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if check.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}
		if check.RowsAffected == 0 {
			logx.Errorw("Update menu parameter error, menu parameter does not find", logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.TargetNotExist)
		}
		origin.MenuId = uint(in.MenuId)
		origin.Type = in.Type
		origin.Value = in.Value
		origin.Key = in.Key
		result := l.svcCtx.DB.Save(&origin)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, errorx.DatabaseError)
		}
		if result.RowsAffected == 0 {
			logx.Errorw("Create menu parameter error", logx.Field("Detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}

}
