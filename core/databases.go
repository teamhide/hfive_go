package core

import (
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

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("postgres", "host="+DbHost+" port="+DbPort+" user="+DbUser+" dbname="+DbName+" password="+DbPassword+" sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
}
