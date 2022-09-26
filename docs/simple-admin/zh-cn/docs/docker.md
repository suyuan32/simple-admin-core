# Docker 镜像

## 生成 API 和 RPC 镜像
修改 Makfile 将版本换成自己的
```makefile
version := 0.0.1
docker:
	sudo docker build -f Dockerfile-api -t simpleadminapi:$(version) .
	sudo docker build -f Dockerfile-rpc -t simpleadminrpc:$(version) .
```

然后执行 make docker 即可生成镜像

### 注意生成的中间镜像可以删除，占用1.9G的空间太大。

## 本地部署 mysql, consul, 和 redis

进入 deploy/docker-compose 文件夹执行
```shell
docker-compose up -d
```

#### 默认数据库账号 root
#### 密码： 123456


### consul 和 redis 没有token和密码.