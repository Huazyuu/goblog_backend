package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"io"
	"log"
	"os"
	"path"
)

// color
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct {
}

// Format 实现Formatter(entry *logrus.Entry)([]byte, error) interface
func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// 颜色配置
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
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

	logger := global.Config.Logger

	// 时间格式自定义
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 文件路劲
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 输出格式
		_, _ = fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", logger.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()
	writer1 := os.Stdout
	writer2, err := os.OpenFile(global.Config.Logger.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("create file %s failed: %v", global.Config.Logger.Path, err)
	}
	mLog.SetOutput(io.MultiWriter(writer2, writer1))    // 输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 行号
	mLog.SetFormatter(&LogFormatter{})                  // 自定义
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	InitDefaultLogger()
	// logrus.Info(" log uploads success")
	return mLog
}
func InitDefaultLogger() {
	writer1 := os.Stdout
	writer2, err := os.OpenFile(global.Config.Logger.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("create file %s failed: %v", global.Config.Logger.Path, err)
	}
	logrus.SetOutput(io.MultiWriter(writer2, writer1))    // 输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // 行号
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低的level
}
