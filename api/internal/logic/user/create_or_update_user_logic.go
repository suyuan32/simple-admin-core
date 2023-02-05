package user

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateUserLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateUserLogic {
	return &CreateOrUpdateUserLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateUserLogic) CreateOrUpdateUser(req *types.CreateOrUpdateUserReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateUser(l.ctx, &core.CreateOrUpdateUserReq{
		Id:           req.Id,
		Avatar:       req.Avatar,
		RoleId:       req.RoleId,
		Mobile:       req.Mobile,
		Email:        req.Email,
		Status:       req.Status,
		Username:     req.Username,
		Nickname:     req.Nickname,
		Password:     req.Password,
		Description:  req.Description,
		HomePath:     req.HomePath,
		DepartmentId: req.DepartmentId,
		PositionId:   req.PositionId,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
