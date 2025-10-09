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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register creates a new user account with username and email
func (l *RegisterLogic) Register(in *core.RegisterReq) (*core.BaseResp, error) {
	// Check if email already exists
	exists, err := l.svcCtx.DB.User.Query().
		Where(user.EmailEQ(in.Email)).
		Exist(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if exists {
		return nil, errorx.NewInvalidArgumentError("login.emailAlreadyExists")
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

	// Create user using existing CreateUser logic
	createLogic := NewCreateUserLogic(l.ctx, l.svcCtx)
	_, err = createLogic.CreateUser(&core.UserInfo{
		Username: &in.Username,
		Password: &in.Password,
		Email:    &in.Email,
		Status:   pointy.GetPointer(uint32(1)), // Active by default
	})
	if err != nil {
		return nil, err
	}

	return &core.BaseResp{
		Msg: "login.registerSuccessTitle",
	}, nil
}
