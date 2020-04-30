package starter

import (
	"GoWebScaffold/infras/dao/mongoDao"
	"fmt"
)

const (
	MongoDBStarterKey = "_mongodb_starter"
)

type MongoDBStarter struct {
	Key string
	BaseStarter
}

func (s *MongoDBStarter) Init(sctx *SystemContext) {
	// 创建一个mongo session 连接池
	client := mongoDao.NewMongoClient(sctx.GetConfig())
	// 把连接池设置到上下文
	sctx.SetMongoClient(client)
	fmt.Println("MongoDB Client Init Successful!")
}
