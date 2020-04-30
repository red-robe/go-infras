package starter

import (
	"GoWebScaffold/infras/dao/mysqlDao"
	"fmt"
)

const (
	MysqlStarterKey = "_mysql_starter"
)

type MysqlStarter struct {
	Key string
	BaseStarter
}

func (s *MysqlStarter) Init(sctx *SystemContext) {
	// 创建一个mysql 连接客户端
	client := mysqlDao.NewMysqlClient(sctx.GetConfig())
	// 把客户端连接设置到上下文
	sctx.SetMysqlClient(client)
	fmt.Println("Mysql Client Init Successful!")
}
