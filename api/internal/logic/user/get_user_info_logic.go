package user

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetUserInfoLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	if l.ctx.Value("userId").(string) == "" {
		return nil, errorx.NewApiError(http.StatusUnauthorized, "Please log in")
	}
	user, err := l.svcCtx.CoreRpc.GetUserById(l.ctx,
		&core.UUIDReq{Uuid: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}

	tenants := make([]types.Tanent, 0)
	for _, v := range user.Tenants {
		tenants = append(tenants, types.Tanent{
			Tenant_Id:  v.TenantId,
			TenantName: v.TenantName,
		})
	}

	return &types.GetUserInfoResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.lang, i18n.Success)},
		Data: types.UserBaseInfo{
			Tenants:  tenants,
			UUID:     user.Uuid,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Roles: types.GetUserRoleInfo{
				RoleName: user.RoleName,
				Value:    user.RoleValue,
			},
		},
	}, nil
}
