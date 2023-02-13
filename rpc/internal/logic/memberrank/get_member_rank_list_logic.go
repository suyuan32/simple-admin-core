package memberrank

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/memberrank"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type GetMemberRankListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankListLogic {
	return &GetMemberRankListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRankListLogic) GetMemberRankList(in *core.MemberRankListReq) (*core.MemberRankListResp, error) {
	var predicates []predicate.MemberRank
	if in.Name != "" {
		predicates = append(predicates, memberrank.NameContains(in.Name))
	}
	if in.Description != "" {
		predicates = append(predicates, memberrank.DescriptionContains(in.Description))
	}
	if in.Remark != "" {
		predicates = append(predicates, memberrank.RemarkContains(in.Remark))
	}
	result, err := l.svcCtx.DB.MemberRank.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.MemberRankListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MemberRankInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Name:        v.Name,
			Description: v.Description,
			Remark:      v.Remark,
			Code:        v.Code,
		})
	}

	return resp, nil
}
