package oauthprovider

import (
	"context"
	"strings"

	"golang.org/x/oauth2"

	"github.com/suyuan32/simple-admin-core/rpc/ent/oauthprovider"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
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
	p, err := l.svcCtx.DB.OauthProvider.Query().Where(oauthprovider.NameEQ(in.Provider)).First(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	var config oauth2.Config
	if v, ok := providerConfig[p.Name]; ok {
		config = v
	} else {
		providerConfig[p.Name] = oauth2.Config{
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   p.AuthURL,
				TokenURL:  p.TokenURL,
				AuthStyle: oauth2.AuthStyle(p.AuthStyle),
			},
			RedirectURL: p.RedirectURL,
			Scopes:      strings.Split(p.Scopes, " "),
		}
		config = providerConfig[p.Name]
	}

	if _, ok := userInfoURL[p.Name]; !ok {
		userInfoURL[p.Name] = p.InfoURL
	}

	url := config.AuthCodeURL(in.State)

	return &core.OauthRedirectResp{Url: url}, nil
}
