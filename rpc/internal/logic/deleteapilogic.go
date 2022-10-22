package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteApiLogic) DeleteApi(in *core.IDReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Delete(&model.Api{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("delete API failed, check the id", logx.Field("apiId", in.ID))
		return nil, status.Error(codes.InvalidArgument, errorx.DeleteFailed)
	}

	logx.Infow("delete API successfully", logx.Field("apiId", in.ID))

	return &core.BaseResp{Msg: errorx.DeleteSuccess}, nil
}
