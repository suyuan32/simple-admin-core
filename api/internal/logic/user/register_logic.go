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
		user, err := l.svcCtx.CoreRpc.CreateUser(l.ctx,
			&core.UserInfo{
				Id:           "",
				Username:     req.Username,
				Password:     req.Password,
				Email:        req.Email,
				Nickname:     req.Username,
				Status:       1,
				HomePath:     "/dashboard",
				RoleId:       1,
				DepartmentId: 1,
				PositionId:   1,
			})
		if err != nil {
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
