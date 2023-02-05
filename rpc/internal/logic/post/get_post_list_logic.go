package post

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/post"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type GetPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostListLogic) GetPostList(in *core.PostListReq) (*core.PostListResp, error) {
	var predicates []predicate.Post
	if in.Name != "" {
		predicates = append(predicates, post.NameContains(in.Name))
	}
	result, err := l.svcCtx.DB.Post.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.PostListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.PostInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			Status:    uint32(v.Status),
			Sort:      v.Sort,
			Name:      v.Name,
			Code:      v.Code,
			Remark:    v.Remark,
		})
	}

	return resp, nil
}
