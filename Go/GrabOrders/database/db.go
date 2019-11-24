package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Username string
	Password string
	Address  string
	DBName   string
}

func GetConnection() *gorm.DB {
	database := &Connection{
		Username: "root",
		Password: "root",
		Address:  "127.0.0.1:3306",
		DBName:   "sample_db",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&multiStatements=true", database.Username, database.Password, database.Address, database.DBName)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Printf("main: mysql open failed: %v", err)
	}

	return db
}
