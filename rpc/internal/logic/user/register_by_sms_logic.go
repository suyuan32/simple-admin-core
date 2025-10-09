package user

import (
	"context"

	"github.com/chimerakang/simple-admin-common/utils/pointy"
	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterBySmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterBySmsLogic {
	return &RegisterBySmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RegisterBySms creates a new user account with SMS verification
func (l *RegisterBySmsLogic) RegisterBySms(in *core.RegisterBySmsReq) (*core.BaseResp, error) {
	// Check if phone number already exists
	exists, err := l.svcCtx.DB.User.Query().
		Where(user.MobileEQ(in.PhoneNumber)).
		Exist(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if exists {
		return nil, errorx.NewInvalidArgumentError("login.phoneAlreadyExists")
	}

	// Check if username already exists
	exists, err = l.svcCtx.DB.User.Query().
		Where(user.UsernameEQ(in.Username)).
		Exist(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if exists {
		return nil, errorx.NewInvalidArgumentError("login.userAlreadyExists")
	}

	// Note: SMS captcha verification happens in API layer

	// Create user
	createLogic := NewCreateUserLogic(l.ctx, l.svcCtx)
	_, err = createLogic.CreateUser(&core.UserInfo{
		Username: &in.Username,
		Password: &in.Password,
		Mobile:   &in.PhoneNumber,
		Status:   pointy.GetPointer(uint32(1)),
	})
	if err != nil {
		return nil, err
	}

	return &core.BaseResp{
		Msg: "login.registerSuccessTitle",
	}, nil
}
