package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/msg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *core.IDReq) (*core.BaseResp, error) {
	var children []*model.Menu
	check := l.svcCtx.DB.Where("parent_id = ?", in.ID).Find(&children)
	if check.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", check.Error.Error()))
		return nil, status.Error(codes.Internal, check.Error.Error())
	}
	if check.RowsAffected != 0 {
		logx.Errorw("delete menu failed, please check its children had been deleted",
			logx.Field("menuId", in.ID))
		return nil, status.Error(codes.InvalidArgument, i18n.ChildrenExistError)
	}

	result := l.svcCtx.DB.Delete(&model.Menu{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("delete menu failed, please check the menu id", logx.Field("menuId", in.ID))
		return nil, status.Error(codes.InvalidArgument, i18n.MenuNotExists)
	}

	logx.Infow("delete menu successfully", logx.Field("menuId", in.ID))
	return &core.BaseResp{Msg: errorx.DeleteSuccess}, nil
}
