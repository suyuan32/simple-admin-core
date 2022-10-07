package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/errorx"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
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
			logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", check.Error.Error()))
			return nil, status.Error(codes.Internal, check.Error.Error())
		}

		if check.RowsAffected == 0 {
			logx.Errorw("Provider not found", logx.Field("Detail", target))
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
			Scopes:      strings.Split(target.Scopes, ","),
		}
		if _, ok := userInfoURL[target.Name]; !ok {
			userInfoURL[target.Name] = target.InfoURL
		}
	}

	// get user information
	content, err := getUserInfo(providerConfig[provider], userInfoURL[provider], in.Code)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
			userUUID := uuid.NewString()
			result := l.svcCtx.DB.Save(&model.User{
				Model:    gorm.Model{},
				UUID:     userUUID,
				Username: u.Email,
				Password: util.BcryptEncrypt(u.Email),
				Nickname: userUUID[30:],
				Email:    u.Email,
			})
			if result.Error != nil {
				logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
				return nil, status.Error(codes.Internal, result.Error.Error())
			}

			roleName, roleValue, err := getRoleInfo(2, l.svcCtx.Redis, l.svcCtx.DB)
			if err != nil {
				return nil, err
			}

			return &core.LoginResp{
				Id:        userUUID,
				RoleName:  roleName,
				RoleValue: roleValue,
				RoleId:    2,
			}, nil
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

	response, err := http.Get(infoURL + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
