# Simple admin tools
Simple admin tools is a tool fork from go-zero.
It provides more addition features than origin project such as:
- go-swagger : it is different with origin which uses @doc comments
- consul: default use consul to do service discovery and configuration
- multi-language support
- optimize error message
- fully support validator and easy to use
- so on

> But it is a little complex to install the goctls and import the dependencies due to forking.\
Let me show you how to build the code by yourself.

> Build goctls

```shell
git clone https://github.com/suyuan32/simple-admin-tools.git

cd tools/goctl

go mod tidy

# output goctls in order not to conflict with goctl
go build -o goctls goctl.go

cp ./goctls $GOPATH/bin/goctls
```
It is easy right?

> How to use it?

> Auto fix all dependencies command

```shell
goctls env check -i -f --verbose
```
Run this command can auto install protoc and so on.

> API command

The command is the same as goctl but you should use goctls instead.
```shell
$ goctl api -h
NAME:
   goctl api - generate api related files

USAGE:
   goctl api command [command options] [arguments...]

COMMANDS:
   new       fast create api service
   format    format api files
   validate  validate api file
   doc       generate doc files
   go        generate go files for provided api in yaml file
   java      generate java files for provided api in api file
   ts        generate ts files for provided api in api file
   dart      generate dart files for provided api in api file
   kt        generate kotlin code for provided api file
   plugin    custom file generator

OPTIONS:
   -o value        output a sample api file
   --home value    the goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
   --remote value  the remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                   The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
   --branch value  the branch of the remote repo, it does work with --remote
   --help, -h      show help
```

> Example:

```shell
goctls api go -api core.api -dir .
```
This means generating go files by core.api's declaration in current directory. -dir set the output path.

> Rpc command

```shell
$ goctl rpc protoc -h
NAME:
   goctl rpc protoc - generate grpc code

USAGE:
   example: goctl rpc protoc xx.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

DESCRIPTION:
   for details, see https://go-zero.dev/cn/goctl-rpc.html

OPTIONS:
   --zrpc_out value  the zrpc output directory
   --style value     the file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]
   --home value      the goctl home path of the template
   --remote value    the remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                     The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
   --branch value    the branch of the remote repo, it does work with --remote
   --verbose, -v     enable log output
```

> Example: \
Generate proto template

```shell
goctl rpc template -o=user.proto
```

```shell
syntax = "proto3";

package user;
option go_package=". /user";

message Request {
  string ping = 1;
}

message Response {
  string pong = 1;
}

service User {
  rpc Ping(Request) returns(Response);
}
```
> Generate go files

```shell
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```
[More](https://go-zero.dev/docs/goctl/zrpc)

> Project go.mod setting

We know that if we want to import fork project we should use replace command.

```mod
module github.com/suyuan32/simple-admin-core

go 1.19

require (
	github.com/casbin/casbin/v2 v2.52.1
	github.com/casbin/gorm-adapter/v3 v3.7.4
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/validator/v10 v10.11.1
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/google/uuid v1.3.0
	github.com/mojocn/base64Captcha v1.3.5
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.0
	github.com/suyuan32/simple-admin-tools/plugins/registry/consul v0.0.0-20220923060146-bde681863b8d
	github.com/zeromicro/go-zero v1.4.1
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gorm.io/gorm v1.23.8
)

replace github.com/zeromicro/go-zero v1.4.1 => github.com/suyuan32/simple-admin-tools v0.0.6
```
> In order to make it easier. You can use the command in goctls.

```shell
goctls migrate --zero-version v1.4.1 --tool-version v0.0.6
```
> It can help you to add replace code but it cannot run multiple times because it will add multiple replace lines
> in the go.mod file. You can just edit go.mod file can modify the simple-admin-tools version manually and run
> **go mod tidy**.

