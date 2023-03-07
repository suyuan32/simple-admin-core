package menuparam

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type UpdateMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuParamLogic {
	return &UpdateMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuParamLogic) UpdateMenuParam(in *core.MenuParamInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.MenuParam.UpdateOneID(in.Id).
		SetNotEmptyType(in.Type).
		SetNotEmptyKey(in.Key).
		SetNotEmptyValue(in.Value).
		SetNotEmptyMenuID(in.MenuId).
		Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
