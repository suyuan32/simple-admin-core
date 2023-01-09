package user

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewUpdateUserProfileLogic(r *http.Request, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.ProfileReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.UpdateProfile(l.ctx, &core.UpdateProfileReq{
		Id:       l.ctx.Value("userId").(string),
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
	})

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
