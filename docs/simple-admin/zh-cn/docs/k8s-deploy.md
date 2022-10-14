# 部署K8s

## 环境需求
- minikube v1.23.0 +
- mysql 5.7+
- redis 6.0 +
- docker

## Minikube 配置
[Minikube](simple-admin/zh-cn/docs/minikube.md)

### K8s配置
#### API 服务
> api/etc/core.yaml
```yaml
Name: core.api
Host: 0.0.0.0 # 需要 0.0.0.0 以便外部访问
Port: 9100
Timeout: 30000

Auth:
  AccessSecret: jS6VKDtsJf3z1n2VKDtsJf3z1n2  # JWT的加密密钥，各个API应保持一致才能解析
  AccessExpire: 259200  # 秒，过期时间

Log:
  ServiceName: coreApiLogger
  Mode: file # 日志模式
  Path: /home/ryan/logs/core/api  # log 保存路径，使用filebeat同步
  Level: info # 日志等级
  Compress: false # 日志压缩
  KeepDays: 7  # 保存时长（天）
  StackCoolDownMillis: 100 # 多少毫秒后再次写入堆栈跟踪。用来避免堆栈跟踪日志过多

RedisConf:
  Host: 127.0.0.1:6379  # 改成自己的redis地址
  Type: node
  # Pass: xxx  # 也可以设置密码

CoreRpc:
  Target: k8s://simple-admin/corerpc-svc:9101 # 格式 k8s://namespace/service-name:port

Captcha:
  KeyLong: 5 # 验证码长度
  ImgWidth: 240 # 验证码图片宽度
  ImgHeight: 80 # 验证码图片高度

DatabaseConf:
  Type: mysql
  Path: "127.0.0.1"  # 修改成自己的mysql地址
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local # gorm时间转换需要如下配置
  DBName: simple_admin # 数据库名，可以自定义
  Username: root   # 用户名
  Password: "123456" # 密码
  MaxIdleConn: 10  # 最大空闲连接
  MaxOpenConn: 100 # 最大连接数
  LogMode: error # log 级别
  LogZap: false # 目前log zap还未实现

# 服务监控
Prometheus:
  Host: 0.0.0.0
  Port: 4000
  Path: /metrics
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

# 服务监控
Prometheus:
  Host: 0.0.0.0
  Port: 4001
  Path: /metrics
```

### Docker镜像编译发布

> 手动编译

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

> 建议使用 gitlab-ci， 项目已默认提供， 需要在 gitlab runner 设置 variable ： $DOCKER_USERNAME 和 $DOCKER_PASSWORD

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

## 部署流程
- 生成docker镜像
- 上传docker仓库
- 在k8s集群使用命令 kubectl apply -f deploy/k8s/coreapi.yaml 等直接部署
> 生成镜像和上传仓库建议直接使用gitlab-ci自动发布

### coreapi 部署文件详解

> coreapi 是所有服务的label和metadata:name \
> 命名空间默认是 simple-admin, 可自行修改

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: core-api
  labels:
    app: core-api
spec:
  replicas: 3
  revisionHistoryLimit: 5
  selector:
    matchLabels:
      app: core-api
  template:
    metadata:
      labels:
        app: core-api
    spec:
      serviceAccountName: endpoints-finder
      containers:
      - name: core-api
        image: ryanpower/core-api:0.0.19 # 主要修改此处镜像
        ports:
        - containerPort: 9100 # 端口， 与 core.yaml 内配置端口相同
        readinessProbe:
          tcpSocket:
            port: 9100
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: 9100
          initialDelaySeconds: 15
          periodSeconds: 20
        resources:
          requests:
            cpu: 100m  # 最低 cpu 需求， 1000m 为一个cpu,测试环境建议小一些
            memory: 100Mi # 本地调试我设置了 100 mb, 正式环境可以为 512Mi
          limits:
            cpu: 1000m # 最高占用 cpu
            memory: 1024Mi # 最高占用的内存
        volumeMounts:
        - name: timezone
          mountPath: /etc/localtime
      volumes:
        - name: timezone
          hostPath:
            path: /usr/share/zoneinfo/Asia/Shanghai # 设置时区

---

apiVersion: v1
kind: Service
metadata:
  name: core-api-svc
  labels:
    app: core-api-svc
spec:
  type: NodePort
  ports:
    - port: 9100
      targetPort: 9100
      name: api
      protocol: TCP
  selector:
    app: core-api

---

apiVersion: v1
kind: Service
metadata:
  name: core-api-svc
  labels:
    app: core-api-svc
spec:
  ports:
    - port: 4000
      name: prometheus
      targetPort: 4000
  selector:
    app: core-api


---
# 服务监控
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: core-rpc
  labels:
    serviceMonitor: prometheus
spec:
  selector:
    matchLabels:
      app: core-rpc-svc
  endpoints:
    - port: prometheus

---
# autoscaling 用于动态增加负载，通过 metric-server 获取使用率，目前获取 metric 还有些问题，近期会修复
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: core-api-hpa-c
  labels:
    app: core-api-hpa-c
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: core-api
  minReplicas: 3  # 最小副本
  maxReplicas: 10 # 最大副本
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 80 # 平均使用率 80%

---

apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: core-api-hpa-m
  labels:
    app: core-api-hpa-m
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: core-api
  minReplicas: 3
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: memory
        target:
          type: Utilization
          averageUtilization: 80

```

> core rpc 和 backend ui 相似

## 前端 nginx 请求设置

> simple-admin-backend-ui/deploy/default.conf

```text
server {
    listen       80;
    listen  [::]:80;
    server_name  localhost;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
        try_files $uri $uri/ /index.html;
    }

    location /sys-api/ {
        proxy_pass  http://core-api-svc.default.svc.cluster.local:9100/;
    }
    
    # location /file-manager/ {
    #     proxy_pass  http://file-api-svc.default.svc.cluster.local:9102/;
    # }
}
```

> 注意 proxy_pass 格式  http://{service-name}.{namespace}.svc.cluster.local:{port}/

#### 快速部署

> 执行 deploy/k8s/setup.sh