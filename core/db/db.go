package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/teamhide/hfive_go/core/configs"
	"log"
)

var db *gorm.DB
var err error

func Init() {
	config := configs.GetConfiguration()
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.DbHost,
			config.DbPort,
			config.DbUser,
			config.DbName,
			config.DbPassword,
		),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
