package api

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/ent/api"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type CreateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateApiLogic) CreateApi(in *core.ApiInfo) (*core.BaseIDResp, error) {
	// if exist , return success
	if in.Path != nil && in.Method != nil {
		check, err := l.svcCtx.DB.API.Query().Where(api.Path(*in.Path), api.Method(*in.Method)).Only(l.ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		if check != nil {
			return &core.BaseIDResp{Id: check.ID, Msg: i18n.CreateSuccess}, nil
		}
	}

	result, err := l.svcCtx.DB.API.Create().
		SetNotNilPath(in.Path).
		SetNotNilDescription(in.Description).
		SetNotNilAPIGroup(in.ApiGroup).
		SetNotNilMethod(in.Method).
		SetNotNilIsRequired(in.IsRequired).
		SetNotNilServiceName(in.ServiceName).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
