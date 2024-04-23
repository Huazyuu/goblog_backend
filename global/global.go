package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/config"
)

var (
	Config *config.Config
	DB     *gorm.DB

	Logger   *logrus.Logger
	MySqlLog logger.Interface
)
