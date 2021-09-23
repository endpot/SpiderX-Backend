package oss

import (
	"github.com/endpot/SpiderX-Backend/app/infra/config"
	"github.com/minio/minio-go/v6"
	"log"
)

var Client *minio.Client

func InitClient() {
	log.Print("OSS Initializing...")

	// Init Client
	var err error
	Client, err = minio.New(
		config.Config.GetString("OSS_ENDPOINT"),
		config.Config.GetString("OSS_KEY"),
		config.Config.GetString("OSS_SECRET"),
		config.Config.GetBool("OSS_USE_SSL"),
	)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	log.Println("OSS Initialized!")
}
