package datastore

import (
	"codly/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strconv"
	"time"
)

type datastore struct {
	engine *gorm.DB
}

func (db *datastore) DB() *gorm.DB {
	return db.engine
}

var Datastore datastore

func init() {
	if Datastore.engine != nil {
		return
	}
	var err error
	dbName := config.Database.DBName
	username := config.Database.Username
	password := config.Database.Password
	host := config.Database.Host
	port := strconv.Itoa(config.Database.Port)
	Datastore.engine, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbName,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		log.Fatal("open database error: " + err.Error())
	}
	sqlDB, _ := Datastore.engine.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(1000)                 //必须小于等于数据库max_connections参数值
	sqlDB.SetConnMaxLifetime(110 * time.Second) //必须小于数据库wait_timeout参数值
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
