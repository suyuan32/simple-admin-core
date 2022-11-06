package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-core/pkg/config"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	RedisConf    redis.RedisConf
	CoreRpc      zrpc.RpcClientConf
	Captcha      Captcha
	DatabaseConf config.DatabaseConf
	CasbinConf   config.CasbinConf
}

type Captcha struct {
	KeyLong   int // captcha length
	ImgWidth  int // captcha width
	ImgHeight int // captcha height
}
