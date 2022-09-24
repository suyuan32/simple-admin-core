# Environment setting

## Environment Requirement
- golang 1.19
- nodejs 18.8.0
- consul
- mysql

## Backend deployment

### simple admin core
simple admin core is the core codes of the system, it must be running. 

#### Default account
username:     admin  \
password:     simple-admin

### Get codes
```bash
git clone https://github.com/suyuan32/simple-admin-core.git
```

### configure 

[Consul](/simple-admin/en/docs/consul.md)

### Sync dependencies

```shell 
go mod tidy
```


### Run rpc service

```bash
cd rpc

go run core.go -f etc/core.yaml
```


### Run api service

```bash
cd api

go run core.go -f etc/core.yaml
```

## Front end configuration

### clone the code

```shell
git clone https://github.com/suyuan32/simple-admin-backend-ui.git
```

### Sync dependencies

```shell
yarn install
```

### Run server

```shell
yarn serve
```

### Compile
```shell
yarn build
```

### Preview
```shell
# build and preview
yarn preview

# preview exist files
yarn preview:dist
```

## Initialize the database

***Important:***  You must create the database before initialize
The database name is the same as your configuration.

```shell
# visit the url 

https://address:port/#/init
```

You can see the ui to do this.

![pic](../../assets/init_en.png)

## **When the initialization finished, you should restart api and rpc service.**