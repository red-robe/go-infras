package starter

import (
	"GoWebScaffold/infras/etcd"
	"context"
)

const (
	EtcdStarterKey = "_etcd_starter"
)

type EtcdStarter struct {
	Key string
	BaseStarter
}

func (s *EtcdStarter) Init(sctx *SystemContext) {
	client := etcd.NewEtcdClient(context.TODO(), sctx.appConfig, nil)
	sctx.SetEtcdClient(client)
}
