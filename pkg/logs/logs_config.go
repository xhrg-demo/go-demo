package logs

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

type LogFormatter struct{}

func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("[%s] [%s] %s:%d %s\n", timestamp, entry.Level.String(), entry.Caller.File, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}

func InitLog() {
	level := logrus.InfoLevel
	//初始化日志组件
	logrus.SetFormatter(new(LogFormatter))
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	path := "/Users/xhrg/temp/logs/main.log"
	logfile, err := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithRotationTime(time.Hour*24*7),
		rotatelogs.WithRotationCount(uint(50)),
		rotatelogs.ForceNewFile(),
	)
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	//输出到日志文件的同事打印控制台
	logrus.SetOutput(io.MultiWriter(logfile, os.Stdout))
}
