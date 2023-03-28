PROJECT=core
GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w

.PHONY: test
test: # Run test for the project | 运行项目测试
	go test -v --cover ./internal/..

.PHONY: fmt
fmt: # Format the codes | 格式化代码
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint: # Run go linter | 运行代码错误分析
	golangci-lint run -D staticcheck

.PHONY: tools
tools: # Install the necessary tools | 安装必要的工具
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest;

.PHONY: docker
docker: # Build the docker image | 构建 docker 镜像
	docker build -f Dockerfile-api -t ${DOCKER_USERNAME}/$(PROJECT)-api:${VERSION} .
	docker build -f Dockerfile-rpc -t ${DOCKER_USERNAME}/$(PROJECT)-rpc:${VERSION} .
	@echo "Build docker successfully"

.PHONY: publish-docker
publish-docker: # Publish docker image | 发布 docker 镜像
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin https://${REPO}
	docker push ${DOCKER_USERNAME}/$(PROJECT)-rpc:${VERSION}
	docker push ${DOCKER_USERNAME}/$(PROJECT)-api:${VERSION}
	@echo "Publish docker successfully"

.PHONY: gen-api
gen-api: # Generate API files | 生成 API 的代码
	goctls api go --api ./api/desc/all.api --dir ./api --trans_err=true
	swagger generate spec --output=./$(PROJECT).yml --scan-models
	@echo "Generate API files successfully"

.PHONY: gen-rpc
gen-rpc: # Generate RPC files from proto | 生成 RPC 的代码
	goctls rpc protoc ./rpc/$(PROJECT).proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
	@echo "Generate RPC files successfully"

.PHONY: gen-ent
gen-ent: # Generate Ent codes | 生成 Ent 的代码
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./rpc/ent/template/*.tmpl" ./rpc/ent/schema
	@echo "Generate Ent files successfully"

.PHONY: gen-rpc-ent-logic
gen-rpc-ent-logic: # Generate logic code from Ent, need model and group params | 根据 Ent 生成逻辑代码, 需要设置 model 和 group
	goctls rpc ent --schema=./rpc/ent/schema --service_name=$(PROJECT) --project_name=$(PROJECT) --o=./rpc --model=$(model) --group=$(group) --proto_out=./rpc/desc/$(shell echo $(model) | tr A-Z a-z).proto
	@echo "Generate logic codes from Ent successfully"

.PHONY: build-win-rpc
build-win-rpc: # Build RPC project for Windows | 构建Windows下的RPC可执行文件
	env CGO_ENABLED=0 GOOS=windows go build -ldflags "$(LDFLAGS)" -o $(PROJECT)-rpc.exe ./rpc/$(PROJECT).go
	@echo "Build RPC project for Windows successfully"

.PHONY: build-mac-rpc
build-mac-rpc: # Build RPC project for MacOS | 构建MacOS下的RPC可执行文件
	env CGO_ENABLED=0 GOOS=darwin go build -ldflags "$(LDFLAGS)" -o $(PROJECT)-rpc ./rpc/$(PROJECT).go
	@echo "Build RPC project for MacOS successfully"

.PHONY: build-linux-rpc
build-linux-rpc: # Build RPC project for Linux | 构建Linux下的RPC可执行文件
	env CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(PROJECT)-rpc ./rpc/$(PROJECT).go
	@echo "Build RPC project for Linux successfully"

.PHONY: build-win-api
build-win-api: # Build API project for Windows | 构建Windows下的API可执行文件
	env CGO_ENABLED=0 GOOS=windows go build -ldflags "$(LDFLAGS)" -o $(PROJECT)-api.exe ./api/$(PROJECT).go
	@echo "Build API project for windows successfully"

.PHONY: build-mac-api
build-mac-api: # Build API project for MacOS | 构建MacOS下的API可执行文件
	env CGO_ENABLED=0 GOOS=darwin go build -ldflags "$(LDFLAGS)" -o $(PROJECT)-api ./api/$(PROJECT).go
	@echo "Build API project for MacOS successfully"

.PHONY: build-linux-api
build-linux-api: # Build API project for Linux | 构建Linux下的API可执行文件
	env CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(PROJECT)-api ./api/$(PROJECT).go
	@echo "Build API project for Linux successfully"

.PHONY: gen-swagger
gen-swagger: # Generate swagger file | 生成 swagger 文件
	swagger generate spec --output=./$(PROJECT).yml --scan-models
	@echo "Generate swagger successfully"

.PHONY: serve-swagger
serve-swagger: # Run the swagger server | 运行 swagger 服务
	lsof -i:36666 | awk 'NR!=1 {print $2}' | xargs killall -9 || true
	swagger serve -F=swagger --port 36666 $(PROJECT).yml
	@echo "Serve swagger-ui successfully"

.PHONY: help
help: # Show help | 显示帮助
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
