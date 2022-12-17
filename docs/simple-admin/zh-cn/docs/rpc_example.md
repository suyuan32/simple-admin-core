# 3分钟开发 RPC 服务

> 首先确认你安装了以下软件:
- simple-admin-tool (goctls) v0.1.3


## 创建 RPC 基本项目
> 创建 example 服务 
> 
```shell
goctls rpc new example --ent=true --module_name=github.com/suyuan32/simple-admin-example-rpc --go_zero_version=v1.4.2 --tool_version=v0.1.2 --port=8080 --gitlab=true
```

### 参数介绍

| 参数              | 介绍                     | 使用方法                                                                                               |
|-----------------|------------------------|----------------------------------------------------------------------------------------------------|
| ent             | 是否启用 ent               | true 为启用                                                                                           |
| module_name     | go.mod 中的module名称      | 如果项目需要被在外部import，需要像上面例子设置为github或者其他地方的仓库网址， 为空则只在本地使用                                            |
| go_zero_version | go zero版本              | 需要到[go-zero](https://github.com/zeromicro/go-zero/releases)查看最新release                             |
| tool_version    | simple admin tools 版本号 | 需要到[tool](https://github.com/suyuan32/simple-admin-tools/releases)查看simple admin  tools 最新 release |
| gitlab          | 是否生成 gitlab-ci.yml     | true 为生成                                                                                           |
| port            | 端口号                    | 服务暴露的端口号                                                                                           |

详细参数请在命令行查看 `goctls rpc new --help`

> 你可以看到如下项目结构 

![Example](../../assets/example_rpc_struct.png)

然后编辑 etc/example.yaml

```yaml
Name: example.rpc
ListenOn: 0.0.0.0:8080

DatabaseConf:
  Type: mysql
  Host: 127.0.0.1
  Port: 3306
  DBName: simple_admin
  Username: root # set your username
  Password: "123456" # set your password
  MaxOpenConn: 100
  SSLMode: disable
  CacheTime: 5

RedisConf:
  Host: 127.0.0.1:6379
  Type: node

Log:
  ServiceName: exampleRpcLogger
  Mode: file
  Path: /home/ryan/data/logs/example/rpc
  Encoding: json
  Level: info
  Compress: false
  KeepDays: 7
  StackCoolDownMillis: 100

Prometheus:
  Host: 0.0.0.0
  Port: 4001
  Path: /metrics

```

> 编辑 schema

进入目录 ent/schema, 修改 example.go, 改名为 student.go 添加 mixin 和字段 address 和 uuid

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age"),
	}
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}


```

> 生成 Ent 代码

```shell
make gen-ent
```

> 生成 Student 逻辑代码

```shell
goctls rpc ent --schema=./ent/schema  --style=go_zero --multiple=false --service_name=example --o=./ --model=Student --group=student

make gen-rpc
```

### 参数介绍

| 参数             | 介绍           | 使用方法                                         |
|----------------|--------------|----------------------------------------------|
| schema         | schema文件地址   | 输入Ent schema 文件夹相对路径                         |
| style          | 文件名格式        | go_zero为蛇形格式                                 |
| service_name   | 服务名称         | 和new 时的名称相同，如example.go的serviceName是 example |
| o              | 输出位置         | 文件输出位置，可以为相对路径，指向main文件目录                    |
| model          | 模型名称         | schema中内部struct名称，如example中的Student          |
| search_key_num | 搜索字段数量（默认为3） | 列表搜索字段数量，只能自动生成string的字段                     |
| group          | 分组名称         | 分组名称用于将不同logic文件放到不同文件夹                      |

详细参数请在命令行查看 `goctls rpc ent --help`

> 快捷命令：gen-rpc-ent-logic model=Student 表示只生成 schema 为 Student 的代码， 为空则全部生成 group 为分组文件夹名称

```shell
make gen-rpc-ent-logic model=Student group=student

make gen-rpc

# 可能需要运行下
go mod tidy 
```

![logic](../../assets/ent_gen_logic.png)

可以看到 crud 代码已生成

> proto 文件代码

```protobuf
syntax = "proto3";

package example;
option go_package="./example";

message Empty {}

message IDReq {
  uint64 id = 1;
}

message IDsReq {
  repeated uint64 ids = 1;
}

message UUIDReq {
  string uuid = 1;
}

message BaseResp {
  string msg = 1;
}

message PageInfoReq {
  uint64 page = 1;
  uint64 page_size = 2;
}


// Student message

message StudentInfo {
  uint64 id = 1;
  int64 created_at = 2;
  int64 updated_at = 3;
  string name = 4;
  int64 age = 5;
}

message StudentListResp {
  uint64 total = 1;
  repeated StudentInfo data = 2;
}

message StudentPageReq {
  uint64 page = 1;
  uint64 page_size = 2;
  string name = 3;
}

service Example {
  // group: base
  rpc initDatabase (Empty) returns (BaseResp);

  // Student management
  // group: student
  rpc createOrUpdateStudent (StudentInfo) returns (BaseResp);
  // group: student
  rpc getStudentList (StudentPageReq) returns (StudentListResp);
  // group: student
  rpc deleteStudent (IDReq) returns (BaseResp);
  // group: student
  rpc batchDeleteStudent (IDsReq) returns (BaseResp);
}

```

使用group注释给 rpc 分组

> 然后代码就可以运行啦 !

```shell
go run example.go -f etc/example.yaml
```

> 如果看到
```shell
Starting server at 127.0.0.1:8080...
```
说明运行成功. 注意后续还需要修改数据库初始化函数，参考 [simple admin file](https://github.com/suyuan32/simple-admin-file/blob/master/api/internal/logic/file/init_database_logic.go)

> 项目地址 https://github.com/suyuan32/simple-admin-example-rpc

> simple admin example api 中如何远程调用该 RPC 

> 添加 config
```go
package config

import (
	"github.com/suyuan32/simple-admin-core/pkg/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
	CasbinConf   config.CasbinConf
	ExampleRpc   zrpc.RpcClientConf
}

```

> 小型网站直接使用
>
> ExampleRpc:
>  Endpoints:
>   - 127.0.0.1:8080
>
> 的方式直连，不需要服务发现， Endpoints 可以有多个

> 添加 example rpc 
### 修改 service context
```go
package svc

import (
	"github.com/suyuan32/simple-admin-example-rpc/exampleclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-example-api/internal/config"
	i18n2 "github.com/suyuan32/simple-admin-example-api/internal/i18n"
	"github.com/suyuan32/simple-admin-example-api/internal/middleware"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	ExampleRpc exampleclient.Example
	Casbin     *casbin.Enforcer
	Authority  rest.Middleware
	Trans      *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.NewRedis()
	if !rds.Ping() {
		logx.Error("initialize redis failed")
		return nil
	}

	cbn, err := c.CasbinConf.NewCasbin(c.DatabaseConf.Type, c.DatabaseConf.GetDSN())
	if err != nil {
		logx.Errorw("initialize casbin failed", logx.Field("detail", err.Error()))
		return nil
	}

	trans := &i18n.Translator{}
	trans.NewBundle(i18n2.LocaleFS)
	trans.NewTranslator()

	return &ServiceContext{
		Config:     c,
		Authority:  middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		Trans:      trans,
		ExampleRpc: exampleclient.NewExample(zrpc.MustNewClient(c.ExampleRpc)),
	}
}
```

> 然后在 logic 直接调用 l.svcCtx.ExampleRpc 即可

> simple admin example api 地址 https://github.com/suyuan32/simple-admin-example-api