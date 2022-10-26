package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeMenuParamListByMenuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeMenuParamListByMenuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeMenuParamListByMenuIdLogic {
	return &GeMenuParamListByMenuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeMenuParamListByMenuIdLogic) GeMenuParamListByMenuId(in *core.IDReq) (*core.MenuParamListResp, error) {
	var paramsList []model.MenuParam
	result := l.svcCtx.DB.Where("menu_id = ?", in.ID).Find(&paramsList)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	resp := &core.MenuParamListResp{}
	resp.Total = uint64(result.RowsAffected)
	for _, v := range paramsList {
		resp.Data = append(resp.Data, &core.MenuParamResp{
			Id:        uint64(v.ID),
			MenuId:    uint64(v.MenuId),
			Type:      v.Type,
			Key:       v.Key,
			Value:     v.Value,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
