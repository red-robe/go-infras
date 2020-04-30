package mysqlDao

import (
	"GoWebScaffold/infras/config"
	"GoWebScaffold/infras/utils"
	"database/sql"
	"fmt"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlClient struct {
	Db *sql.DB
}

func NewMysqlClient(config *config.AppConfig) *MysqlClient {
	var err error
	client := new(MysqlClient)
	client.Db, err = manager.New(config.MysqlConf.DbName, config.MysqlConf.DbUser, config.MysqlConf.DbPasswd, config.MysqlConf.DbHost).Set(
		manager.SetCharset(config.MysqlConf.ChartSet),                                   // 设置编码类型：utf8
		manager.SetAllowCleartextPasswords(config.MysqlConf.AllowCleartextPasswords),    // 开发环境中设置允许明文密码通信
		manager.SetInterpolateParams(config.MysqlConf.InterpolateParams),                // 设置允许占位符参数
		manager.SetTimeout(time.Duration(config.MysqlConf.Timeout)*time.Second),         // 连接超时时间
		manager.SetReadTimeout(time.Duration(config.MysqlConf.ReadTimeout)*time.Second), // 读超时时间
		manager.SetParseTime(config.MysqlConf.ParseTime),                                // 将数据库的datetime时间格式转换为go time包数据类型
	).Port(int(config.MysqlConf.DbPort)).Open(config.MysqlConf.PING)
	utils.FailHandler(err)

	client.Db.SetConnMaxLifetime(time.Duration(config.MysqlConf.ConnMaxLifetime) * time.Second) // 设置最大的连接时间，1分钟
	client.Db.SetMaxIdleConns(int(config.MysqlConf.MaxIdleConns))                               // 设置最大的闲置连接数
	client.Db.SetMaxOpenConns(int(config.MysqlConf.MaxOpenConns))                               // 设置最大的连接数

	err = client.Db.Ping()
	if err != nil {
		utils.FailHandler(err)
	} else {
		fmt.Println("Mysql pool init ready!")
	}

	return client
}
