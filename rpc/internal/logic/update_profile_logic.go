package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/model"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProfileLogic) UpdateProfile(in *core.UpdateProfileReq) (*core.BaseResp, error) {
	var origin model.User
	result := l.svcCtx.DB.Where("uuid = ?", in.Uuid).First(&origin)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Fail to find user, please check the UUID", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
	}

	origin.Email = in.Email
	origin.Mobile = in.Mobile
	origin.Nickname = in.Nickname
	origin.Avatar = in.Avatar

	result = l.svcCtx.DB.Save(&origin)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Fail to update the user profile", logx.Field("detail", origin))
		return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
	}

	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
