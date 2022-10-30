package logic

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/common/msg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/model"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateTokenLogic {
	return &CreateOrUpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Token management
func (l *CreateOrUpdateTokenLogic) CreateOrUpdateToken(in *core.TokenInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		result := l.svcCtx.DB.Create(&model.Token{
			Model:     gorm.Model{},
			UUID:      in.UUID,
			Token:     in.Token,
			Status:    in.Status,
			Source:    in.Source,
			ExpiredAt: time.Unix(in.ExpiredAt, 0),
		})
		if result.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw("Token already exists", logx.Field("detail", in))
			return nil, status.Error(codes.InvalidArgument, msg.DictionaryAlreadyExists)
		}

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin model.Token
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			logx.Errorw(logmsg.TargetNotFound, logx.Field("detail", in))
			return nil, status.Error(codes.InvalidArgument, errorx.TargetNotExist)
		}

		if check.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}

		origin.UUID = in.UUID
		origin.Token = in.Token
		origin.Status = in.Status
		origin.Source = in.Source
		origin.ExpiredAt = time.Unix(in.ExpiredAt, 0)

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
