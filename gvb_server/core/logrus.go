package core

// 日志系统

import (
	"bytes"
	"fmt"
	"gvb_server/global"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	log := global.Config.Logger
	// 等级颜色选择
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日期格式
	timeStamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		// 自定义输出格式
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timeStamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timeStamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	log := global.Config.Logger
	mylog := logrus.New()               // 新建一个实例
	mylog.SetOutput(os.Stdout)          // 设置输出类型
	mylog.SetReportCaller(log.Showline) // 开启返回函数名和行号
	mylog.SetFormatter(&LogFormatter{}) // 设置自定义的formatter

	level, err := logrus.ParseLevel(log.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mylog.SetLevel(level) // 设置最低的level
	return mylog
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout) // 设置输出类型
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
