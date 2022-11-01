package logic

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/errorx"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

var providerConfig = make(map[string]oauth2.Config)

// userInfoURL used to store infoURL in database | 用来存储获取用户信息网址数据
var userInfoURL = make(map[string]string)

type OauthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthLoginLogic) OauthLogin(in *core.OauthLoginReq) (*core.OauthRedirectResp, error) {
	var provider model.OauthProvider
	check := l.svcCtx.DB.Where("name = ?", in.Provider).First(&provider)
	if check.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", check.Error.Error()))
		return nil, status.Error(codes.Internal, check.Error.Error())
	}

	if check.RowsAffected == 0 {
		logx.Errorw("Provider not found", logx.Field("detail", in))
		return nil, status.Error(codes.NotFound, errorx.TargetNotExist)
	}

	var config oauth2.Config
	if v, ok := providerConfig[provider.Name]; ok {
		config = v
	} else {
		providerConfig[provider.Name] = oauth2.Config{
			ClientID:     provider.ClientID,
			ClientSecret: provider.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   provider.AuthURL,
				TokenURL:  provider.TokenURL,
				AuthStyle: oauth2.AuthStyle(provider.AuthStyle),
			},
			RedirectURL: provider.RedirectURL,
			Scopes:      strings.Split(provider.Scopes, " "),
		}
		config = providerConfig[provider.Name]
	}

	if _, ok := userInfoURL[provider.Name]; !ok {
		userInfoURL[provider.Name] = provider.InfoURL
	}

	url := config.AuthCodeURL(in.State)

	return &core.OauthRedirectResp{Url: url}, nil
}
