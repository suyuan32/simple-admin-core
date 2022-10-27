package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/model"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailByDictionaryNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailByDictionaryNameLogic {
	return &GetDetailByDictionaryNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDetailByDictionaryNameLogic) GetDetailByDictionaryName(in *core.DictionaryDetailReq) (*core.DictionaryDetailList, error) {
	var dict model.Dictionary
	result := l.svcCtx.DB.Preload("detail").Where("name = ?", in.Name).First(&dict)

	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	resp := &core.DictionaryDetailList{}
	resp.Total = uint64(len(dict.Detail))
	for _, v := range dict.Detail {
		resp.Data = append(resp.Data, &core.DictionaryDetail{
			Id:        uint64(v.ID),
			Title:     v.Title,
			Key:       v.Key,
			Value:     v.Value,
			Status:    v.Status,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			ParentId:  int64(v.DictionaryID),
		})
	}

	return resp, nil
}
