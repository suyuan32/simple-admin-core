package api

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetApiList(l.ctx,
		&core.ApiListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			Path:        req.Path,
			Description: req.Description,
			Method:      req.Method,
			ApiGroup:    req.Group,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ApiListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	// 接口分组翻译
	var group = make(map[string]string)
	for _, v := range data.Data {
		// 检测是否存在分组翻译，避免多次获取
		if _, exist := group[*v.ApiGroup]; !exist {
			group[*v.ApiGroup] = l.svcCtx.Trans.Trans(l.ctx, "apiGroup."+*v.ApiGroup)
		}
		// 因为 ApiInfo 结构体的 Group 字段为指针 因此需要一个额外的临时变量
		trans := group[*v.ApiGroup]
		resp.Data.Data = append(resp.Data.Data,
			types.ApiInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Path:        v.Path,
				Trans:       l.svcCtx.Trans.Trans(l.ctx, *v.Description),
				Description: v.Description,
				Method:      v.Method,
				Group:       &trans, // 分组翻译
				// Group:       v.ApiGroup,
			})
	}
	return resp, nil
}
