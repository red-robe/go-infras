package starter

import (
	"GoWebScaffold/infras/dao/RedisDao"
	"fmt"
)

const (
	RedisStarterKey = "_redis_starter"
)

type RedisStarter struct {
	Key string
	BaseStarter
}

func (s *RedisStarter) Init(sctx *SystemContext) {
	// 创建一个redis连接池
	pool := RedisDao.NewRedisPool(sctx.GetConfig())
	// 把连接池设置到上下文
	sctx.SetRedisPool(pool)
	fmt.Println("Redis Pool Client Init Successful!")
}
