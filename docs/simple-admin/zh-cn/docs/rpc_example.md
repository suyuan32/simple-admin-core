# 3分钟开发 RPC 服务

> 首先确认你安装了以下软件:
- simple-admin-tool (goctls)


## 创建 RPC 项目
> 创建 messaging 服务 (--Ent=true表示使用Ent)
```shell
goctls rpc new messaging --Ent=true
```
> 你可以看到如下项目结构 

![Example](../../assets/example_rpc_struct.png)

> 你需要运行以下命令来替换默认的go-zero.

```shell
cd messaging

goctls migrate --zero-version v1.4.2 --tool-version v0.0.9
```

> 版本号可以去Github查看最新版本. \
运行完命令后 go.mod 会变成:

```text
module messaging

go 1.19

require (
	entgo.io/ent v0.11.4
	github.com/suyuan32/simple-admin-core v0.1.6
	github.com/zeromicro/go-zero v1.4.2
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	ariga.io/atlas v0.7.3-0.20221011160332-3ca609863edd // indirect
	ariga.io/entcache v0.1.0 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/casbin/casbin/v2 v2.52.1 // indirect
	github.com/casbin/ent-adapter v0.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/openzipkin/zipkin-go v0.4.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/v3 v3.5.5 // indirect
	go.opentelemetry.io/otel v1.11.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.11.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.11.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.11.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.11.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.11.0 // indirect
	go.opentelemetry.io/otel/sdk v1.11.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/automaxprocs v1.5.1 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20220602131408-e326c6e8e9c8 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.22.9 // indirect
	k8s.io/apimachinery v0.22.9 // indirect
	k8s.io/client-go v0.22.9 // indirect
	k8s.io/klog/v2 v2.40.1 // indirect
	k8s.io/utils v0.0.0-20220706174534-f6158b442e7c // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace github.com/zeromicro/go-zero v1.4.2 => github.com/suyuan32/simple-admin-tools v0.0.9
```

> 注意 migrate 不要多次运行，会导致多次添加 replace 

然后编辑 etc/messaging.yaml

```yaml
Name: messaging.rpc
ListenOn: 0.0.0.0:9106

DatabaseConf:
  Type: mysql
  Host: 127.0.0.1
  Port: 3306
  DBName: simple_admin
  Username: # set your username
  Password: # set your password
  MaxOpenConn: 100
  SSLMode: false
  CacheTime: 5

RedisConf:
  Host: 127.0.0.1:6379
  Type: node

Log:
  ServiceName: messagingRpcLogger
  Mode: file
  Path: /home/data/logs/messaging/rpc
  Encoding: json
  Level: info
  Compress: false
  KeepDays: 7
  StackCoolDownMillis: 100

Prometheus:
  Host: 0.0.0.0
  Port: 4006
  Path: /metrics

```

> 编辑 schema

进入目录 ent/schema, 修改 messaging.go, 改名为 email.go 添加 mixin 和字段 address 和 uuid

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type Email struct {
	ent.Schema
}

// Fields of the Email.
func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Comment("email address"),
		field.String("uuid").Comment("user uuid"),
	}
}

func (Email) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return nil
}

```

> 生成 Ent 代码

```shell
make gen-ent
```

> 生成 Email 逻辑代码

model=Email 表示只生成 schema 为 Email 的代码， 为空则全部生成

```shell
make gen-rpc-ent-logic model=Email

# 可能需要运行下
go mod tidy 
```

![logic](../../assets/ent_gen_logic.png)

可以看到 crud 代码已生成

> 然后代码就可以运行啦 !

```shell
go run messaging.go -f etc/messaging.yaml
```

> 如果看到
```shell
Starting server at 127.0.0.1:9106...
```

说明运行成功. 注意后续还需要修改数据库初始化函数，参考 [simple admin file](https://github.com/suyuan32/simple-admin-file/blob/master/api/internal/logic/file/init_database_logic.go)

> simple admin core 的 API 中如何远程调用该 RPC 

> 添加 config
```go
package config

import (
	"github.com/suyuan32/simple-admin-tools/plugins/registry/consul"
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf 
	Auth          rest.AuthConf
	RedisConf     redis.RedisConf    
	CoreRpc       zrpc.RpcClientConf 
	MessagingRpc  zrpc.RpcClientConf 
	Captcha       Captcha          
	DB            gormsql.GORMConf  
}

type Captcha struct {
	KeyLong   int   // captcha length
	ImgWidth  int   // captcha width
	ImgHeight int   // captcha height
}
```

> 小型网站直接使用
>
> MessagingRpc:
>  Endpoints:
>   - 127.0.0.1:9106
>
> 的方式直连，不需要服务发现， Endpoints 可以有多个

> 添加 example rpc 
### 修改 service context
```go
package svc

import (
	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   coreclient.Core
	Redis     *redis.Redis
	Casbin    *casbin.Enforcer
	Trans     *i18n.Translator
	MessagingRpc messagingclient.Messaging
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize redis
	rds := c.RedisConf.NewRedis()
	if !rds.Ping() {
		logx.Error("initialize redis failed")
		return nil
	}
	logx.Info("initialize redis connection successfully")

	// initialize casbin connection
	cbn, err := c.CasbinConf.NewCasbin(c.DatabaseConf.Type, c.DatabaseConf.GetDSN())
	if err != nil {
		logx.Errorw("Initialize casbin failed", logx.Field("detail", err.Error()))
		return nil
	}

	// initialize translator
	trans := &i18n.Translator{}
	trans.NewBundle(i18n.LocaleFS)
	trans.NewTranslator()

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		CoreRpc:   coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		MessagingRpc: messagingclient.NewMessaging(zrpc.MustNewClient(c.MessagingRpc)),
		Redis:     rds,
		Casbin:    cbn,
		Trans:     trans,
	}
}

```

> 然后在 logic 直接调用 l.svcCtx.MessagingRpc 即可

