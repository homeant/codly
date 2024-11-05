package datastore

import (
	"gorm.io/gorm"

	"github.com/glebarez/sqlite"
)

func InitDB() (*gorm.DB, error) {
	dbConn, err := gorm.Open(sqlite.Open("codly.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
