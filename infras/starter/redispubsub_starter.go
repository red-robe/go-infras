package starter

import "GoWebScaffold/infras/mq/redisPubSub"

const (
	RedisPubsubStarterKey = "_redis_pubsub_starter"
)

type RedisPubSubStarter struct {
	Key string
	BaseStarter
}

func (s *RedisPubSubStarter) Init(sctx *SystemContext) {
	redisPS := redisPubSub.RedisMqInit(sctx.GetConfig(), sctx.GetCommonLogger())
	sctx.SetRedisPubSubMPool(redisPS)
}
