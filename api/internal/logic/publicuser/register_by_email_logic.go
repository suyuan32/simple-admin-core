package publicuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByEmailLogic {
	return &RegisterByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *RegisterByEmailLogic) RegisterByEmail(req *types.RegisterByEmailReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.RegisterVerify != "email" && l.svcCtx.Config.ProjectConf.RegisterVerify != "sms_or_email" {
		return nil, errorx.NewCodeAbortedError("login.registerTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get(l.ctx, config.RedisCaptchaPrefix+req.Email).Result()
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for email validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}

	if captchaData == req.Captcha {
		_, err := l.svcCtx.CoreRpc.CreateUser(l.ctx,
			&core.UserInfo{
				Username:     &req.Username,
				Password:     &req.Password,
				Email:        &req.Email,
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

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.Email).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		resp = &types.BaseMsgResp{
			Msg: l.svcCtx.Trans.Trans(l.ctx, "login.signupSuccessTitle"),
		}
		return resp, nil
	}

	return nil, errorx.NewInvalidArgumentError(
		l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
}
