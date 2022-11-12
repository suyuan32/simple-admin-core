package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateOrUpdateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryLogic {
	return &CreateOrUpdateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// dictionary management service
func (l *CreateOrUpdateDictionaryLogic) CreateOrUpdateDictionary(in *core.DictionaryInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		err := l.svcCtx.DB.Dictionary.Create().
			SetTitle(in.Title).
			SetName(in.Name).
			SetStatus(uint8(in.Status)).
			SetDesc(in.Desc).
			Exec(l.ctx)

		if err != nil {
			if ent.IsConstraintError(err) {
				logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInvalidArgumentError(errorx.CreateFailed)
			}
			logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		exist, err := l.svcCtx.DB.Dictionary.Query().Where(dictionary.IDEQ(in.Id)).Exist(l.ctx)
		if err != nil {
			return nil, err
		}

		if err != nil {
			logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
			return nil, status.Error(codes.Internal, err.Error())
		}

		if !exist {
			logx.Errorw(logmsg.TARGET_NOT_FOUND, logx.Field("id", in.Id))
			return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
		}

		err = l.svcCtx.DB.Dictionary.UpdateOneID(in.Id).
			SetTitle(in.Title).
			SetName(in.Name).
			SetStatus(uint8(in.Status)).
			SetDesc(in.Desc).
			Exec(l.ctx)

		if err != nil {
			logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInvalidArgumentError(logmsg.UPDATE_FAILED)
		}

		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}
