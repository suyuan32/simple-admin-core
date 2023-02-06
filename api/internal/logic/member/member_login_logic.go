package member

import (
	"context"
	"net/http"
	"time"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewMemberLoginLogic(r *http.Request, svcCtx *svc.ServiceContext) *MemberLoginLogic {
	return &MemberLoginLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *MemberLoginLogic) MemberLogin(req *types.MemberLoginReq) (resp *types.MemberLoginResp, err error) {
	result, err := l.svcCtx.CoreRpc.MemberLogin(l.ctx, &core.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	token, err := utils.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, result.Id, "rankId", time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, int64(result.RankId))
	if err != nil {
		return nil, err
	}

	// add token into database
	expiredAt := time.Now().Add(time.Second * 259200).Unix()
	_, err = l.svcCtx.CoreRpc.CreateOrUpdateToken(l.ctx, &core.TokenInfo{
		Id:        "",
		CreatedAt: 0,
		Uuid:      result.Id,
		Token:     token,
		Source:    "core_member",
		Status:    1,
		ExpiredAt: expiredAt,
	})

	if err != nil {
		return nil, err
	}

	return &types.MemberLoginResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
			Data: "",
		},
		Data: types.MemberLoginRespInfo{
			Id:       result.Id,
			Avatar:   result.Avatar,
			Nickname: result.Nickname,
			RankId:   result.RankId,
			Token:    token,
			Expire:   uint64(expiredAt),
		},
	}, nil
}
