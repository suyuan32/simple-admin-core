package member

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateMemberLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateMemberLogic {
	return &CreateOrUpdateMemberLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateMemberLogic) CreateOrUpdateMember(req *types.CreateOrUpdateMemberReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateMember(l.ctx,
		&core.MemberInfo{
			Id:       req.Id,
			Status:   req.Status,
			Username: req.Username,
			Password: req.Password,
			Nickname: req.Nickname,
			RankId:   req.RankId,
			Mobile:   req.Mobile,
			Email:    req.Email,
			Avatar:   req.Avatar,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
