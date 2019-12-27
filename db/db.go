package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

const (
	DbHost     = "localhost"
	DbPort     = "5432"
	DbUser     = "hfive"
	DbPassword = "hfive"
	DbName     = "hfive"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			DbHost,
			DbPort,
			DbUser,
			DbName,
			DbPassword,
		),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
