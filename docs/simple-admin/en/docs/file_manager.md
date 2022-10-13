## File manager service

> Get code
```shell
git clone https://github.com/suyuan32/simple-admin-file.git
```

> Modify configuration file

```yaml
Name: file_manager_0
Host: 0.0.0.0
Port: 9102
MaxBytes: 1073741824 # max content length : 1 gb
Timeout: 30000 # bigger max bytes need longer timeout

Auth:
  AccessSecret: jS6VKDtsJf3z1n2VKDtsJf3z1n2
  AccessExpire: 259200  # Seconds

Log:
  ServiceName: fileManagerLogger
  Mode: file
  Path: /home/ryan/logs/file/api
  Level: info
  Compress: false
  KeepDays: 7
  StackCoolDownMillis: 100

RedisConf:
  Host: 127.0.0.1:6379
  Type: node

DatabaseConf:
  Type: mysql
  Path: 127.0.0.1
  Port: 3306
  Config: charset=utf8mb4&parseTime=True&loc=Local
  DBName: simple_admin_file
  Username:
  Password:
  MaxIdleConn: 10
  MaxOpenConn: 100
  LogMode: error
  LogZap: false

UploadConf:
  MaxImageSize: 33554432  # 32 mb the maximum size of image
  MaxVideoSize: 1073741824 # 1gb the maximum size of video
  MaxAudioSize: 33554432  # 32mb the maximum size of audio
  MaxOtherSize: 10485760  # 10 mb the maximum size of other type
  PrivateStorePath: /home/ryan/www/private  # private path 
  PublicStorePath: /home/ryan/www/public  # public path for every one access e.g. nginx path
```

> You should use nginx to set PublicStorePath as static path for front end.
> Make sure AccessSecret is the same as simple-admin-core' api set
> The configuration is similar as core
> Run code the same as core

### K8s Deployment
> It is similar with core api.

You should do these step:
- deploy the images via fileapi.yaml
- modify simple-admin-backend-ui/deploy/default.conf, uncomment the file manager rule
- update ingress configmap
- update ingress controller