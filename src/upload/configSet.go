package upload

import "github.com/spf13/viper"

type ossConfigValue struct {
	accessKeyID     string
	accessKeySecret string
	endpoint        string
	bucketName      string
}

//为配置设置默认值
func (o *ossConfigValue) ossConfigSet() {
	if viper.GetString("oss.accessKeyID") != "" {
		o.accessKeyID = viper.GetString("oss.accessKeyID")
	} else {
		o.accessKeyID = "123"
	}

	if viper.GetString("oss.accessKeySecret") != "" {
		o.accessKeySecret = viper.GetString("oss.accessKeySecret")
	} else {
		o.accessKeySecret = "123"
	}
	if viper.GetString("oss.endpoint") != "" {
		o.endpoint = viper.GetString("oss.endpoint")
	} else {
		o.endpoint = "oss-cn-zhangjiakou-internal.aliyuncs.com"
		//o.endpoint = "oss-cn-zhangjiakou.aliyuncs.com"

	}
	if viper.GetString("oss.bucketName") != "" {
		o.bucketName = viper.GetString("oss.bucketName")
	} else {
		o.bucketName = "backup"
	}

}
