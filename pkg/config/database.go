package config

import (
	"database/sql"
	"fmt"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"

	"time"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
)

const DefaultMaxOpenCon = 100

type DatabaseConf struct {
	Host         string
	Port         int
	Username     string `json:",optional"`
	Password     string `json:",optional"`
	DbName       string `json:",optional"`
	SSLMode      bool   `json:",optional"`
	Type         string `json:",optional"` // "postgres" or "mysql"
	MaxOpenConns *int   `json:",optional,default=100"`
	Debug        bool   `json:",optional,default=false"`
	AutoMigrate  bool   `json:",optional,default=false"`
}

func (c DatabaseConf) NewEntOption(redisConf redis2.RedisConf) ([]ent.Option, error) {
	var entOpts []ent.Option

	if c.Debug {
		logx.Info("Enabling Ent Client Request Debug")
		entOpts = append(entOpts, ent.Log(logx.Info))
		entOpts = append(entOpts, ent.Debug())
	}

	switch c.Type {
	case "mysql":
		driver, err := c.getEntDriver("mysql", dialect.MySQL, c.MysqlDSN(), redisConf)
		if err != nil {
			return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
		}
		entOpts = append(entOpts, ent.Driver(driver))
	case "postgres":
		driver, err := c.getEntDriver("postgres", dialect.Postgres, c.PostgresDSN(), redisConf)
		if err != nil {
			return nil, fmt.Errorf("failed to open the connection to postgresql: %v", err)
		}
		entOpts = append(entOpts, ent.Driver(driver))
	default:
		return nil, fmt.Errorf("unknown database type '%s'", c.Type)
	}

	return entOpts, nil
}

func (c DatabaseConf) getEntDriver(dbtype string, dialect string, dsn string, redisConf redis2.RedisConf) (*entcache.Driver, error) {
	db, err := sql.Open(dbtype, dsn)

	if err != nil {
		logx.Infof("failed to open the connection to %s: %v", dbtype, err)
		return nil, err
	}

	db.SetMaxOpenConns(*c.MaxOpenConns)
	driver := entsql.OpenDB(dialect, db)

	rdb := redis.NewClient(&redis.Options{Addr: redisConf.Host})

	cacheDrv := entcache.NewDriver(
		driver,
		entcache.TTL(time.Minute),
		entcache.Levels(
			entcache.NewLRU(256),
			entcache.NewRedis(rdb),
		),
	)

	return cacheDrv, nil
}

func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", c.Username, c.Password, c.Host, c.Port, c.DbName)
}

func (c DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", c.Host, c.Username, c.Password,
		c.DbName, c.Port, c.SSLMode)
}
