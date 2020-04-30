package starter

import (
	"GoWebScaffold/infras/logger"
	"io"
)

const (
	LoggerStarterKey = "_logger_starter"
)

type LoggerStarter struct {
	Key string
	BaseStarter
}

//初始化日志记录器并设置到应用启动器上下文
func (s *LoggerStarter) Init(sctx *SystemContext) {
	// 获取日志记录启动器的相关选项资源,此处需要异步记录器
	options := sctx.GetStarterOptions(LoggerStarterKey)
	var sw []io.Writer
	for _, opt := range options {
		if writer, ok := opt.Resource.(io.Writer); ok {
			sw = append(sw, writer)
		}
	}
	sctx.SetCommonLogger(logger.CommonLogger(sctx.GetConfig(), sw...))
	sctx.SetSyncErrorLogger(logger.SyncErrorLogger(sctx.GetConfig()))
}
