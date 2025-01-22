package core

import (
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"github.com/sirupsen/logrus"
)

func InitAddrDB() *geoip2.DBReader {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		logrus.Error("InitAddrDB err:", err)
	}
	return db
}
