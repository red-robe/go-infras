package RedisDao

import (
	"GoWebScaffold/infras/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

type RedisPool struct {
	pool *redis.Pool
}

func NewRedisPool(config *config.AppConfig) *RedisPool {
	redisPool := new(RedisPool)

	// 配置并获得一个连接池对象的指针
	redisPool.pool = &redis.Pool{
		// 最大活动链接数。0为无限
		MaxActive: int(config.RedisConf.MaxActive),
		// 最大闲置链接数，0为无限
		MaxIdle: int(config.RedisConf.MaxIdle),
		// 闲置链接超时时间
		IdleTimeout: time.Duration(config.RedisConf.IdleTimeout) * time.Second,
		// 连接池的连接拨号
		Dial: func() (redis.Conn, error) {
			// 连接
			redisAddr := config.RedisConf.DbHost + ":" + strconv.Itoa(int(config.RedisConf.DbPort))
			conn, err := redis.Dial("tcp", redisAddr)
			if err != nil {
				fmt.Println("redis dial fatal:", err.Error())
				return nil, err
			}
			// 权限认证
			if config.RedisConf.DbAuth {
				if _, err := conn.Do("Auth", config.RedisConf.DbPasswd); err != nil {
					fmt.Println("redis auth fatal:", err.Error())
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},

		// 定时检测连接是否可用
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("Ping")
			if err != nil {
				fmt.Println("")
			}
			return err
		},
	}

	// 一般启动后不关闭连接池
	fmt.Println("Redis pool init ready!")
	return redisPool
}

// 从Redis连接池获取一个连接
func (p *RedisPool) GetRedisConn() redis.Conn {
	conn := p.pool.Get()
	return conn
}
