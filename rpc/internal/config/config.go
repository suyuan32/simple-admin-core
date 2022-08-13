package config

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB        DatabaseConf    `json:"DatabaseConf" yaml:"DatabaseConf"`
	LogConf   LogConf         `json:"LogConf" yaml:"LogConf"`
	RedisConf redis.RedisConf `json:"RedisConf" yaml:"RedisConf"`
}

type DatabaseConf struct {
	Type        string `json:"Type" yaml:"Type"`               // type of database: mysql, postpres
	Path        string `json:"Path" yaml:"Path"`               // address
	Port        int    `json:"Port" yaml:"Port"`               // port
	Config      string `json:"Config" yaml:"Config"`           // extra config such as charset=utf8mb4&parseTime=True
	Dbname      string `json:"DBName" yaml:"DBName"`           // database name
	Username    string `json:"Username" yaml:"Username"`       // username
	Password    string `json:"Password" yaml:"Password"`       // password
	MaxIdleConn int    `json:"MaxIdleConn" yaml:"MaxIdleConn"` // the maximum number of connections in the idle connection pool
	MaxOpenConn int    `json:"MaxOpenConn" yaml:"MaxOpenConn"` // the maximum number of open connections to the database
	LogMode     string `json:"LogMode" yaml:"LogMode"`         // open gorm's global logger
	LogZap      bool   `json:"LogZap" yaml:"LogZap"`
}

func (d *DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", d.Username, d.Password, d.Path, d.Port, d.Dbname, d.Config)
}

func (d *DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s", d.Path, d.Username, d.Password,
		d.Dbname, d.Port, d.Config)
}

type LogConf struct {
	ServiceName         string `json:",optional"`                                    // service name
	Mode                string `json:",default=console,options=console|file|volume"` // mode: console-output to console，file-output to file，，volume-output to the docker volume
	Path                string `json:",default=logs"`                                // store path
	Level               string `json:",default=info,options=info|error|severe"`      // the level to be shown
	Compress            bool   `json:",optional"`                                    // gzip compress
	KeepDays            int    `json:",optional"`                                    // the period to be stored
	StackCooldownMillis int    `json:",default=100"`                                 // the period between two writing (ms)
}
