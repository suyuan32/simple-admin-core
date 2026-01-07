package publicuser

import (
	"context"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

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
	if l.svcCtx.Config.ProjectConf.LoginVerify != "captcha" && l.svcCtx.Config.ProjectConf.LoginVerify != "all" {
		return nil, errorx.NewCodeAbortedError("login.loginTypeForbidden")
	}

	if ok := l.svcCtx.Captcha.Verify(config.RedisCaptchaPrefix+req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.GetUserByUsername(l.ctx,
			&core.UsernameReq{
				Username: req.Username,
			})
		if err != nil {
			if e, ok := status.FromError(err); ok {
				if e.Message() == i18n.TargetNotFound {
					return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
				}
			}

			return nil, err
		}

		if user.Status != nil && *user.Status != uint32(common.StatusNormal) {
			return nil, errorx.NewCodeInvalidArgumentError("login.userBanned")
		}

		if !encrypt.BcryptCheck(req.Password, *user.Password) {
			return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
		}

		token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", user.Id), jwt.WithOption("roleId",
				strings.Join(user.RoleCodes, ",")), jwt.WithOption("deptId", user.DepartmentId))
		if err != nil {
			return nil, err
		}

		// add token into database
		expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
		_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
			Uuid:      user.Id,
			Token:     pointy.GetPointer(token),
			Source:    pointy.GetPointer("core_user"),
			Status:    pointy.GetPointer(uint32(common.StatusNormal)),
			Username:  user.Username,
			ExpiredAt: pointy.GetPointer(expiredAt),
		})

		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.CaptchaId).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		resp = &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
			Data: types.LoginInfo{
				UserId: *user.Id,
				Token:  token,
				Expire: uint64(expiredAt),
			},
		}
		return resp, nil
	}

	return nil, errorx.NewCodeInvalidArgumentError("login.wrongCaptcha")
}
