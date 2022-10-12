# Local Development Setting

> Environment Requirement
- golang 1.19
- nodejs 18.8.0
- mysql 5.7 +
- redis 6.0 +

> Backend Setting

## simple admin core
simple admin core is the core service for simple admin. It offers user management, authorization,
menu management and API management and so on. It must be running.

> Default Account

username:     admin  \
password:     simple-admin

> Clone the core code 
```bash
git clone https://github.com/suyuan32/simple-admin-core.git
```

> Local development setting
#### API Service
##### Notice: You should add core_dev.yaml for development to avoid conflicting with core.yaml in production.
> Add api/etc/core_dev.yaml
```yaml
Name: core.api
Host: 0.0.0.0 # ip can be 0.0.0.0 or 127.0.0.1, it should be 0.0.0.0 if you want to access from another host
Port: 9100
Timeout: 30000

Auth:
  AccessSecret: jS6VKDtsJf3z1n2VKDtsJf3z1n2  # JWT encrypt code
  AccessExpire: 259200  # seconds, expire duration

Log:
  ServiceName: coreApiLogger
  Mode: file
  Path: /home/ryan/logs/core/api  # log saving path，use filebeat to sync
  Level: info
  Compress: false
  KeepDays: 7  # Save time (Day)
  StackCoolDownMillis: 100

RedisConf:
  Host: 127.0.0.1:6379  # Change to your own redis address
  Type: node
  # Pass: xxx  # You can also set the password 

# The main difference between k8s and local development
CoreRpc:
  Endpoints:
    - 127.0.0.1:9101 # the same as rpc address

Captcha:
  KeyLong: 5 # captcha key length
  ImgWidth: 240 # captcha image width
  ImgHeight: 80 # captcha image height

DatabaseConf:
  Type: mysql
  Path: "127.0.0.1"  # change to your own mysql address
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local # set the config for time convert in gorm
  DBName: simple_admin # database name, you can set your own name
  Username: root   # username 
  Password: "123456" # password
  MaxIdleConn: 10  # the maximum number of connections in the idle connection pool
  MaxOpenConn: 100 # the maximum number of open connections to the database
  LogMode: error
  LogZap: false

```

> Small website use endpoint connect directly
>
> CoreRpc:
>  Endpoints:
>   - 127.0.0.1:9101
>
> it does not need service discovery，when you develop locally, you should also use this mode. There can be several endpoints.

> Add rpc/etc/core_dev.yaml
```yaml
Name: core.rpc
ListenOn: 0.0.0.0:9101  # ip can be 0.0.0.0 or 127.0.0.1, it should be 0.0.0.0 if you want to access from another host

DatabaseConf:
  Type: mysql
  Path: "127.0.0.1"  # change to your own mysql address
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local # set the config for time convert in gorm
  DBName: simple_admin # database name, you can set your own name
  Username: root   # username 
  Password: "123456" # password
  MaxIdleConn: 10  # the maximum number of connections in the idle connection pool
  MaxOpenConn: 100 # the maximum number of open connections to the database
  LogMode: error
  LogZap: false

Log:
  ServiceName: coreApiLogger
  Mode: file
  Path: /home/ryan/logs/core/api  # log saving path，use filebeat to sync
  Level: info
  Compress: false
  KeepDays: 7  # Save time (Day)
  StackCoolDownMillis: 100


RedisConf:
  Host: 127.0.0.1:6379  # Change to your own redis address
  Type: node
  # Pass: xxx  # You can also set the password 
```

> Sync dependencies

```shell 
go mod tidy
```

> Run rpc service

```bash
cd rpc

go run core.go -f etc/core_dev.yaml
```

> Run api service

```bash
cd api

go run core.go -f etc/core_dev.yaml
```

> Front end setting
>
> Clone the code

```shell
git clone https://github.com/suyuan32/simple-admin-backend-ui.git
```

> Sync dependencies

```shell
yarn install
```

> Run in development mode

```shell
yarn serve
```

> Build

```shell
yarn build
```

> Preview

```shell
# build and preview
yarn preview

# preview existing directory
yarn preview:dist
```

> Notice: Set the API address

> .env.development
```text
# Whether to open mock
VITE_USE_MOCK = false

# public path
VITE_PUBLIC_PATH = /

# Cross-domain proxy, you can configure multiple
# Please note that no line breaks
VITE_PROXY = [["/sys-api","http://localhost:9100"],["/file-manager","http://localhost:9102"]]

VITE_BUILD_COMPRESS = 'none'

# Delete console
VITE_DROP_CONSOLE = false

# Basic interface address SPA
VITE_GLOB_API_URL=

# File upload address， optional
VITE_GLOB_UPLOAD_URL=/upload

# File store address
VITE_FILE_STORE_URL=http://localhost:8080

# Interface prefix
VITE_GLOB_API_URL_PREFIX=
```
> Mainly modify sys-api in VITE_PROXY， use  localhost or 127.0.0.1 to connect local service，\ 
> you can also set your own address, file-manager is the API for upload it is optional

## Initialize database
***Important:*** You should create the database before initialization, the database name should be the same as core_dev.yaml.

```shell
# visit the address
https://address:port/init

# default is
https://localhost:3100/init
```

> You can see

![pic](../../assets/init_zh_cn.png)

> File manager service is optional [File Manager](/simple-admin/en/docs/file_manager.md)
## **After initialization, you should restart api and rpc service.**