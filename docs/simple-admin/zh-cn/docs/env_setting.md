# 环境配置

## 环境需求
- golang 1.19
- nodejs 18.8.0
- consul
- mysql 

## 后端部署

### simple admin core
simple admin core 是核心代码，必须运行

### 下载代码 
```bash
git clone https://github.com/suyuan32/simple-admin-core.git
```

### 配置

[Consul](/simple-admin/zh-cn/docs/consul.md)

### 配置依赖

```shell 
go mod tidy
```


### 运行 rpc 服务

```bash
cd rpc

go run core.go -f etc/core.yaml
```


### 运行 api 服务

```bash
cd api

go run core.go -f etc/core.yaml
```

## 前端配置

### 下载代码

```shell
git clone https://github.com/suyuan32/simple-admin-backend-ui.git
```

### 下载依赖

```shell
yarn install
```

### 运行

```shell
yarn serve
```

### 编译
```shell
yarn build
```

### 预览
```shell
# build and preview
yarn preview

# preview exist files
yarn preview:dist
```

## 初始化数据库
***重要:*** 在初始化数据库前必须先创建数据库, 数据库名称和配置文件中的名称相同.

```shell
# 访问前端地址端口

https://address:port/#/init
```

进入界面

![pic](../../assets/init_zh_cn.png)