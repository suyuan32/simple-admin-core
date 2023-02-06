package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryDetailLogic {
	return &CreateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryDetailLogic) CreateDictionaryDetail(in *core.DictionaryDetail) (*core.BaseResp, error) {
	exist, err := l.svcCtx.DB.Dictionary.Query().Where(dictionary.IDEQ(in.DictionaryId)).Exist(l.ctx)
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	if !exist {
		logx.Errorw("the parent menu does not exist", logx.Field("detail", in))
		return nil, statuserr.NewInvalidArgumentError("menu.parentNotExist")
	}

	err = l.svcCtx.DB.DictionaryDetail.Create().
		SetTitle(in.Title).
		SetKey(in.Key).
		SetValue(in.Value).
		SetStatus(uint8(in.Status)).
		SetDictionaryID(in.DictionaryId).
		Exec(l.ctx)

	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.CreateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &core.BaseResp{Msg: i18n.CreateSuccess}, nil

	return &core.BaseResp{}, nil
}
