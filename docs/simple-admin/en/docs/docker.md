# Docker Image

## Build API and RPC image
Modify Makfile and change the version
```makefile
version := 0.0.1
docker:
	sudo docker build -f Dockerfile-api -t simpleadminapi:$(version) .
	sudo docker build -f Dockerfile-rpc -t simpleadminrpc:$(version) .
```

Run **make docker**  to build the images.

### The middle images can be removed after build because it cost too many space(about 1.9G).

## Deploy mysql, consul and redis locally

cd deploy/docker-compose and run
```shell
docker-compose up -d
```

#### default mysql account: root
#### password： 123456

### consul and redis do not have password adn token.