package user

import (
	"context"

	"github.com/chimerakang/simple-admin-common/utils/pointy"
	"github.com/chimerakang/simple-admin-common/utils/uuidx"

	"github.com/chimerakang/simple-admin-core/rpc/ent"
	"github.com/chimerakang/simple-admin-core/rpc/ent/user"

	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *core.UUIDReq) (*core.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(uuidx.ParseUUIDString(in.Id))).WithRoles().WithDepartments().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.UserInfo{
		Nickname:       &result.Nickname,
		Avatar:         &result.Avatar,
		RoleIds:        GetRoleIds(result.Edges.Roles),
		RoleName:       GetRoleNames(result.Edges.Roles),
		Mobile:         &result.Mobile,
		Email:          &result.Email,
		Status:         pointy.GetPointer(uint32(result.Status)),
		Id:             pointy.GetPointer(result.ID.String()),
		Username:       &result.Username,
		HomePath:       &result.HomePath,
		Password:       &result.Password,
		Description:    &result.Description,
		DepartmentId:   &result.DepartmentID,
		DepartmentName: &result.Edges.Departments.Name,
		CreatedAt:      pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:      pointy.GetPointer(result.UpdatedAt.UnixMilli()),
	}, nil
}

func GetRoleIds(data []*ent.Role) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}

func GetRoleNames(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Name)
	}
	return codes
}

func GetRoleCodes(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Code)
	}
	return codes
}

func GetPositionIds(data []*ent.Position) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}
