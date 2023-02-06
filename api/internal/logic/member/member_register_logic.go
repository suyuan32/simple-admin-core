package member

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewMemberRegisterLogic(r *http.Request, svcCtx *svc.ServiceContext) *MemberRegisterLogic {
	return &MemberRegisterLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *MemberRegisterLogic) MemberRegister(req *types.MemberRegisterReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateMember(l.ctx,
		&core.MemberInfo{
			Id:       "",
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Nickname: req.Username,
			Status:   1,
			RankId:   1,
		})
	if err != nil {
		l.Logger.Error("register logic: create user err: ", err.Error())
		return nil, err
	}
	resp = &types.BaseMsgResp{
		Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg),
	}
	return resp, nil
}
