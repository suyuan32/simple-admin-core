package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *core.GetUserListReq) (*core.UserListResp, error) {
	db := l.svcCtx.DB.Model(&model.User{})
	var users []model.User

	if in.Username != "" {
		db = db.Where("username LIKE ?", "%"+in.Username+"%")
	}

	if in.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+in.Nickname+"%")
	}

	if in.Email != "" {
		db = db.Where("email = ?", in.Email)
	}

	if in.Mobile != "" {
		db = db.Where("group = ?", in.Mobile)
	}

	if in.RoleId != 0 {
		db = db.Where("role_id = ?", in.RoleId)
	}

	result := db.Limit(int(in.PageSize)).Offset(int((in.Page - 1) * in.PageSize)).
		Order("role_id, id desc").Find(&users)

	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	resp := &core.UserListResp{}
	resp.Total = uint32(result.RowsAffected)
	for _, v := range users {
		resp.Data = append(resp.Data, &core.UserInfoResp{
			Id:       uint64(v.ID),
			Avatar:   v.Avatar,
			RoleId:   v.RoleId,
			Mobile:   v.Mobile,
			Email:    v.Email,
			Status:   v.Status,
			Username: v.Username,
			UUID:     v.UUID,
			Nickname: v.Nickname,
			CreateAt: v.CreatedAt.UnixMilli(),
			UpdateAt: v.UpdatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
