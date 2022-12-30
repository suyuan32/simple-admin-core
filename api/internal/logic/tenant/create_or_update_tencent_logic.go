package tenant

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTencentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateTencentLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateTencentLogic {
	return &CreateOrUpdateTencentLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateTencentLogic) CreateOrUpdateTencent(req *types.CreateOrUpdateTenantReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
