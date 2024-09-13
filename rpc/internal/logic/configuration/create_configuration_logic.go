package configuration

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateConfigurationLogic {
	return &CreateConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateConfigurationLogic) CreateConfiguration(in *core.ConfigurationInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Configuration.Create().
		SetNotNilSort(in.Sort).
		SetNotNilState(in.State).
		SetNotNilName(in.Name).
		SetNotNilKey(in.Key).
		SetNotNilValue(in.Value).
		SetNotNilCategory(in.Category).
		SetNotNilRemark(in.Remark).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
