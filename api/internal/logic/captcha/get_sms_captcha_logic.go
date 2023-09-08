package captcha

import (
	"context"
	"github.com/duke-git/lancet/v2/random"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"
	"github.com/zeromicro/go-zero/core/errorx"
	"strconv"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsCaptchaLogic {
	return &GetSmsCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSmsCaptchaLogic) GetSmsCaptcha(req *types.SmsCaptchaReq) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeInvalidArgumentError("captcha.mcmsNotEnabled")
	}

	captcha := random.RandInt(10000, 99999)
	_, err = l.svcCtx.McmsRpc.SendSms(l.ctx, &mcms.SmsInfo{
		PhoneNumber: []string{req.PhoneNumber},
		Params:      []string{strconv.Itoa(captcha)},
		TemplateId:  &l.svcCtx.Config.ProjectConf.SmsTemplateId,
		AppId:       &l.svcCtx.Config.ProjectConf.SmsAppId,
		SignName:    &l.svcCtx.Config.ProjectConf.SmsSignName,
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Redis.Setex("CAPTCHA_"+req.PhoneNumber, strconv.Itoa(captcha), l.svcCtx.Config.ProjectConf.EmailCaptchaExpiredTime)
	if err != nil {
		logx.Errorw("failed to write sms captcha to redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}
