# Deploy service into  K8s 

## Environment Requirement
- minikube v1.23.0 +
- mysql 5.7+
- redis 6.0 +
- docker

## Minikube Setting
[Minikube](simple-admin/en/docs/minikube.md)

### K8s Setting
#### API service
> api/etc/core.yaml
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

CoreRpc:
  Target: k8s://simple-admin/corerpc-svc:9101 # format  k8s://namespace/service-name:port

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

> rpc/etc/core.yaml
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

### Docker image publish

#### Mammal
```shell
# Set the env variables
export VERSION=0.0.1  # version
export DOCKER_USERNAME=xxx # docker repository username
export DOCKER_PASSWORD=xxx # docker repository password
export REPO=docker.io  # docker repository path

# build the image
make docker

# publish the image
make publish-docker
```

Recommend to use gitlab-ci. The project had been provided .gitlab-ci.yml， You need set variable ： $DOCKER_USERNAME 和 $DOCKER_PASSWORD in gitlab runner.

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

## Deploy pipeline 
- build docker image
- upload to docker repository 
- run in k8s ->  kubectl apply -f deploy/k8s/coreapi.yaml
> You can use gitlab-ci to automatically build and push docker image
### coreapi k8s deployment file tutorial
> coreapi the name of service, you can find in label and metadata:name
> Namespace is simple-admin by default, you can change your own namespace
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: coreapi
  namespace: simple-admin
  labels:
    app: coreapi
spec:
  replicas: 3
  revisionHistoryLimit: 5
  selector:
    matchLabels:
      app: coreapi
  template:
    metadata:
      labels:
        app: coreapi
    spec:
      serviceAccountName: endpoints-finder
      containers:
      - name: coreapi
        image: ryanpower/coreapi:0.0.19 # mainly change this image
        ports:
        - containerPort: 9100 # port， the same as core.yaml 
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
            cpu: 100m  # the lowest cpu requirement， 1000m is one cpu, set lower in development env
            memory: 100Mi # I set 100 mb memory for local test, in production you can set 512Mi
          limits:
            cpu: 1000m # the maximum cpu requirement， 1000m is one cpu, set lower in development env
            memory: 1024Mi # the maximum memory usage
        volumeMounts:
        - name: timezone
          mountPath: /etc/localtime
      volumes:
        - name: timezone
          hostPath:
            path: /usr/share/zoneinfo/Asia/Shanghai # set time zone

---

apiVersion: v1
kind: Service
metadata:
  name: coreapi-svc
  namespace: simple-admin
spec:
  type: NodePort
  ports:
  - port: 9100
    targetPort: 9100
    name: http
    protocol: TCP
  selector:
    app: coreapi

---
# autoscaling is used to auto-scaling the replicas， use metric-server to get usage info，but it has some bugs now, it will fix in the future.
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: coreapi-hpa-c
  namespace: simple-admin
  labels:
    app: coreapi-hpa-c
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: coreapi
  minReplicas: 3  # the min replicas number
  maxReplicas: 10 # the max replicas number
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 80 # average usage 80%

---

apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: coreapi-hpa-m
  namespace: simple-admin
  labels:
    app: coreapi-hpa-m
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: coreapi
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

> corerpc and backendui are similar.

## Frontend nginx setting

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
        proxy_pass  http://coreapi-svc.simple-admin.svc.cluster.local:9100/;
    }
    
    # location /file-manager/ {
    #     proxy_pass  http://fileapi-svc:9102/;
    # }
}
```

> Notice: proxy_pass format:   http://{service-name}.{namespace}.svc.cluster.local:{port}/