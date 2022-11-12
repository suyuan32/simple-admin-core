package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuParamLogic {
	return &CreateOrUpdateMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuParamLogic) CreateOrUpdateMenuParam(in *core.CreateOrUpdateMenuParamReq) (*core.BaseResp, error) {
	if in.Id == 0 {
		err := l.svcCtx.DB.MenuParam.Create().
			SetType(in.Type).
			SetKey(in.Key).
			SetValue(in.Value).
			SetMenusID(in.MenuId).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.CreateFailed)
			default:
				logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		exist, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(in.MenuId)).Exist(l.ctx)
		if err != nil {
			logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}

		if !exist {
			logx.Errorw("menu not found", logx.Field("menuId", in.Id))
			return nil, statuserr.NewInvalidArgumentError("menu.menuNotExists")
		}

		err = l.svcCtx.DB.MenuParam.UpdateOneID(in.Id).
			SetType(in.Type).
			SetKey(in.Key).
			SetValue(in.Value).
			SetMenusID(in.MenuId).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotFound)
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.UpdateFailed)
			default:
				logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}

}
