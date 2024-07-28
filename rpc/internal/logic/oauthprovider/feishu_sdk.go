package oauthprovider

import (
	"context"
	"encoding/json"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkauthen "github.com/larksuite/oapi-sdk-go/v3/service/authen/v1"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/oauth2"
)

var failedError error = errorx.NewInvalidArgumentError("failed to get oauth data")

// feishu funcs

func GetFeishuUserInfo(c oauth2.Config, code string) ([]byte, error) {
	client := lark.NewClient(c.ClientID, c.ClientSecret)

	req := larkauthen.NewCreateOidcAccessTokenReqBuilder().
		Body(larkauthen.NewCreateOidcAccessTokenReqBodyBuilder().
			GrantType(`authorization_code`).
			Code(code).
			Build()).
		Build()

	// 发起请求
	infoResp, err := client.Authen.OidcAccessToken.Create(context.Background(), req)

	if err != nil {
		logx.Error("failed to get access token from feishu", logx.Field("detail", err))
		return nil, failedError
	}

	if !infoResp.Success() {
		logx.Error("failed to get access token from feishu", logx.Field("detail", err))
		return nil, failedError
	}

	resp, err := client.Authen.UserInfo.Get(context.Background(), larkcore.WithUserAccessToken(*infoResp.Data.AccessToken))

	if err != nil {
		logx.Error("failed to get access token from feishu", logx.Field("detail", err))
		return nil, failedError
	}

	if !resp.Success() {
		logx.Error("failed to get access token from feishu", logx.Field("detail", err))
		return nil, failedError
	}

	uInfo := &userInfo{
		NickName: *resp.Data.Name,
		Email:    *resp.Data.Email,
		Picture:  *resp.Data.AvatarUrl,
		Mobile:   *resp.Data.Mobile,
	}

	result, err := json.Marshal(uInfo)

	if err != nil {
		logx.Error("failed to get access token from feishu", logx.Field("detail", err))
		return nil, failedError
	}

	return result, nil
}
