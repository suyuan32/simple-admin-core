package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(in *core.DictionaryPageReq) (*core.DictionaryList, error) {
	var dict []model.Dictionary
	db := l.svcCtx.DB.Model(&model.Dictionary{})

	if in.Name != "" {
		db = db.Where("name LIKE ?", "%"+in.Name+"%")
	}

	if in.Title != "" {
		db = db.Where("title LIKE ?", "%"+in.Title+"%")
	}
	resp := &core.DictionaryList{}
	var total int64
	db.Count(&total)
	resp.Total = uint64(total)
	result := db.Limit(int(in.PageSize)).Offset(int((in.Page - 1) * in.PageSize)).Find(&dict)

	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	for _, v := range dict {
		resp.Data = append(resp.Data, &core.DictionaryInfo{
			Id:        uint64(v.ID),
			Title:     v.Title,
			Name:      v.Name,
			Status:    v.Status,
			Desc:      v.Desc,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
