# 部署K8s

## 环境需求
- minikube v1.23.0 +
- mysql 5.7+
- redis 6.0 +
- docker

# minikube addons enable metrics-server

### K8s配置
#### API 服务
> api/etc/core.yaml
```yaml
Name: core.api
Host: 0.0.0.0
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

CoreRpc:
  Target: k8s://simple-admin/corerpc-svc:9101 # 格式 k8s://namespace/service name:port

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
ListenOn: 0.0.0.0:9101

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

### Docker镜像编译发布

#### 手动编译
```shell
# 设置环境变量
export VERSION=0.0.1  # 版本号
export DOCKER_USERNAME=xxx # docker仓库用户名
export DOCKER_PASSWORD=xxx # docker仓库密码
export REPO=docker.io  # docker仓库地址

# 生成镜像
make docker

# 发布镜像
make publish-docker
```

建议使用 gitlab-ci， 项目已默认提供

```text
variables:
  VERSION: 0.0.19
  REPO: docker.io

stages:
  - info
  - build
  - publish
  - clean

info-job:
  stage: info
  script:
    - echo "Start build version $VERSION"
    - export VERSION=$VERSION
    - export DOCKER_USERNAME=$DOCKER_USERNAME
    - export DOCKER_PASSWORD=$DOCKER_PASSWORD
    - export REPO=$REPO

build-job:
  stage: build
  script:
    - echo "Compiling the code and build docker image..."
    - make docker
    - echo "Compile complete."

deploy-job:
  stage: publish
  environment: production
  script:
    - echo "Publish docker images..."
    - make publish-docker
    - echo "Docker images successfully published."

clean-job:
  stage: clean
  script:
    # 删除所有 none 镜像 | delete all none images
    - docker images |grep none|awk '{print $3}'|xargs docker rmi
    - echo "Delete all none images successfully."
```

