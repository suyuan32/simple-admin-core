# 环境配置

## 环境需求
- golang 1.19
- nodejs 18.8.0
- etcd
- mysql 

## 后端部署

### simple admin core
simple admin core 是核心代码，必须运行

### 下载代码 
```bash
git clone https://github.com/suyuan32/simple-admin-core.git
```

### 配置api

api/etc/core.yaml

```yaml
Name: core.api  # 服务名称
Host: 0.0.0.0  # 监听地址
Port: 8500  # 端口号
Auth:
  AccessSecret:    # jwt 加密秘钥， 推荐使用随机字符串
  AccessExpire:   #  过期时间（秒）
Log:
  ServiceName: coreApiLogger # logger名
  Mode: file  # 日志类型 file为存储为文件
  Path: /data/logs/api  # 文件存储路径
  Level: info  # 日志级别 info/error
  Compress: false # 压缩
  KeepDays: 7 # 保存时间
  StackCooldownMillis: 100
RedisConf:
  Host: 127.0.0.1:6379  # redis 地址
  Type: node       # 节点类型
  Pass: xxx    # 密码
CoreRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379  # etcd 地址
    Key: core.rpc  # rpc服务名称，要与rpc设置的名称一一致
Captcha:
  KeyLong: 5   # 验证码长度
  ImgWidth: 240  # 验证码图片宽度
  ImgHeight: 80  # 验证码图片高度
DatabaseConf:
  Type: mysql  # 服务器类型 mysql 或 postgresql
  Path: 127.0.0.1  # 地址
  Port: 3306  
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin # 数据库名称
  Username:    # 用户名
  Password:    # 密码
  MaxIdleConn: 10  # 最大空闲连接数
  MaxOpenConn: 100  # 最大连接数
  LogMode: error  # 日志级别
  LogZap: false   # 使用 zap 日志
```

### 配置rpc

rpc/etc/core.yaml

```yaml
Name: core.rpc    # 服务名称
ListenOn: 127.0.0.1:8501   # 监听地址
Etcd:
  Hosts:
    - 127.0.0.1:2379   # etcd地址
  Key: core.rpc        # 服务名称，用于服务发现
  User: root           # 用户名
  Pass: xxx            # 密码
DatabaseConf:
  Type: mysql  # 服务器类型 mysql 或 postgresql
  Path: 127.0.0.1  # 地址
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin # 数据库名称
  Username:    # 用户名
  Password:    # 密码
  MaxIdleConn: 10  # 最大空闲连接数
  MaxOpenConn: 100  # 最大连接数
  LogMode: error  # 日志级别
  LogZap: false   # 使用 zap 日志
Log:
  ServiceName: coreRpcLogger # logger名
  Mode: file  # 日志类型 file为存储为文件
  Path: /data/logs/rpc  # 文件存储路径
  Level: info  # 日志级别 info/error
  Compress: false # 压缩
  KeepDays: 7 # 保存时间
  StackCooldownMillis: 100
RedisConf:
  Host: 127.0.0.1:6379  # redis 地址
  Type: node       # 节点类型
  Pass: xxx    # 密码
```

### 配置依赖

```shell 
go mod tidy
```


### 运行 rpc 服务

```bash
cd rpc

go run core.go -f etc/core.yaml
```


### 运行 api 服务

```bash
cd api

go run core.go -f etc/core.yaml
```

## 前端配置

### 下载代码

```shell
git clone https://github.com/suyuan32/simple-admin-backend-ui.git
```

### 下载依赖

```shell
yarn install
```

### 运行

```shell
yarn serve
```