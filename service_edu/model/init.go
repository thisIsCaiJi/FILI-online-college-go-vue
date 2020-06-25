package model


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/thisIsCaiJi/online_college/service_edu/config"
)


type Database struct {
	Self *gorm.DB
}

// 单例
var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: GetDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		// "Asia%2FShanghai",  // 必须是 url.QueryEscape 的
		"Local",
	)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		logrus.Fatalf("数据库连接失败. 数据库名字: %s. 错误信息: %s", name, err)
	} else {
		logrus.Infof("数据库连接成功, 数据库名字: %s", name)
	}

	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(true)
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxOpenConns(20000)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(0)
}

func InitDB() *gorm.DB {
	conf := config.GetDataBaseConfig()
	return openDB(
		conf.Username,
		conf.Password,
		conf.Addr,
		conf.Name,
	)
}

func GetDB() *gorm.DB {
	return InitDB()
}
