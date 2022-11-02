package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProviderLogic {
	return &DeleteProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProviderLogic) DeleteProvider(in *core.IDReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Delete(&model.OauthProvider{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Delete provider failed, check the id", logx.Field("ProviderId", in.ID))
		return nil, status.Error(codes.InvalidArgument, errorx.DeleteFailed)
	}

	logx.Infow("Delete provider successfully", logx.Field("ProviderId", in.ID))

	return &core.BaseResp{Msg: errorx.DeleteSuccess}, nil
}
