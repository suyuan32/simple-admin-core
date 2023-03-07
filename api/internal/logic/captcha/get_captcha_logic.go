package captcha

import (
	"context"

	"github.com/mojocn/base64Captcha"
	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Store  *captcha.RedisStore
	driver *base64Captcha.DriverDigit
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
	if driver == nil || Store == nil {
		initStoreAndDriver(l.svcCtx.Config, l.svcCtx.Redis)
	}
	gen := base64Captcha.NewCaptcha(driver, Store)
	if id, b64s, err := gen.Generate(); err != nil {
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

func initStoreAndDriver(c config.Config, r *redis.Redis) {
	driver = base64Captcha.NewDriverDigit(c.Captcha.ImgHeight, c.Captcha.ImgWidth,
		c.Captcha.KeyLong, 0.7, 80)
	Store = captcha.NewRedisStore(r)
}
