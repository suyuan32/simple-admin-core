# Custom configuration | 独立配置
# Service name | 项目名称
SERVICE=Core
# Service name in specific style | 项目经过style格式化的名称
SERVICE_STYLE=core
# Service name in lowercase | 项目名称全小写格式
SERVICE_LOWER=core
# Service name in snake format | 项目名称下划线格式
SERVICE_SNAKE=core
# Service name in snake format | 项目名称短杠格式
SERVICE_DASH=core

# The project version, if you don't use git, you should set it manually | 项目版本，如果不使用git请手动设置
VERSION=$(shell git describe --tags --always)

# The project file name style | 项目文件命名风格
PROJECT_STYLE=go_zero

# Whether to use i18n | 是否启用 i18n
PROJECT_I18N=true

# Swagger type, support yml,json | Swagger 文件类型，支持yml,json
SWAGGER_TYPE=json

# Ent enabled features | Ent 启用的官方特性
ENT_FEATURE=sql/execquery,intercept,sql/modifier

# The arch of the build | 构建的架构
GOARCH=amd64

# ---- You may not need to modify the codes below | 下面的代码大概率不需要更改 ----

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w

.PHONY: test
test: # Run test for the project | 运行项目测试
	go test -v --cover ./api/internal/..
	go test -v --cover ./rpc/internal/..

.PHONY: fmt
fmt: # Format the codes | 格式化代码
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint: # Run go linter | 运行代码错误分析
	golangci-lint run -D staticcheck

.PHONY: tools
tools: # Install the necessary tools | 安装必要的工具
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: docker
docker: # Build the docker image | 构建 docker 镜像
	docker build -f Dockerfile-api -t ${DOCKER_USERNAME}/$(SERVICE_DASH)-api:${VERSION} .
	docker build -f Dockerfile-rpc -t ${DOCKER_USERNAME}/$(SERVICE_DASH)-rpc:${VERSION} .
	@echo "Build docker successfully"

.PHONY: publish-docker
publish-docker: # Publish docker image | 发布 docker 镜像
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin https://${REPO}
	docker push ${DOCKER_USERNAME}/$(SERVICE_DASH)-rpc:${VERSION}
	docker push ${DOCKER_USERNAME}/$(SERVICE_DASH)-api:${VERSION}
	@echo "Publish docker successfully"

.PHONY: gen-api
gen-api: # Generate API files | 生成 API 的代码
	goctls api go --api ./api/desc/all.api --dir ./api --trans_err=true --style=$(PROJECT_STYLE)
	swagger generate spec --output=./$(SERVICE_STYLE).$(SWAGGER_TYPE) --scan-models
	@echo "Generate API files successfully"

.PHONY: gen-rpc
gen-rpc: # Generate RPC files from proto | 生成 RPC 的代码
	goctls rpc protoc ./rpc/$(SERVICE_STYLE).proto --style=$(PROJECT_STYLE) --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc --style=$(PROJECT_STYLE)
	@echo "Generate RPC files successfully"

.PHONY: gen-ent
gen-ent: # Generate Ent codes | 生成 Ent 的代码
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./rpc/ent/template/*.tmpl" ./rpc/ent/schema --feature $(ENT_FEATURE)
	@echo "Generate Ent files successfully"

.PHONY: gen-rpc-ent-logic
gen-rpc-ent-logic: # Generate logic code from Ent, need model and group params | 根据 Ent 生成逻辑代码, 需要设置 model 和 group
	goctls rpc ent --schema=./rpc/ent/schema --style=$(PROJECT_STYLE) --import_prefix=/rpc --service_name=$(SERVICE) --project_name=$(SERVICE_STYLE) -o=./rpc --model=$(model) --group=$(group) --i18n=$(PROJECT_I18N) --proto_out=./rpc/desc/$(shell echo $(model) | tr A-Z a-z).proto --overwrite=true
	@echo "Generate logic codes from Ent successfully"

.PHONY: build-win
build-win: # Build project for Windows | 构建Windows下的可执行文件
	env CGO_ENABLED=0 GOOS=windows GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_rpc.exe ./rpc/$(SERVICE_STYLE).go
	env CGO_ENABLED=0 GOOS=windows GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_api.exe ./api/$(SERVICE_STYLE).go
	@echo "Build project for Windows successfully"

.PHONY: build-mac
build-mac: # Build project for MacOS | 构建MacOS下的可执行文件
	env CGO_ENABLED=0 GOOS=darwin GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_rpc ./rpc/$(SERVICE_STYLE).go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_api ./api/$(SERVICE_STYLE).go
	@echo "Build project for MacOS successfully"

.PHONY: build-linux
build-linux: # Build project for Linux | 构建Linux下的可执行文件
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_rpc ./rpc/$(SERVICE_STYLE).go
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE_STYLE)_api ./api/$(SERVICE_STYLE).go
	@echo "Build project for Linux successfully"

.PHONY: gen-swagger
gen-swagger: # Generate swagger file | 生成 swagger 文件
	swagger generate spec --output=./$(SERVICE_STYLE).$(SWAGGER_TYPE) --scan-models
	@echo "Generate swagger successfully"

.PHONY: serve-swagger
serve-swagger: # Run the swagger server | 运行 swagger 服务
	lsof -i:36666 | awk 'NR!=1 {print $2}' | xargs killall -9 || true
	swagger serve -F=swagger --port 36666 $(SERVICE_STYLE).$(SWAGGER_TYPE)
	@echo "Serve swagger-ui successfully"

.PHONY: help
help: # Show help | 显示帮助
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
