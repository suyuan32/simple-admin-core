package captcha

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chimerakang/simple-admin-common/config"
	"github.com/chimerakang/simple-admin-common/i18n"
	"github.com/chimerakang/simple-admin-message-center/types/mcms"
	"github.com/duke-git/lancet/v2/random"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/chimerakang/simple-admin-core/api/internal/svc"
	"github.com/chimerakang/simple-admin-core/api/internal/types"

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

	var params []string
	if l.svcCtx.Config.ProjectConf.SmsParamsType == "json" {
		params = []string{fmt.Sprintf("{\"code\":%d}", captcha)}
	} else {
		params = []string{strconv.Itoa(captcha)}
	}

	_, err = l.svcCtx.McmsRpc.SendSms(l.ctx, &mcms.SmsInfo{
		PhoneNumber: []string{req.PhoneNumber},
		Params:      params,
		TemplateId:  &l.svcCtx.Config.ProjectConf.SmsTemplateId,
		AppId:       &l.svcCtx.Config.ProjectConf.SmsAppId,
		SignName:    &l.svcCtx.Config.ProjectConf.SmsSignName,
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Redis.Set(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber, strconv.Itoa(captcha), time.Duration(l.svcCtx.Config.ProjectConf.EmailCaptchaExpiredTime)*time.Second).Err()
	if err != nil {
		logx.Errorw("failed to write sms captcha to redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}
