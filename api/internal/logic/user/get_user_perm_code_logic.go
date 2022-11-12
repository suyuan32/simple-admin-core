package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetUserPermCodeLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetUserPermCodeLogic) GetUserPermCode() (resp *types.PermCodeResp, err error) {
	roleId := l.ctx.Value("roleId")
	if roleId == nil {
		return nil, &errorx.ApiError{
			Code: http.StatusUnauthorized,
			Msg:  "login.requireLogin",
		}
	}
	return &types.PermCodeResp{BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.lang, i18n.Success)},
		Data: []string{fmt.Sprintf("%v", roleId)}}, nil
}
