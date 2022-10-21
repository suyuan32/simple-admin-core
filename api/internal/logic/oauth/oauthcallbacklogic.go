package oauth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewOauthCallbackLogic(r *http.Request, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *OauthCallbackLogic) OauthCallback() (resp *types.CallbackResp, err error) {
	result, err := l.svcCtx.CoreRpc.OauthCallback(context.Background(), &core.CallbackReq{
		State: l.r.FormValue("state"),
		Code:  l.r.FormValue("code"),
	})

	if err != nil {
		return nil, err
	}

	token, err := user.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, result.Id, time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, int64(result.RoleId))

	// add token into database
	expireAt := time.Now().Add(time.Second * 259200).Unix()
	_, err = l.svcCtx.CoreRpc.CreateOrUpdateToken(context.Background(), &core.TokenInfo{
		Id:       0,
		CreateAt: 0,
		UUID:     result.Id,
		Token:    token,
		Source:   strings.Split(l.r.FormValue("state"), "-")[1],
		Status:   1,
		ExpireAt: expireAt,
	})

	if err != nil {
		return nil, err
	}

	return &types.CallbackResp{
		UserId: result.Id,
		Role: types.RoleInfoSimple{
			RoleName: result.RoleName,
			Value:    result.RoleValue,
		},
		Token:  token,
		Expire: uint64(time.Now().Add(time.Second * 259200).Unix()),
	}, nil
}
