package member

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/member"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type GetMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberListLogic {
	return &GetMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberListLogic) GetMemberList(in *core.MemberListReq) (*core.MemberListResp, error) {
	var predicates []predicate.Member
	if in.Username != "" {
		predicates = append(predicates, member.UsernameContains(in.Username))
	}
	if in.Nickname != "" {
		predicates = append(predicates, member.NicknameContains(in.Nickname))
	}
	if in.Mobile != "" {
		predicates = append(predicates, member.MobileContains(in.Mobile))
	}
	if in.Email != "" {
		predicates = append(predicates, member.EmailContains(in.Email))
	}
	if in.RankId != 0 {
		predicates = append(predicates, member.RankIDEQ(in.RankId))
	}

	result, err := l.svcCtx.DB.Member.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.MemberListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MemberInfo{
			Id:        v.ID.String(),
			CreatedAt: v.CreatedAt.UnixMilli(),
			Status:    uint32(v.Status),
			Username:  v.Username,
			Password:  v.Password,
			Nickname:  v.Nickname,
			RankId:    v.RankID,
			Mobile:    v.Mobile,
			Email:     v.Email,
			Avatar:    v.Avatar,
		})
	}

	return resp, nil
}
