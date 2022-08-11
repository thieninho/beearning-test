package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:password2605@/traning_user"), &gorm.Config{})
	if err != nil {
		log.Panicf("could not connect")
	}
}
