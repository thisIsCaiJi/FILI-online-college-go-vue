package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/thisIsCaiJi/online_college/service_edu/config"
)

var db *gorm.DB

func init() {
	if gormDB, err := createDatabase(); err == nil {
		db = gormDB
	} else {
		logrus.WithError(err).Fatalln("create database connection failed")
	}
	//enable Gorm mysql log
	db.LogMode(true)
}

//Close clear db collection
func Close() {
	if db != nil {
		db.Close()
	}
}

func createDatabase() (*gorm.DB, error) {
	dpconf := config.GetDataBaseConfig()
	dbAddr := dpconf.Addr
	dbName := dpconf.Name
	dbUser := dpconf.Username
	dbPassword := dpconf.Password
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbAddr, dbName)
	return gorm.Open("mysql", conn)
}
