package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleListLogic) GetRoleList(in *core.PageInfoReq) (*core.RoleListResp, error) {
	var roles []*model.Role
	resp := &core.RoleListResp{}
	var total int64
	l.svcCtx.DB.Model(&model.Role{}).Count(&total)
	resp.Total = uint64(total)
	result := l.svcCtx.DB.Limit(int(in.PageSize)).Offset(int((in.Page - 1) * in.PageSize)).Find(&roles)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	for _, v := range roles {
		resp.Data = append(resp.Data, &core.RoleInfo{
			Id:            uint64(v.ID),
			Name:          v.Name,
			Value:         v.Value,
			DefaultRouter: v.DefaultRouter,
			Status:        v.Status,
			Remark:        v.Remark,
			OrderNo:       v.OrderNo,
			CreatedAt:     v.CreatedAt.UnixMilli(),
		})
	}
	return resp, nil
}
