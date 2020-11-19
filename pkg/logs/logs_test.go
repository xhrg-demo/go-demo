package logs

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestLogs(t *testing.T) {
	InitLog()
	i := 1
	for true {
		i = i + 1
		logrus.Infof("i=%v", i)
		time.Sleep(time.Second * 1)
	}
}
