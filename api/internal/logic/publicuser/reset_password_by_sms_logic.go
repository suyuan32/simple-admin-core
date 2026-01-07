package publicuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordBySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordBySmsLogic {
	return &ResetPasswordBySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ResetPasswordBySmsLogic) ResetPasswordBySms(req *types.ResetPasswordBySmsReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.ResetVerify != "sms" && l.svcCtx.Config.ProjectConf.ResetVerify != "sms_or_email" {
		return nil, errorx.NewCodeAbortedError("login.resetTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Result()
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for sms validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}

	if captchaData == req.Captcha {
		userData, err := l.svcCtx.CoreRpc.GetUserList(l.ctx, &core.UserListReq{
			Page:     1,
			PageSize: 1,
			Mobile:   &req.PhoneNumber,
		})
		if err != nil {
			return nil, err
		}

		if userData.Total == 0 {
			return nil, errorx.NewCodeInvalidArgumentError("login.userNotExist")
		}

		result, err := l.svcCtx.CoreRpc.UpdateUser(l.ctx, &core.UserInfo{Id: userData.Data[0].Id, Password: &req.Password})
		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
	}

	return nil, errorx.NewInvalidArgumentError(
		l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
}
