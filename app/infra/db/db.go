package db

import (
	"fmt"
	. "github.com/endpot/SpiderX-Backend/app/infra/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Eloquent *gorm.DB

func InitEloquent() {
	log.Print("DB Initializing...")

	// Generate connection string
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		Config.GetString("DB_USER"),
		Config.GetString("DB_PASS"),
		Config.GetString("DB_HOST"),
		Config.GetString("DB_PORT"),
		Config.GetString("DB_NAME"),
	)

	// Init Eloquent
	var err error
	Eloquent, err = gorm.Open("mysql", conn)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	if Eloquent.Error != nil {
		log.Fatalf("Err: %v", Eloquent.Error)
	}

	// Print Eloquent log
	Eloquent.LogMode(true)

	log.Println("DB Initialized!")
}
