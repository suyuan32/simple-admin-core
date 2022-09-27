## 文件上传服务

### 下载代码
```shell
git clone https://github.com/suyuan32/simple-admin-file.git
```

### 修改配置文件

添加 kv 进 consul

#### key
fileApiConf

#### value
```yaml
Name: file_manager_0
Host: 0.0.0.0
Port: 9102
MaxBytes: 1073741824 # max content length : 1 gb
Timeout: 30000 # bigger max bytes need longer timeout
Auth:
  AccessSecret:
  AccessExpire: 259200  # Seconds
Log:
  ServiceName: fileManager
  Mode: file
  Path: /home/ryan/logs/file/api
  Level: info
  Compress: false
  KeepDays: 7
  StackCooldownMillis: 100
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
  MaxImageSize: 33554432  # 32 mb
  MaxVideoSize: 1073741824 # 1gb
  MaxAudioSize: 33554432  # 32mb
  MaxOtherSize: 10485760  # 10 mb
  PrivateStorePath: /home/ryan/www/private  # private
  PublicStorePath: /home/ryan/www/public  # public path for every one access e.g. nginx path
```
你可以使用nginx 将 PublicStorePath 转发为静态地址方便前端调用

确保 AccessSecret 和 simple-admin-core的api配置文件内一致
配置方式参考core
运行方式同理