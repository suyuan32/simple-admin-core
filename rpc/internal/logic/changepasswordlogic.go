package logic

import (
	"context"
	"github.com/suyuan32/simple-admin-core/common/message"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
	}

	if ok := util.BcryptCheck(in.OldPassword, target.Password); ok {
		target.Password = util.BcryptEncrypt(in.NewPassword)
		result = l.svcCtx.DB.Updates(&target)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, errorx.DatabaseError)
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, message.WrongUsernameOrPassword)
	}
	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
