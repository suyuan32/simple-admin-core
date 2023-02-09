package dictionarydetail

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(in *core.DictionaryDetailInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.DictionaryDetail.UpdateOneID(in.Id).
		SetNotEmptyStatus(uint8(in.Status)).
		SetNotEmptyTitle(in.Title).
		SetNotEmptyKey(in.Key).
		SetNotEmptyValue(in.Value).
		SetNotEmptySort(in.Sort).
		SetNotEmptyDictionaryID(in.DictionaryId).
		Exec(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
