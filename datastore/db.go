package datastore

import (
	"codly/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

func InitDB(config model.DatabaseConfig) (*gorm.DB, error) {
	// 构建MySQL连接字符串
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DBName
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
