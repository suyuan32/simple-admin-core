package publicuser

import (
	"context"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *LoginByEmailLogic) LoginByEmail(req *types.LoginByEmailReq) (resp *types.LoginResp, err error) {
	if l.svcCtx.Config.ProjectConf.LoginVerify != "email" && l.svcCtx.Config.ProjectConf.LoginVerify != "sms_or_email" &&
		l.svcCtx.Config.ProjectConf.LoginVerify != "all" {
		return nil, errorx.NewCodeAbortedError("login.loginTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get(l.ctx, config.RedisCaptchaPrefix+req.Email).Result()
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for email validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}

	if captchaData == req.Captcha {
		userData, err := l.svcCtx.CoreRpc.GetUserList(l.ctx, &core.UserListReq{
			Page:     1,
			PageSize: 1,
			Email:    &req.Email,
		})
		if err != nil {
			return nil, err
		}

		if userData.Total == 0 {
			return nil, errorx.NewCodeInvalidArgumentError("login.userNotExist")
		}

		if *userData.Data[0].Status != uint32(common.StatusNormal) {
			return nil, errorx.NewCodeInvalidArgumentError("login.userBanned")
		}

		token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", userData.Data[0].Id), jwt.WithOption("roleId",
				strings.Join(userData.Data[0].RoleCodes, ",")), jwt.WithOption("deptId", userData.Data[0].DepartmentId))
		if err != nil {
			return nil, err
		}

		// add token into database
		expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
		_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
			Uuid:      userData.Data[0].Id,
			Token:     pointy.GetPointer(token),
			Source:    pointy.GetPointer("core_user"),
			Status:    pointy.GetPointer(uint32(common.StatusNormal)),
			Username:  userData.Data[0].Username,
			ExpiredAt: pointy.GetPointer(expiredAt),
		})

		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.Email).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		resp = &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
			Data: types.LoginInfo{
				UserId: *userData.Data[0].Id,
				Token:  token,
				Expire: uint64(expiredAt),
			},
		}
		return resp, nil
	}

	return nil, errorx.NewCodeInvalidArgumentError("login.wrongCaptcha")
}
