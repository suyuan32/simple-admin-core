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

type CreateOrUpdateProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateProviderLogic {
	return &CreateOrUpdateProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// oauth management
func (l *CreateOrUpdateProviderLogic) CreateOrUpdateProvider(in *core.ProviderInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		data := &model.OauthProvider{
			Model:        gorm.Model{},
			Name:         in.Name,
			ClientID:     in.ClientId,
			ClientSecret: in.ClientSecret,
			RedirectURL:  in.RedirectUrl,
			Scopes:       in.Scopes,
			AuthURL:      in.AuthUrl,
			TokenURL:     in.TokenUrl,
			AuthStyle:    int(in.AuthStyle),
			InfoURL:      in.InfoUrl,
		}
		result := l.svcCtx.DB.Create(&data)

		if result.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			logx.Errorw(msg.ApiAlreadyExists, logx.Field("detail", data))
			return nil, status.Error(codes.InvalidArgument, msg.ApiAlreadyExists)
		}

		logx.Infow(logmsg.CreateSuccess, logx.Field("detail", in))

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		var origin *model.OauthProvider
		check := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			logx.Errorw(errorx.TargetNotExist, logx.Field("id", in.Id))
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}
		if check.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}

		data := &model.OauthProvider{
			Model:        gorm.Model{ID: origin.ID, CreatedAt: origin.CreatedAt, UpdatedAt: time.Now()},
			Name:         in.Name,
			ClientID:     in.ClientId,
			ClientSecret: in.ClientSecret,
			RedirectURL:  in.RedirectUrl,
			Scopes:       in.Scopes,
			AuthURL:      in.AuthUrl,
			TokenURL:     in.TokenUrl,
			AuthStyle:    int(in.AuthStyle),
			InfoURL:      in.InfoUrl,
		}
		result := l.svcCtx.DB.Save(&data)
		if result.Error != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			logx.Errorw(errorx.UpdateFailed)
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}

		delete(providerConfig, in.Name)

		logx.Infow(logmsg.UpdateSuccess, logx.Field("detail", data))
		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}
