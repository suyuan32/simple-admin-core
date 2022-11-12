package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
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
	exist, err := l.svcCtx.DB.Menu.Query().Where(menu.ParentID(in.Id)).Exist(l.ctx)
	if err != nil {
		logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
		return nil, statuserr.NewInternalError(errorx.DatabaseError)
	}

	if exist {
		logx.Errorw("delete menu failed, please check its children had been deleted",
			logx.Field("menuId", in.Id))
		return nil, statuserr.NewInvalidArgumentError("menu.deleteChildrenDesc")
	}

	err = utils.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		err = l.svcCtx.DB.Menu.Update().ClearParams().Exec(l.ctx)

		if err != nil {
			logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
			return statuserr.NewInternalError(errorx.DatabaseError)
		}

		err = l.svcCtx.DB.Menu.DeleteOneID(in.Id).Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return statuserr.NewInvalidArgumentError(errorx.TargetNotFound)
			default:
				logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
				return statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		return nil
	})

	if err != nil {
		logx.Errorf("delete dictionary failed, error : %s", err.Error())
		return nil, err
	}

	return &core.BaseResp{Msg: errorx.DeleteSuccess}, nil
}
