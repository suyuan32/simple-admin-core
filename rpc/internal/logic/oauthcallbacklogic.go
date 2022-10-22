package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/errorx"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
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
		var target model.OauthProvider
		check := l.svcCtx.DB.Where("name = ?", provider).First(&target)
		if check.Error != nil {
			logx.Errorw(logmessage.DatabaseError, logx.Field("detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}

		if check.RowsAffected == 0 {
			logx.Errorw("provider not found", logx.Field("detail", target))
			return nil, status.Error(codes.InvalidArgument, errorx.TargetNotExist)
		}

		providerConfig[provider] = oauth2.Config{
			ClientID:     target.ClientID,
			ClientSecret: target.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   target.AuthURL,
				TokenURL:  target.TokenURL,
				AuthStyle: oauth2.AuthStyle(target.AuthStyle),
			},
			RedirectURL: target.RedirectURL,
			Scopes:      strings.Split(target.Scopes, " "),
		}
		if _, ok := userInfoURL[target.Name]; !ok {
			userInfoURL[target.Name] = target.InfoURL
		}
	}

	// get user information
	content, err := getUserInfo(providerConfig[provider], userInfoURL[provider], in.Code)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// find or register user
	var u userInfo
	err = json.Unmarshal(content, &u)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if u.Email != "" {
		var targetUser model.User
		check := l.svcCtx.DB.Where("email = ?", u.Email).First(&targetUser)
		if check.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, u.Email)
		} else {
			roleName, roleValue, err := getRoleInfo(targetUser.RoleId, l.svcCtx.Redis, l.svcCtx.DB)
			if err != nil {
				return nil, err
			}

			return &core.LoginResp{
				Id:        targetUser.UUID,
				RoleName:  roleName,
				RoleValue: roleValue,
				RoleId:    targetUser.RoleId,
			}, nil
		}
	}

	return nil, status.Error(codes.InvalidArgument, errorx.Failed)
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
