package db

import (
	"database/sql"
	"fmt"
	"github.com/endpot/SpiderX-Backend/app/infra/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	log.Print("DB Initializing...")

	// Generate connection string
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		config.Config.GetString("DB_USER"),
		config.Config.GetString("DB_PASS"),
		config.Config.GetString("DB_HOST"),
		config.Config.GetString("DB_PORT"),
		config.Config.GetString("DB_NAME"),
	)

	// Init DB
	var err error
	DB, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	log.Println("DB Initialized!")
}
