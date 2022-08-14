package user

import (
	"context"
	"net/http"
	"time"

	"github.com/suyuan32/simple-admin-core/api/common/errorx"
	"github.com/suyuan32/simple-admin-core/api/internal/logic/captcha"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.Login(context.Background(),
			&core.LoginReq{
				Username: req.Username,
				Password: req.Password,
			})
		if err != nil {
			return nil, err
		}

		token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, user.Id, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, int64(user.RoleId))
		if err != nil {
			return nil, err
		}
		resp = &types.LoginResp{
			UserId: user.Id,
			Token:  token,
			Expire: uint64(time.Now().Add(time.Second * 259200).Unix()),
			Role: types.RoleInfoSimple{
				Value:    user.RoleValue,
				RoleName: user.RoleName,
			},
		}
		return resp, nil
	} else {
		return nil, errorx.NewApiError(http.StatusBadRequest, errorx.WrongCaptcha)
	}
}

func (l *LoginLogic) getJwtToken(secretKey, uuid string, iat, seconds, roleId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = uuid
	claims["roleId"] = roleId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
