package user

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/captcha"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/enum"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewRegisterLogic(r *http.Request, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseMsgResp, err error) {
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.CreateOrUpdateUser(l.ctx,
			&core.CreateOrUpdateUserReq{
				Id:       0,
				Username: req.Username,
				Password: req.Password,
				Email:    req.Email,
				Nickname: req.Username,
			})
		if err != nil {
			l.Logger.Error("register logic: create user err: ", err.Error())
			return nil, err
		}
		resp = &types.BaseMsgResp{
			Msg: l.svcCtx.Trans.Trans(l.lang, user.Msg),
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(enum.InvalidArgument,
			l.svcCtx.Trans.Trans(l.lang, "login.wrongCaptcha"))
	}
}
