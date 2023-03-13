package user

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.User.Create().
		SetUsername(in.Username).
		SetPassword(encrypt.BcryptEncrypt(in.Password)).
		SetNickname(in.Nickname).
		SetEmail(in.Email).
		SetMobile(in.Mobile).
		SetAvatar(in.Avatar).
		AddRoleIDs(in.RoleIds...).
		SetHomePath(in.HomePath).
		SetDescription(in.Description).
		SetDepartmentID(in.DepartmentId).
		AddPositionIDs(in.PositionIds...).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
