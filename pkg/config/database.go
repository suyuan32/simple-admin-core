package config

import (
	"database/sql"
	"fmt"

	"ariga.io/entcache"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"

	"time"
)

type DatabaseConf struct {
	Host         string
	Port         int
	Username     string `json:",optional"`
	Password     string `json:",optional"`
	DBName       string `json:",optional"`
	SSLMode      string `json:",optional"`
	Type         string `json:",optional"` // "postgres" or "mysql"
	MaxOpenConns *int   `json:",optional,default=100"`
	Debug        bool   `json:",optional,default=false"`
	CacheTime    int    `json:",optional,default=10"`
}

func (c DatabaseConf) NewCacheDriver(redisConf redis2.RedisConf) *entcache.Driver {
	db, err := sql.Open(c.Type, c.GetDSN())
	logx.Must(err)

	db.SetMaxOpenConns(*c.MaxOpenConns)
	driver := entsql.OpenDB(c.Type, db)

	rdb := redis.NewClient(&redis.Options{Addr: redisConf.Host})

	cacheDrv := entcache.NewDriver(
		driver,
		entcache.TTL(time.Duration(c.CacheTime)*time.Second),
		entcache.Levels(
			entcache.NewLRU(256),
			entcache.NewRedis(rdb),
		),
	)

	return cacheDrv
}

func (c DatabaseConf) NewNoCacheDriver() *entsql.Driver {
	db, err := sql.Open(c.Type, c.GetDSN())
	logx.Must(err)

	db.SetMaxOpenConns(*c.MaxOpenConns)
	driver := entsql.OpenDB(c.Type, db)

	return driver
}

func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", c.Username, c.Password, c.Host, c.Port, c.DBName)
}

func (c DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

func (c DatabaseConf) GetDSN() string {
	switch c.Type {
	case "mysql":
		return c.MysqlDSN()
	case "postgres":
		return c.PostgresDSN()
	default:
		return "mysql"
	}
}
