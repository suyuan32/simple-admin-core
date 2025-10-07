VERSION=$(shell git describe --tags --always)

.PHONY: docker
docker: # Compile and build the docker | 编译并构建 docker 镜像
	pnpm install
	pnpm build:simple-admin-core
	docker build -f Dockerfile -t ${DOCKER_USERNAME}/backend-ui-vben5:${VERSION} .

.PHONY: docker-not-build
docker-not-build: # Build the docker without compiling | 不编译直接构建镜像
	docker build -f Dockerfile -t ${DOCKER_USERNAME}/backend-ui-vben5:${VERSION} .

.PHONY: publish-docker
publish-docker: # Publish the docker | 发布镜像
	docker push ${DOCKER_USERNAME}/backend-ui-vben5:${VERSION}

.PHONY: run-docker
run-docker: # Run the docker image | 运行 docker 镜像
	docker volume create backendui-vben5
	docker run -d --name ${DOCKER_USERNAME}/backend-ui:${VERSION} -p 80:80 -v backendui-vben5:/etc/nginx --network docker-compose_simple-admin ${DOCKER_USERNAME}/backend-ui-vben5:${VERSION}

.PHONY: help
help: # Show help | 显示帮助
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
