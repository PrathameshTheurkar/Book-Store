package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	godotenv.Load()
	d, err := gorm.Open("mysql", os.Getenv("SQL_PATH"))
	if err != nil {
		fmt.Println("Error while Connecting to Database")
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
