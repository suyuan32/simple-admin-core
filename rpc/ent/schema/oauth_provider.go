package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type OauthProvider struct {
	ent.Schema
}

func (OauthProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Comment("the provider's name | 提供商名称"),
		field.String("client_id").Comment("the client id | 客户端 id"),
		field.String("client_secret").Comment("the client secret | 客户端密钥"),
		field.String("redirect_url").Comment("the redirect url | 跳转地址"),
		field.String("scopes").Comment("the scopes | 权限范围"),
		field.String("auth_url").Comment("the auth url of the provider | 认证地址"),
		field.String("token_url").Comment("the token url of the provider | 获取 token地址"),
		field.Uint64("auth_style").Comment("the auth style, 0: auto detect 1: third party log in 2: log in with username and password"),
		field.String("info_url").Comment("the URL to request user information by token | 用户信息请求地址"),
	}
}

func (OauthProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (OauthProvider) Edges() []ent.Edge {
	return nil
}

func (OauthProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_oauth_providers"},
	}
}
