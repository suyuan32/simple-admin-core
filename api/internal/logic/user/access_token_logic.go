package user

import (
	"context"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/orm/ent/entctx/rolectx"
	"github.com/suyuan32/simple-admin-common/orm/ent/entctx/userctx"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessTokenLogic {
	return &AccessTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *AccessTokenLogic) AccessToken() (resp *types.RefreshTokenResp, err error) {
	userId, err := userctx.GetUserIDFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	roleIds, err := rolectx.GetRoleIDFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	userData, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.UUIDReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	if userData.Status != nil && *userData.Status != uint32(common.StatusNormal) {
		return nil, errorx.NewApiUnauthorizedError(i18n.Failed)
	}

	token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
		int64(l.svcCtx.Config.ProjectConf.AccessTokenPeriod)*60*60, jwt.WithOption("userId", userId), jwt.WithOption("roleId",
			strings.Join(roleIds, ",")), jwt.WithOption("deptId", userData.DepartmentId))
	if err != nil {
		return nil, err
	}

	// add token into database
	expiredAt := time.Now().Add(time.Hour * time.Duration(l.svcCtx.Config.ProjectConf.AccessTokenPeriod)).UnixMilli()
	_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
		Uuid:      &userId,
		Token:     pointy.GetPointer(token),
		Source:    pointy.GetPointer("core_user_access_token"),
		Status:    pointy.GetPointer(uint32(common.StatusNormal)),
		Username:  userData.Username,
		ExpiredAt: pointy.GetPointer(expiredAt),
	})

	return &types.RefreshTokenResp{
		BaseDataInfo: types.BaseDataInfo{Msg: i18n.Success},
		Data:         types.RefreshTokenInfo{Token: token, ExpiredAt: expiredAt},
	}, nil
}
