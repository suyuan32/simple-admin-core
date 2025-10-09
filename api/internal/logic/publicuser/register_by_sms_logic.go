package publicuser

import (
	"context"

	"github.com/chimerakang/simple-admin-common/config"
	"github.com/chimerakang/simple-admin-common/enum/errorcode"
	"github.com/chimerakang/simple-admin-common/i18n"
	"github.com/chimerakang/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/chimerakang/simple-admin-core/api/internal/svc"
	"github.com/chimerakang/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterBySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterBySmsLogic {
	return &RegisterBySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *RegisterBySmsLogic) RegisterBySms(req *types.RegisterBySmsReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.RegisterVerify != "sms" && l.svcCtx.Config.ProjectConf.RegisterVerify != "sms_or_email" {
		return nil, errorx.NewCodeAbortedError("login.registerTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Result()
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for sms validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}

	if captchaData == req.Captcha {
		_, err := l.svcCtx.CoreRpc.CreateUser(l.ctx,
			&core.UserInfo{
				Username:     &req.Username,
				Password:     &req.Password,
				Mobile:       &req.PhoneNumber,
				Nickname:     &req.Username,
				Status:       pointy.GetPointer(uint32(1)),
				HomePath:     &l.svcCtx.Config.ProjectConf.DefaultRegisterHomePath,
				RoleIds:      []uint64{l.svcCtx.Config.ProjectConf.DefaultRoleId},
				DepartmentId: pointy.GetPointer(l.svcCtx.Config.ProjectConf.DefaultDepartmentId),
				PositionIds:  []uint64{l.svcCtx.Config.ProjectConf.DefaultPositionId},
			})
		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		resp = &types.BaseMsgResp{
			Msg: l.svcCtx.Trans.Trans(l.ctx, "login.signupSuccessTitle"),
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
	}
}
