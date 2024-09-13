package configuration

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigurationLogic {
	return &UpdateConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigurationLogic) UpdateConfiguration(in *core.ConfigurationInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Configuration.UpdateOneID(*in.Id).
		SetNotNilSort(in.Sort).
		SetNotNilState(in.State).
		SetNotNilName(in.Name).
		SetNotNilKey(in.Key).
		SetNotNilValue(in.Value).
		SetNotNilCategory(in.Category).
		SetNotNilRemark(in.Remark).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
