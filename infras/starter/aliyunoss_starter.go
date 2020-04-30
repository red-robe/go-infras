package starter

import "GoWebScaffold/infras/oss/aliyunOss"

const (
	AliyunOssStarterKey = "_aliyun_oss_starter"
)

type AliyunOssStarter struct {
	Key string
	BaseStarter
}

func (s *AliyunOssStarter) Init(sctx *SystemContext) {
	client := aliyunOss.Init(sctx.GetConfig())
	sctx.SetAliyunOSSClient(client)
}
