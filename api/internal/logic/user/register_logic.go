package user

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/captcha"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/msg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.SimpleMsg, err error) {
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.CreateOrUpdateUser(l.ctx,
			&core.CreateOrUpdateUserReq{
				Id:       0,
				Username: req.Username,
				Password: req.Password,
				Email:    req.Email,
			})
		if err != nil {
			l.Logger.Error("register logic: create user err: ", err.Error())
			return nil, err
		}
		resp = &types.SimpleMsg{
			Msg: user.Msg,
		}
		return resp, nil
	} else {
		return nil, errorx.NewApiError(http.StatusBadRequest, i18n.WrongCaptcha)
	}
}
