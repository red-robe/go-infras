package starter

import (
	"GoWebScaffold/infras/config"
	"GoWebScaffold/infras/dao/RedisDao"
	"GoWebScaffold/infras/dao/mongoDao"
	"GoWebScaffold/infras/dao/mysqlDao"
	"GoWebScaffold/infras/etcd"
	"GoWebScaffold/infras/mq/natsMq"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/garyburd/redigo/redis"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/nats-io/nats.go"
	qiuniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
	"go.uber.org/zap"
)

//系统启动器上下文
type SystemContext struct {
	starterOptions  starterOptionMap
	appConfig       *config.AppConfig
	commonLogger    *zap.Logger
	syncErrorLogger *zap.Logger
	mongoClient     *mongoDao.MongoClient
	mysqlClient     *mysqlDao.MysqlClient
	redisPool       *RedisDao.RedisPool
	etcdClient      *etcd.EtcdClient
	natsMQPool      *natsMq.NatsPool
	redisPubSubPool *redigo.Pool
	qiniuOSS        *qiuniuOss.Mac
	aliyunOSS       *aliOss.Client
}

type starterOptionMap map[string][]*StarterOption

// 设置启动器资源到启动器上下文
func (s *SystemContext) SetStarterOptions(opts ...*StarterOption) {
	if s.starterOptions == nil {
		s.starterOptions = make(starterOptionMap)
	}
	for _, o := range opts {
		if s.starterOptions[o.StarterKey] == nil {
			s.starterOptions[o.StarterKey] = make([]*StarterOption, 0)
		}
		s.starterOptions[o.StarterKey] = append(s.starterOptions[o.StarterKey], o)
	}
}

//从启动器上下文获取启动器资源
func (s *SystemContext) GetStarterOptions(starterKey string) []*StarterOption {
	if s.starterOptions != nil {
		if opts, ok := s.starterOptions[starterKey]; ok {
			return opts
		}
	}
	return nil
}

// 设置配置信息到启动器上下文
func (s *SystemContext) SetConfig(appConf *config.AppConfig) {
	if s.appConfig == nil {
		s.appConfig = appConf
	}
}

//从启动器上下文获取配置信息
func (s *SystemContext) GetConfig() *config.AppConfig {
	if s.appConfig == nil {
		panic("配置还为被初始化")
	}
	return s.appConfig
}

// 设置通用日志记录器到启动器上下文
func (s *SystemContext) SetCommonLogger(logger *zap.Logger) {
	if s.commonLogger == nil {
		s.commonLogger = logger
	}
}

//从启动器上下文获取通用日志记录器实例
func (s *SystemContext) GetCommonLogger() *zap.Logger {
	if s.commonLogger == nil {
		panic("系统通用日志记录器还未启动")
	}
	return s.commonLogger
}

// 设置异步错误日志记录器到启动器上下文
func (s *SystemContext) SetSyncErrorLogger(logger *zap.Logger) {
	if s.syncErrorLogger == nil {
		s.syncErrorLogger = logger
	}
}

//从启动器上下文获取异步错误日志记录器实例
func (s *SystemContext) GetSyncErrorLogger() *zap.Logger {
	if s.syncErrorLogger == nil {
		panic("系统异步错误日志记录器还未启动")
	}
	return s.syncErrorLogger
}

// 设置mongodb连接客户端
func (s *SystemContext) SetMongoClient(mp *mongoDao.MongoClient) {
	if s.mongoClient == nil {
		s.mongoClient = mp
	}
}

// 获取mongodb连接客户端
func (s *SystemContext) GetMongoClient() *mongoDao.MongoClient {
	if s.mongoClient == nil {
		panic("Mongodb Client 未初始化")
	}
	return s.mongoClient
}

// 设置mysql连接客户端
func (s *SystemContext) SetMysqlClient(m *mysqlDao.MysqlClient) {
	if s.mysqlClient == nil {
		s.mysqlClient = m
	}
}

// 获取mysql连接客户端
func (s *SystemContext) GetMysqlClient() *mysqlDao.MysqlClient {
	if s.mysqlClient == nil {
		panic("Mysql Client 未初始化")
	}
	return s.mysqlClient
}

// 设置redis连接池
func (s *SystemContext) SetRedisPool(p *RedisDao.RedisPool) {
	if s.redisPool == nil {
		s.redisPool = p
	}
}

// 获取redis连接客户端
func (s *SystemContext) GetRedisConn() redis.Conn {
	if s.redisPool == nil {
		panic("Redis 连接池未初始化")
	}
	return s.redisPool.GetRedisConn()
}

// 设置etcd连接客户端
func (s *SystemContext) SetEtcdClient(e *etcd.EtcdClient) {
	if s.etcdClient == nil {
		s.etcdClient = e
	}
}

// 获取etcd连接客户端
func (s *SystemContext) GetEtcdClient() *etcd.EtcdClient {
	if s.etcdClient == nil {
		panic("ETCD Client 未初始化")
	}
	return s.etcdClient
}

// 设置natsMQ连接客户端
func (s *SystemContext) SetNatsMQPool(mq *natsMq.NatsPool) {
	if s.natsMQPool == nil {
		s.natsMQPool = mq
	}
}

// 获取NatsMQ连接客户端
func (s *SystemContext) GetNatsMQConn() (*nats.Conn, error) {
	if s.natsMQPool == nil {
		panic("NatsMQ Pool 未初始化")
	}
	return s.natsMQPool.Get()
}

// 设置redisPubSub连接池
func (s *SystemContext) SetRedisPubSubMPool(mq *redigo.Pool) {
	if s.redisPubSubPool == nil {
		s.redisPubSubPool = mq
	}
}

// 获取redisPubSub连接
func (s *SystemContext) GetRedisPubSubConn() redigo.Conn {
	if s.redisPubSubPool == nil {
		panic("redisPubSub Pool 未初始化")
	}
	return s.redisPubSubPool.Get()
}

// 设置aliyunOSS连接客户端
func (s *SystemContext) SetAliyunOSSClient(c *aliOss.Client) {
	if s.aliyunOSS == nil {
		s.aliyunOSS = c
	}
}

// 获取aliyunOss连接
func (s *SystemContext) GetAliyunOSSConn() *aliOss.Client {
	if s.aliyunOSS == nil {
		panic("AliyunOSS Client 未初始化")
	}
	return s.aliyunOSS
}

// 设置qiuniuOSS连接客户端
func (s *SystemContext) SetQiniuOSS(c *qiuniuOss.Mac) {
	if s.qiniuOSS == nil {
		s.qiniuOSS = c
	}
}

// 获取qiniuOss连接
func (s *SystemContext) GetQiniuOSS() *qiuniuOss.Mac {
	if s.qiniuOSS == nil {
		panic("QiniuOSS Client 未初始化")
	}
	return s.qiniuOSS
}
