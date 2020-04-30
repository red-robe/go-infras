package starter

// 主程序启动时引入infras包自动注册基础资源启动器
func init() {
	//配置启动器必须首先启动
	Register(&ConfigStarter{Key: ConfigStarterKey})
	Register(&LoggerStarter{Key: LoggerStarterKey})
	Register(&MongoDBStarter{Key: MongoDBStarterKey})
	Register(&MysqlStarter{Key: MysqlStarterKey})
}
