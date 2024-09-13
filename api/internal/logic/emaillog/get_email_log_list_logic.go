package emaillog

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailLogListLogic {
	return &GetEmailLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailLogListLogic) GetEmailLogList(req *types.EmailLogListReq) (resp *types.EmailLogListResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetEmailLogList(l.ctx,
		&mcms.EmailLogListReq{
			Page:       req.Page,
			PageSize:   req.PageSize,
			Target:     req.Target,
			Subject:    req.Subject,
			Provider:   req.Provider,
			SendStatus: req.SendStatus,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.EmailLogListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.EmailLogInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Target:     v.Target,
				Subject:    v.Subject,
				Content:    v.Content,
				SendStatus: v.SendStatus,
				Provider:   v.Provider,
			})
	}
	return resp, nil
}
