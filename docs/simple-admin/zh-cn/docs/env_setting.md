# 环境配置

## 环境需求
- golang 1.19
- nodejs 18.8.0
- mysql 5.7 +
- redis 6.0 +
- etcd

## 后端部署

### simple admin core
simple admin core 是核心代码，主要负责用户注册鉴权和充当网关的角色以及后台的各类配置。

#### 默认账号
username:     admin  \
password:     simple-admin

### 下载代码 
```bash
git clone https://github.com/suyuan32/simple-admin-core.git
```

### 本地调试配置
#### API 服务
#### 注意本地测试最好创建 core_dev.yaml 与部署文件core.yaml区分开
> api/etc/core_dev.yaml
```yaml
Name: core.api
Host: 0.0.0.0 # ip可以是0.0.0.0也可以是127.0.0.1
Port: 9100
Timeout: 30000

Auth:
  AccessSecret: jS6VKDtsJf3z1n2VKDtsJf3z1n2  # JWT的加密密钥，各个API应保持一致才能解析
  AccessExpire: 259200  # 秒，过期时间

Log:
  ServiceName: coreApiLogger
  Mode: file
  Path: /home/ryan/logs/core/api  # log 保存路径，使用filebeat同步
  Level: info
  Compress: false
  KeepDays: 7  # 保存时长（天）
  StackCoolDownMillis: 100

RedisConf:
  Host: 127.0.0.1:6379  # 改成自己的redis地址
  Type: node
  # Pass: xxx  # 也可以设置密码

# 与 k8s 主要是此处服务发现不同
CoreRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: core.rpc

Captcha:
  KeyLong: 5
  ImgWidth: 240
  ImgHeight: 80

DatabaseConf:
  Type: mysql
  Path: "127.0.0.1"  # 修改成自己的mysql地址
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin
  Username: root   # 用户名
  Password: "123456" # 密码
  MaxIdleConn: 10  # 最大空闲连接
  MaxOpenConn: 100 # 最大连接数
  LogMode: error
  LogZap: false

```

> rpc/etc/core.yaml
```yaml
Name: core.rpc
ListenOn: 0.0.0.0:9101  # ip可以是0.0.0.0也可以是127.0.0.1
# 需要添加 ETCD 配置
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: user.rpc

DatabaseConf:
  Type: mysql
  Path: "127.0.0.1"  # 修改成自己的mysql地址
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin
  Username: root   # 用户名
  Password: "123456" # 密码
  MaxIdleConn: 10  # 最大空闲连接
  MaxOpenConn: 100 # 最大连接数
  LogMode: error
  LogZap: false

Log:
  ServiceName: coreRpcLogger
  Mode: file
  Path: /home/ryan/logs/core/rpc  # log 保存路径，使用filebeat同步
  Encoding: json
  Level: info
  Compress: false
  KeepDays: 7  # 保存时长（天）
  StackCoolDownMillis: 100

RedisConf:
  Host: 192.168.50.216:6379   # 改成自己的redis地址
  Type: node
  # Pass: xxx  # 也可以设置密码
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

### 编译
```shell
yarn build
```

### 预览
```shell
# build and preview
yarn preview

# preview exist files
yarn preview:dist
```

## 初始化数据库
***重要:*** 在初始化数据库前必须先创建数据库, 数据库名称和配置文件中的名称相同.

```shell
# 访问前端地址端口
https://address:port/init

```
进入界面

![pic](../../assets/init_zh_cn.png)

## **初始化完成后需要重启 api 和 rpc。**