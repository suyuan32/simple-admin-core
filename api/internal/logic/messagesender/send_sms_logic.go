package messagesender

import (
	"context"
	"strings"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *SendSmsLogic) SendSms(req *types.SendSmsReq) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	var result *mcms.BaseUUIDResp
	if req.TemplateId == nil {
		result, err = l.svcCtx.McmsRpc.SendSms(l.ctx, &mcms.SmsInfo{
			PhoneNumber: []string{req.PhoneNumber},
			Params:      strings.Split(req.Params, ","),
			TemplateId:  &l.svcCtx.Config.ProjectConf.SmsTemplateId,
			AppId:       &l.svcCtx.Config.ProjectConf.SmsAppId,
			SignName:    &l.svcCtx.Config.ProjectConf.SmsSignName,
			Provider:    nil,
		})
	} else {
		result, err = l.svcCtx.McmsRpc.SendSms(l.ctx, &mcms.SmsInfo{
			PhoneNumber: []string{req.PhoneNumber},
			Params:      strings.Split(req.Params, ","),
			TemplateId:  req.TemplateId,
			AppId:       req.AppId,
			SignName:    req.SignName,
			Provider:    req.Provider,
		})
	}

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
