# Environment setting

## Environment Requirement
- golang 1.19
- nodejs 18.8.0
- etcd
- mysql

## Backend deployment

### simple admin core
simple admin core is the core codes of the system, it must be running.

### Get codes
```bash
git clone https://github.com/suyuan32/simple-admin-core.git
```

### configure api

api/etc/core.yaml

```yaml
Name: core.api  # service name
Host: 0.0.0.0  # bind address
Port: 8500  # port 
Auth:
  AccessSecret:    # jwt encrypt secret
  AccessExpire:   #  expire time (second)
Log:
  ServiceName: coreApiLogger # logger name
  Mode: file  # type file is storing as file | console is output on console
  Path: /data/logs/api  # store path
  Level: info  # level:  info/error
  Compress: false # compress file
  KeepDays: 7 # store period
  StackCooldownMillis: 100
RedisConf:
  Host: 127.0.0.1:6379  # redis address
  Type: node       # node type
  Pass: xxx    # password
CoreRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379  # etcd address
    Key: core.rpc  # rpc service name (the same as configuration in rpc)
Captcha:
  KeyLong: 5   # captcha answer length
  ImgWidth: 240  # captcha image width
  ImgHeight: 80  # captcha image height
DatabaseConf:
  Type: mysql  # database type: mysql or postgresql
  Path: 127.0.0.1  # address
  Port: 3306  
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin # database name
  Username:    # username 
  Password:    # password
  MaxIdleConn: 10  # max idle connections
  MaxOpenConn: 100  # max opening connections
  LogMode: error  # log level
  LogZap: false   # use zap logger
```

### Configure rpc

rpc/etc/core.yaml

```yaml
Name: core.rpc    # service name 
ListenOn: 127.0.0.1:8501   # bind address
Etcd:
  Hosts:
    - 127.0.0.1:2379   # etcd address
  Key: core.rpc        # service name, used for service discover
  User: root           # username 
  Pass: xxx            # password
DatabaseConf:
  Type: mysql  # database type: mysql or postgresql
  Path: 127.0.0.1  # address
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin # database name
  Username:    # username 
  Password:    # password
  MaxIdleConn: 10  # max idle connections
  MaxOpenConn: 100  # max opening connections
  LogMode: error  # log level
  LogZap: false   # use zap logger
Log:
  ServiceName: coreApiLogger # logger name
  Mode: file  # type file is storing as file | console is output on console
  Path: /data/logs/api  # store path
  Level: info  # level:  info/error
  Compress: false # compress file
  KeepDays: 7 # store period
  StackCooldownMillis: 100
RedisConf:
  Host: 127.0.0.1:6379  # redis address
  Type: node       # node type
  Pass: xxx    # password
```

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

## Initialize the database

***Important:***  You must create the database before initialize
The database name is the same as your configuration.

```shell
# visit the url 


https://address:port/#/init
```

You can see the ui to do this.

![pic](../../assets/init_en.png)
