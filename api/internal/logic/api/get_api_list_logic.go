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
		// 在角色管理页面调用该接口时为了获取所有接口，每页大小被设置为10000
		// 因此可以根据每页数据大小来判断，该请求是否为来自角色管理页面
		// 如果是来自API管理的请求则不去获取分组翻译，直接显示原本字段
		if req.PageSize > 1000 {
			// 检测是否存在分组翻译，避免多次获取
			if _, exist := group[*v.ApiGroup]; !exist {
				group[*v.ApiGroup] = l.svcCtx.Trans.Trans(l.ctx, "apiGroup."+*v.ApiGroup)
			}
			// 直接将分组字段更新为分组翻译内容
			*v.ApiGroup = group[*v.ApiGroup]
		}
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
				Group:       v.ApiGroup, // 根据页大小决定显示原本字段还是分组字段
				Method:      v.Method,
			})
	}
	return resp, nil
}
