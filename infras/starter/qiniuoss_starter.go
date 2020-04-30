package starter

import "GoWebScaffold/infras/oss/qiniuOss"

const (
	QiniuOssStarterKey = "_qiniu_oss_starter"
)

type QiniuOssStarter struct {
	Key string
	BaseStarter
}

func (s *QiniuOssStarter) Init(sctx *SystemContext) {
	mac := qiniuOss.Init(sctx.GetConfig())
	sctx.SetQiniuOSS(mac)
}
