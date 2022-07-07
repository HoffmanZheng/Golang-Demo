package models

import (
	"github.com/astaxie/beego/core/logs"
	"github.com/astaxie/beego/server/web"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error
var logger1 = logs.GetBeeLogger()

func init() {
	mysqladmin, err := web.AppConfig.String("mysqladmin")
	mysqlpwd, err := web.AppConfig.String("mysqlpwd")
	mysqldb, err := web.AppConfig.String("mysqldb")
	logger1.Info(mysqladmin + mysqlpwd + mysqldb)
	DB, err = gorm.Open("mysql", mysqladmin+":"+mysqlpwd+"@"+mysqldb+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		logger1.Error("failed to connect Mysql", err)
	} else {
		logger1.Info("connect to Mysql successfully!")
	}
}
