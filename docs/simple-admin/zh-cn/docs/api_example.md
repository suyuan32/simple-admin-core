# 3 分钟开发 API 服务

首先确认你安装了以下软件:
- simple-admin-tool (goctls) v0.1.0 +


## 创建 API 项目
创建 example
```shell
goctls api new example --i18n=true --casbin=true --goZeroVersion=v1.4.2 --toolVersion=v0.0.9 --transErr=true --moduleName=github.com/suyuan32/simple-admin-example-api --port=8081

```

### 参数介绍

| 参数            | 介绍                     | 使用方法                                                    |
|---------------|------------------------|---------------------------------------------------------|
| i18n          | 是否启用 i18n              | true 为启用                                                |
| casbin        | 是否启用 casbin            | true 为启用                                                |
| moduleName    | go.mod 中的module名称      | 如果项目需要被在外部import，需要像上面例子设置为github或者其他地方的仓库网址， 为空则只在本地使用 |
| goZeroVersion | go zero版本              | 需要到github 查看最新release                                   |
| toolVersion   | simple admin tools 版本号 | 需要到github查看simple admin  tools 最新 release               |
| port          | 端口号                    | 服务暴露的端口号                                                |

> 你可以看到以下结构

![Example](../../assets/example-struct.png)


> 然后编辑 etc/example.yaml

```yaml
Name: example.api
Host: 0.0.0.0
Port: 8081
Timeout: 30000

Auth:
  AccessSecret: # the same as core
  AccessExpire: 259200

Log:
  ServiceName: exampleApiLogger
  Mode: file
  Path: /home/ryan/data/logs/example/api
  Level: info
  Compress: false
  KeepDays: 7
  StackCoolDownMillis: 100

Prometheus:
  Host: 0.0.0.0
  Port: 4000
  Path: /metrics


RedisConf:
  Host: 127.0.0.1:6379
  Type: node

DatabaseConf:
  Type: mysql
  Host: 127.0.0.1
  Port: 3306
  DBName: simple_admin
  Username: root # set your username
  Password: "123456" # set your password
  MaxOpenConn: 100
  SSLMode: false
  CacheTime: 5

CasbinConf:
  ModelText: |
    [request_definition]
    r = sub, obj, act
    [policy_definition]
    p = sub, obj, act
    [role_definition]
    g = _, _
    [policy_effect]
    e = some(where (p.eft == allow))
    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act

ExampleRpc:
  Endpoints:
    - 127.0.0.1:8080
```

> 运行代码

```shell
go run example.go -f etc/example.yaml
```

> 如果看到

```shell
Starting server at 127.0.0.1:8081...
```

说明运行成功.

## 代码生成（基于Proto）

```shell
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --serviceName=example --o=./ --model=Student --rpcName=Example --grpcPackage=github.com/suyuan32/simple-admin-example-rpc/example
```
| 参数          | 介绍                | 使用方法                                                           |
|-------------|-------------------|----------------------------------------------------------------|
| proto       | proto文件地址         | 输入proto文件的绝对路径                                                 |
| style       | 文件名格式             | go_zero为蛇形格式                                                   |
| serviceName | 服务名称              | 和new 时的名称相同，如example.go的serviceName是 example                   |
| o           | 输出位置              | 文件输出位置，可以为相对路径，指向main文件目录                                      |
| model       | 模型名称              | schema中内部struct名称，如example中的Student                            |
| rpcName     | RPC名称             | 输入Example则生成文件会生成l.svcCtx.ExampleRpc                           |
| grpcPackage | RPC *_grpc.go 包路径 | 在example中是github.com/suyuan32/simple-admin-example-rpc/example |

生成效果

![pic](../../assets/api_gen_struct.png)

> 详情查看 simple admin example api 地址 https://github.com/suyuan32/simple-admin-example-api

注意还需要手动添加下 service_context, config, etc, ExampleRpc