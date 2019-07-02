package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbName,
	)

	return gorm.Open("mysql", connectionString)
}
