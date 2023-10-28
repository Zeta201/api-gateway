package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("DB_PORT")
	// 		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		log.Fatal(err.Error())
		// log.
		// log.Fatal("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
