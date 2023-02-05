package member

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

type GetMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMemberListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMemberListLogic {
	return &GetMemberListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMemberListLogic) GetMemberList(req *types.MemberListReq) (resp *types.MemberListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMemberList(l.ctx,
		&core.MemberListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Username: req.Username,
			Nickname: req.Nickname,
			Mobile:   req.Mobile,
			Email:    req.Email,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MemberListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MemberInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:   v.Status,
				Username: v.Username,
				Password: "",
				Nickname: v.Nickname,
				RankId:   v.RankId,
				Mobile:   v.Mobile,
				Email:    v.Email,
				Avatar:   v.Avatar,
			})
	}
	return resp, nil
}
