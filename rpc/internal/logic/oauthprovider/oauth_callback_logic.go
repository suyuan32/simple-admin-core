package oauthprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	user2 "github.com/suyuan32/simple-admin-core/rpc/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type userInfo struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Picture  string `json:"picture"`
}

func NewOauthCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthCallbackLogic) OauthCallback(in *core.CallbackReq) (*core.LoginResp, error) {
	provider := strings.Split(in.State, "-")[1]
	if _, ok := providerConfig[provider]; !ok {
		p, err := l.svcCtx.DB.OauthProvider.Query().Where(oauthprovider.NameEQ(provider)).First(l.ctx)
		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		providerConfig[provider] = oauth2.Config{
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
		if _, ok := userInfoURL[p.Name]; !ok {
			userInfoURL[p.Name] = p.InfoURL
		}
	}

	// get user information
	content, err := getUserInfo(providerConfig[provider], userInfoURL[provider], in.Code)
	if err != nil {
		return nil, statuserr.NewInvalidArgumentError(err.Error())
	}

	// find or register user
	var u userInfo
	err = json.Unmarshal(content, &u)
	if err != nil {
		return nil, statuserr.NewInternalError(err.Error())
	}

	if u.Email != "" {
		targetUser, err := l.svcCtx.DB.User.Query().Where(user.EmailEQ(u.Email)).First(l.ctx)
		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError("login.userNotExist")
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		roleName, roleValue, err := user2.GetRoleInfo(targetUser.RoleID, l.svcCtx.Redis, l.svcCtx.DB, l.ctx)
		if err != nil {
			return nil, err
		}

		return &core.LoginResp{
			Id:        targetUser.ID.String(),
			RoleName:  roleName,
			RoleValue: roleValue,
			RoleId:    targetUser.RoleID,
		}, nil
	}

	return nil, status.Error(codes.InvalidArgument, i18n.Failed)
}

func getUserInfo(c oauth2.Config, infoURL string, code string) ([]byte, error) {
	token, err := c.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	var response *http.Response
	if c.Endpoint.AuthStyle == 1 {
		response, err = http.Get(infoURL + token.AccessToken)
		if err != nil {
			return nil, fmt.Errorf("failed getting user info: %s", err.Error())
		}
	} else if c.Endpoint.AuthStyle == 2 {
		client := &http.Client{}
		request, err := http.NewRequest("GET", infoURL, nil)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+token.AccessToken)

		response, err = client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("failed getting user info: %s", err.Error())
		}
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
