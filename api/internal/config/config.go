package config

import (
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	RedisConf    redis.RedisConf
	CoreRpc      zrpc.RpcClientConf
	Captcha      Captcha
	DatabaseConf gormsql.GORMConf
}

type Captcha struct {
	KeyLong   int // captcha length
	ImgWidth  int // captcha width
	ImgHeight int // captcha height
}
