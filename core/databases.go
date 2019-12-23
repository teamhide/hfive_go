package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() {
	db, err := gorm.Open("postgres", "host=hfive port=5432 user=hfive dbname=hfive password=hfive")
}
