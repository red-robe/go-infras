package starter

import "GoWebScaffold/infras/mq/natsMq"

const (
	NatsMQStarterKey = "_natsmq_starter"
)

type NatsMQStarter struct {
	Key string
	BaseStarter
}

func (s *NatsMQStarter) Init(sctx *SystemContext) {
	natsMqPool := natsMq.NatsMqPoolInit(sctx.GetConfig(), sctx.GetCommonLogger())
	sctx.SetNatsMQPool(natsMqPool)
}
