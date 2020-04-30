package starter

import (
	"GoWebScaffold/infras/config"
)

const (
	ConfigStarterKey = "_config_starter"
)

type ConfigStarter struct {
	Key string
	BaseStarter
}

//启动器初始化
func (s *ConfigStarter) Init(sctx *SystemContext) {
	//启动时先加载配置文件并解析,设置解析后的配置信息到应用启动器上下文
	sctx.SetConfig(config.Parse())
}
