package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/common/message"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *core.ChangePasswordReq) (*core.BaseResp, error) {
	var target model.User
	result := l.svcCtx.DB.Where("uuid = ?", in.Uuid).First(&target)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}
	if result.RowsAffected == 0 {
		logx.Errorw("User does not exist", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
	}

	if ok := util.BcryptCheck(in.OldPassword, target.Password); ok {
		target.Password = util.BcryptEncrypt(in.NewPassword)
		result = l.svcCtx.DB.Updates(&target)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, errorx.DatabaseError)
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
	} else {
		logx.Errorw("Old password is wrong", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.InvalidArgument, message.WrongPassword)
	}
	logx.Infow("Change password successful", logx.Field("UUID", in.Uuid))
	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
