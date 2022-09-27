# 项目部署 docker
接下来讲述如何将项目部署在本地docker, k8s 同理
## 首先使用 deploy-compose 初始化数据库，redis， consul
 
目录是 deploy/docker-compose

```dockerfile
version: '3'

volumes:
  mysql:
  redis:
  consul:

networks:
  simple-admin:
    driver: bridge

services:
  mysql:
    image: mysql:8.0.21
    container_name: mysql
    command: mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
    restart: always
    ports:
      - '3306:3306'
    environment:
      MYSQL_DATABASE: 'simple-admin'
      MYSQL_ROOT_PASSWORD: '123456'
    volumes:
      - mysql:/var/lib/mysql
    networks:
      simple-admin:
        aliases:
          - mysqlserver

  redis:
    image: redis:7.0.5-alpine
    container_name: redis
    restart: always
    ports:
      - '6379:6379'
    volumes:
      - redis:/data
    networks:
      simple-admin:
        aliases:
          - redisserver

  consul:
    image: docker.io/bitnami/consul:latest
    container_name: consul
    volumes:
      - consul:/bitnami/consul
    ports:
      - '8300:8300'
      - '8301:8301'
      - '8301:8301/udp'
      - '8500:8500'
      - '8600:8600'
      - '8600:8600/udp'
    networks:
      simple-admin:
        aliases:
          - consulserver

```

然后运行

```shell
docker-compose up -d
```

默认创建了 docker-compose_simple-admin 网络， 后续 corerpc 和 coreapi 也会添加到这个网络

## 修改 etc/core.yaml

### API
```yaml
Consul:
  Host: consulserver:8500 # consul endpoint
  ListenOn: coreapi:9100  # this is used for other services to find this service
  #Token: 'f0512db6-76d6-f25e-f344-a98cc3484d42' # consul ACL token (optional)
  Key: core.api
  Meta:
    Protocol: grpc
  Tag:
    - core
    - api
```

### RPC 

```yaml
Consul:
  Host: consulserver:8500 # consul endpoint
  ListenOn: corerpc:9101
  #Token: 'f0512db6-76d6-f25e-f344-a98cc3484d42' # consul ACL token (optional)
  Key: core.rpc
  Meta:
    Protocol: grpc
  Tag:
    - core
    - rpc
```

#### consulserver 和 docker compose 中一致
#### 注意 etc 内的配置中的  ListenOn 是用于 api获取对应地址的，而 consul 中虽然也是ListenOn， 但是它是容器内的监听地址，所以是 0.0.0.0

## Consul 内设置修改

### coreApiConf

```yaml
Name: core.api
Host: 0.0.0.0
Port: 9100
Timeout: 30000
Auth:
  AccessSecret: jS6VKDtsJf3z1n2VKDtsJf3z1n2
  AccessExpire: 259200  # Seconds
Log:
  ServiceName: coreApiLogger
  Mode: file
  Path: /home/ryan/logs/core/api
  Level: info
  Compress: false
  KeepDays: 7
  StackCooldownMillis: 100
RedisConf:
  Host: redisserver:6379  # change to your
  Type: node
CoreRpc:
  Target: consul://consulserver:8500/core.rpc?wait=14s
Captcha:
  KeyLong: 5
  ImgWidth: 240
  ImgHeight: 80
DatabaseConf:
  Type: mysql
  Path: mysqlserver
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin
  Username: root
  Password: "123456"
  MaxIdleConn: 10
  MaxOpenConn: 100
  LogMode: error
  LogZap: false
```


### coreRpcConf

```yaml
Name: core.rpc
ListenOn: 0.0.0.0:9101
DatabaseConf:
  Type: mysql
  Path: mysqlserver
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin
  Username: root
  Password: "123456"
  MaxIdleConn: 10
  MaxOpenConn: 100
  LogMode: error
  LogZap: false
Log:
  ServiceName: coreRpcLogger
  Mode: file
  Path: /home/ryan/logs/core/rpc
  Level: info
  Compress: false
  KeepDays: 7
  StackCooldownMillis: 100
RedisConf:
  Host: redisserver:6379
  Type: node
```

## 编译 docker image
修改自己的版本
```makefile
version := 0.0.8
docker:
	sudo docker build -f Dockerfile-api -t coreapi:$(version) .
	sudo docker build -f Dockerfile-rpc -t corerpc:$(version) .

run-docker:
	sudo docker run -d --name corerpc-$(version) --network docker-compose_simple-admin --network-alias corerpc -p 9101:9101 corerpc:$(version)
	sudo docker run -d --name coreapi-$(version) --network docker-compose_simple-admin --network-alias coreapi -p 9100:9100 coreapi:$(version)

run-docker-rpc:
	sudo docker run -d --name corerpc-$(version) --network docker-compose_simple-admin --network-alias corerpc -p 9101:9101 corerpc:$(version)

run-docker-api:
	sudo docker run -d --name coreapi-$(version) --network docker-compose_simple-admin --network-alias coreapi -p 9100:9100 coreapi:$(version)

```

到项目根目录执行

```shell
make docker 
```

然后执行

```shell
make run-docker
```

即可部署好 API 和 RPC

# Simple Admin UI 部署

## 注意： VITE_PROXY 在生产环境需要 Nginx 做转发。

## 编译 docker image

进入 simple-amind-backend-ui/

### 修改 makefile
```makefile
version := 0.0.6
user := ryan

docker:
	yarn install
	yarn build
	sudo docker build -f Dockerfile -t backendui:$(version) .

run-docker:
	sudo docker volume create backendui
	sudo docker run -d --name backendui-$(version) -p 80:80 -v backendui:/etc/nginx --network docker-compose_simple-admin backendui:$(version)
```
修改 version 和 user.

### 运行
```shell
make docker

make run-docker 
```

进入容器修改 /etc/nginx/conf.d/default.conf

```text
server {
    listen       80;
    listen  [::]:80;
    server_name  localhost;

    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    location /sys-api/ {
        proxy_pass  http://coreapi:9100/;
    }
    
    location /file-manager/ {
        proxy_pass  http://coreapi:9102/;
    }

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}

    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
    #
    #location ~ \.php$ {
    #    root           html;
    #    fastcgi_pass   127.0.0.1:9000;
    #    fastcgi_index  index.php;
    #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
    #    include        fastcgi_params;
    #}

    # deny access to .htaccess files, if Apache's document root
    # concurs with nginx's one
    #
    #location ~ /\.ht {
    #    deny  all;
    #}
}
```

执行

```shell
nginx -s reload
```

即可实现访问

```shell
http://localhost/
```

