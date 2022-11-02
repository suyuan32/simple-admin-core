package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(in *core.IDReq) (*core.BaseResp, error) {
	tx := l.svcCtx.DB.Begin()
	childResult := tx.Exec(fmt.Sprintf("delete from dictionary_details where dictionary_id = %d", in.ID))
	if childResult.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", childResult.Error.Error()))
		tx.Rollback()
		return nil, status.Error(codes.Internal, childResult.Error.Error())
	}

	result := tx.Delete(&model.Dictionary{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		tx.Rollback()
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("delete dictionary failed, check the id", logx.Field("DictionaryId", in.ID))
		return nil, status.Error(codes.InvalidArgument, errorx.DeleteFailed)
	}

	tx.Commit()

	logx.Infow("delete dictionary successfully", logx.Field("DictionaryId", in.ID))

	return &core.BaseResp{Msg: errorx.DeleteSuccess}, nil
}
