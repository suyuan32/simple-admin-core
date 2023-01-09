package user

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetUserProfileLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetUserProfileLogic) GetUserProfile() (resp *types.ProfileResp, err error) {
	result, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.UUIDReq{
		Id: l.ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ProfileResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data = types.ProfileInfo{
		Nickname: result.Nickname,
		Avatar:   result.Avatar,
		Mobile:   result.Mobile,
		Email:    result.Email,
	}

	return resp, nil
}
