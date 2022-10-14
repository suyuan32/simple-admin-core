# Minikube 初始化环境配置

## Minikube 初始化

### 安装

[Minikube安装](https://minikube.sigs.k8s.io/docs/start/)


### 启动 minikube
> 建议限制内存使用，本地运行，我设置的是 3g

```shell
minikube start --memory 3g
```

### minikube 启用插件
> 启用 ingress 和 metrics (metrics是gozero默认使用的,用于自动扩展)
```shell
minikube addons enable metrics-server
minikube addons enable ingress
```

### 添加别名
```shell
alias kubectl="minikube kubectl --"
```

### 部署服务
> clone 代码，进入 deploy/k8s 文件夹
```shell
# 创建 namespace
kubectl create namespace simple-admin

# 添加服务注册发现账号
kubectl apply -f auth.yaml

# 修改 core-rpc.yaml 中的镜像为自己的，然后执行
kubectl apply -f core-rpc.yaml

# 修改 core-api.yaml 中的镜像为自己的，然后执行
kubectl apply -f core-api.yaml

# 修改 backend-ui.yaml 中的镜像为自己的，然后执行
kubectl apply -f backend-ui.yaml
```

### Ingress 配置
参考官方文档 [Official Document](https://minikube.sigs.k8s.io/docs/tutorials/nginx_tcp_udp_ingress/)

### 首先创建 ingress

```shell
kubectl apply -f ingress.yaml
```
> Ingress 文件
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: simple-admin-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$1

spec:
  rules:
    - host: simple-admin.com # 域名
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: backendui-svc
                port:
                  number: 80
    - host: simple-admin.com # 域名
      http:
        paths:
          - path: /sys-api/
            pathType: Prefix
            backend:
              service:
                name: coreapi-svc
                port:
                  number: 9100
```

### 注册 tcp 服务到 ingress configmap 中

```shell
# 注册 coreapi
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"9100":"simple-admin/coreapi-svc:9100"}}'
# 注册 backendui, 将 8080 端口映射到backendui的80端口
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"8080":"simple-admin/backendui-svc:80"}}'
```

### 注册服务到 ingress controller中
```shell
kubectl patch deployment ingress-nginx-controller --patch "$(cat ingress-patch.yaml)" -n ingress-nginx

```

> 查看 ingress 地址
```shell
kubectl get ingress
```
> 返回
```shell
NAME                   CLASS   HOSTS                               ADDRESS        PORTS   AGE
simple-admin-ingress   nginx   simple-admin.com,simple-admin.com   192.168.49.2   80      2m8s
```

### 修改本地 hosts
> ip 和 ingress 中一致
```shell
# 添加dns
192.168.49.2    simple-admin.com
```

### 本地访问

```shell
http://simple-admin.com:8080/
```
即可看到效果

