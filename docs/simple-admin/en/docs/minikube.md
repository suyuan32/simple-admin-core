# Minikube initialization

### Install

[Minikube Install](https://minikube.sigs.k8s.io/docs/start/)


### Start minikube
> Recommend to set the memory limit in local development，my setting is 3gb

```shell
minikube start --memory 3g
```

### minikube add plugins
> Enable ingress adn metrics (metrics is used in go-zero for auto-scaling monitor)

```shell
minikube addons enable metrics-server
minikube addons enable ingress
```

### Add alias
```shell
alias kubectl="minikube kubectl --"
```

### Deploy the service
> clone the code，cd deploy/k8s

```shell
# add service discovery account
kubectl apply -f auth.yaml

# modify core-rpc.yaml to set your own image
kubectl apply -f core-rpc.yaml

# modify core-api.yaml to set your own image
kubectl apply -f core-api.yaml

# modify backend-ui.yaml to set your own image
kubectl apply -f backend-ui.yaml
```

> You can just run simple-admin-core/deploy/k8s/setup.sh to finish the job.

### Ingress Setting
Reference [Official Document](https://minikube.sigs.k8s.io/docs/tutorials/nginx_tcp_udp_ingress/)

### Firstly, add ingress

```shell
kubectl apply -f ingress.yaml
```
> Ingress.yaml file

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: simple-admin-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$1

spec:
  rules:
    - host: simple-admin.com # your domain
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: backendui-svc
                port:
                  number: 80
    - host: simple-admin.com # your domain
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

### Register tcp service into ingress configmap 

```shell
# register coreapi
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"9100":"simple-admin/coreapi-svc:9100"}}'
# register backendui, connect 8080 port to 80 in backendui service
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"8080":"simple-admin/backendui-svc:80"}}'
```

### Register service into  ingress controller
```shell
kubectl patch deployment ingress-nginx-controller --patch "$(cat ingress-patch.yaml)" -n ingress-nginx

```

> You can just run simple-admin-core/deploy/k8s/setup-ingress.sh to finish the job.

> browse ingress IP address 

```shell
kubectl get ingress
```
> You can see

```shell
NAME                   CLASS   HOSTS                               ADDRESS        PORTS   AGE
simple-admin-ingress   nginx   simple-admin.com,simple-admin.com   192.168.49.2   80      2m8s
```

### Modify local hosts
> ip is the same as ingress 

```shell
# add dns
192.168.49.2    simple-admin.com
```

### Visit the site

```shell
http://simple-admin.com:8080/
```

> You can see the web page.

