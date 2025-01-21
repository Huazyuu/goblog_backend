package log_stash

import (
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logrus.Info(os.Getwd())
	core.InitCore("../../settings.yaml")
	global.DB = core.InitGorm()
	log := New("192.168.0.0", "xxx")
	log.Warn("test")

}
