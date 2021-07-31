package LogGet

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func writerPath() {
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	 `WithMaxAge` 设置文件清理前的最长保存时间
	 `WithRotationCount` 设置文件清理前最多保存的个数

	*/
	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，的自动清理掉。
	viper.SetDefault("log.Path", "log/")
	viper.SetDefault("log.LogLevel", "info")

	path := viper.GetString("log.path")
	writer, _ := rotatelogs.New(
		path+"ossupload"+".%Y%m%d%H%M"+".log",
		//rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(720)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(720)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	//fmt.Println(writer)
	logrus.SetOutput(writer)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//var logLevel = viper.GetString("log.LogLevel")
	setLevel(viper.GetString("log.LogLevel"))
}

func Info(message string, Context ...*map[string]string) {
	writerPath()
	logger := logrus.WithFields(logrus.Fields{"Context": &Context})
	logger.Infoln(message)

}

func Warn(message string, Context ...*map[string]string) {

	writerPath()
	logger := logrus.WithFields(logrus.Fields{"Context": &Context})
	logger.Warnln(message)
}

func Error(message string, Context ...*map[string]string) {
	writerPath()
	logger := logrus.WithFields(logrus.Fields{"Context": &Context})
	logger.Errorln(message)
}

func Fatal(message string, Context ...*map[string]string) {
	writerPath()
	logger := logrus.WithFields(logrus.Fields{"Context": &Context})
	logger.Fatalln(message)

}
func Debug(message string, Context ...*map[string]string) {
	writerPath()
	logger := logrus.WithFields(logrus.Fields{"Context": &Context})
	logger.Debugln(message)
}

func setLevel(lvl string) {
	switch strings.ToLower(lvl) {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

}
