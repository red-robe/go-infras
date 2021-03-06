package aliyunOss

import (
	"GoWebScaffold/infras/config"
	"GoWebScaffold/infras/utils"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func Init(appConf *config.AppConfig) *aliOss.Client {
	// Aliyun OSS初始化
	if appConf.OssConf.Aliyun.Switch {
		var err error
		var aoss *aliOss.Client
		aoss, err = aliOss.New(
			appConf.OssConf.Aliyun.Endpoint,
			appConf.OssConf.Aliyun.AccessKeyId,
			appConf.OssConf.Aliyun.AccessKeySecret,
			aliOss.Timeout(int64(appConf.OssConf.Aliyun.ConnTimeout), int64(appConf.OssConf.Aliyun.RWTimeout)),
			aliOss.UseCname(appConf.OssConf.Aliyun.UseCname),
			aliOss.EnableCRC(appConf.OssConf.Aliyun.EnableCRC),
		)
		utils.FailHandler(err)
		return aoss
	}
	return nil
}
