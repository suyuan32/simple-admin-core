package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/common/message"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CreateOrUpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateUserLogic {
	return &CreateOrUpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateUserLogic) CreateOrUpdateUser(in *core.CreateOrUpdateUserReq) (*core.BaseResp, error) {
	if in.Id == 0 {
		var u model.User
		check := l.svcCtx.DB.Where("username = ? OR email = ?", in.Username, in.Email).First(&u)

		if check.RowsAffected != 0 {
			logx.Errorw("Username or email address had been used", logx.Field("Username", in.Username),
				logx.Field("Email", in.Email))
			return nil, status.Error(codes.InvalidArgument, message.UserAlreadyExists)
		}

		data := &model.User{
			UUID:     uuid.NewString(),
			Username: in.Username,
			Nickname: in.Username,
			Password: util.BcryptEncrypt(in.Password),
			Email:    in.Email,
			RoleId:   in.RoleId,
			Avatar:   in.Avatar,
			Mobile:   in.Mobile,
			Status:   in.Status,
		}

		result := l.svcCtx.DB.Create(&data)

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		logx.Infow("Create user successfully", logx.Field("Detail", data))
		return &core.BaseResp{
			Msg: errorx.Success,
		}, nil
	} else {
		var origin model.User
		result := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("User does not find", logx.Field("UserId", in.Id))
			return nil, status.Error(codes.InvalidArgument, message.UserNotExists)
		}

		data := &model.User{
			Model:    gorm.Model{ID: origin.ID, CreatedAt: origin.CreatedAt},
			UUID:     origin.UUID,
			Username: in.Username,
			Nickname: in.Username,
			Password: util.BcryptEncrypt(in.Password),
			Email:    in.Email,
			RoleId:   in.RoleId,
			Avatar:   in.Avatar,
			Mobile:   in.Mobile,
			Status:   in.Status,
		}

		result = l.svcCtx.DB.Save(&data)

		if result.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		logx.Infow("Update user successfully", logx.Field("Detail", data))
		return &core.BaseResp{
			Msg: errorx.Success,
		}, nil
	}
}
