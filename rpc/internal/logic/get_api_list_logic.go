package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/model"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiListLogic) GetApiList(in *core.ApiPageReq) (*core.ApiListResp, error) {
	db := l.svcCtx.DB.Model(&model.Api{})
	var apis []model.Api

	if in.Path != "" {
		db = db.Where("path LIKE ?", "%"+in.Path+"%")
	}

	if in.Description != "" {
		db = db.Where("description LIKE ?", "%"+in.Description+"%")
	}

	if in.Method != "" {
		db = db.Where("method = ?", in.Method)
	}

	if in.Group != "" {
		db = db.Where("api_group = ?", in.Group)
	}
	resp := &core.ApiListResp{}
	var total int64
	db.Count(&total)
	resp.Total = uint64(total)

	result := db.Limit(int(in.Page.PageSize)).Offset(int((in.Page.Page - 1) * in.Page.PageSize)).
		Order("api_group desc").Find(&apis)

	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	for _, v := range apis {
		resp.Data = append(resp.Data, &core.ApiInfo{
			Id:          uint64(v.ID),
			CreatedAt:   v.CreatedAt.UnixMilli(),
			Path:        v.Path,
			Description: v.Description,
			Group:       v.ApiGroup,
			Method:      v.Method,
		})
	}
	return resp, nil
}
