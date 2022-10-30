package logic

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/common/msg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
	"github.com/suyuan32/simple-admin-core/rpc/model"
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
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logx.Errorw("user does not exist", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
	}

	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	if ok := util.BcryptCheck(in.OldPassword, target.Password); ok {
		target.Password = util.BcryptEncrypt(in.NewPassword)
		result = l.svcCtx.DB.Updates(&target)
		if result.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, errorx.DatabaseError)
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
	} else {
		logx.Errorw("old password is wrong", logx.Field("UUID", in.Uuid))
		return nil, status.Error(codes.InvalidArgument, msg.WrongPassword)
	}
	logx.Infow("change password successful", logx.Field("UUID", in.Uuid))
	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
