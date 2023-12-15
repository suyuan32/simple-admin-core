package captcha

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaResp, err error) {
	if id, b64s, _, err := l.svcCtx.Captcha.Generate(); err != nil {
		logx.Errorw("fail to generate captcha", logx.Field("detail", err.Error()))
		return &types.CaptchaResp{
			BaseDataInfo: types.BaseDataInfo{Code: errorcode.Internal, Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Failed)},
			Data:         types.CaptchaInfo{},
		}, nil
	} else {
		resp = &types.CaptchaResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
			Data: types.CaptchaInfo{
				CaptchaId: id,
				ImgPath:   b64s,
			},
		}
		return resp, nil
	}
}
