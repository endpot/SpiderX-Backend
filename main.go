package main

import (
	"context"
	. "github.com/endpot/SpiderX-Backend/app/infra/config"
	. "github.com/endpot/SpiderX-Backend/app/infra/db"
	. "github.com/endpot/SpiderX-Backend/app/router"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title SpiderX API
// @version 0.0.1
// @description SpiderX is a private tracker server.

// @contact.name EndPot
// @contact.url https://github.com/endpot

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host spiderx.endpot.com
// @BasePath /

// @schemes https

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// Print Logo
	printLogo()

	// Init
	InitConfig()
	InitEloquent()

	// Close DB Connection
	defer Eloquent.Close()

	// Construct server
	r := InitRouter()
	srv := &http.Server{
		Addr:    Config.GetString("APP_LISTEN_ADDR"),
		Handler: r,
	}

	go func() {
		// Listen and serve
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Err: %s\n", err)
		}
	}()

	log.Println("Server URL: http://" + Config.GetString("APP_LISTEN_ADDR") + "/")
	log.Println("API Doc URL: http://" + Config.GetString("APP_LISTEN_ADDR") + "/_doc/index.html")
	log.Println("Enter Control + C Shutdown Server")

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func printLogo() {
	log.Println(
		`
      ___           ___                    _____          ___           ___           ___
     /  /\         /  /\      ___         /  /::\        /  /\         /  /\         /__/|
    /  /:/_       /  /::\    /  /\       /  /:/\:\      /  /:/_       /  /::\       |  |:|
   /  /:/ /\     /  /:/\:\  /  /:/      /  /:/  \:\    /  /:/ /\     /  /:/\:\      |  |:|
  /  /:/ /::\   /  /:/~/:/ /__/::\     /__/:/ \__\:|  /  /:/ /:/_   /  /:/~/:/    __|__|:|
 /__/:/ /:/\:\ /__/:/ /:/  \__\/\:\__  \  \:\ /  /:/ /__/:/ /:/ /\ /__/:/ /:/___ /__/::::\____
 \  \:\/:/~/:/ \  \:\/:/      \  \:\/\  \  \:\  /:/  \  \:\/:/ /:/ \  \:\/:::::/    ~\~~\::::/
  \  \::/ /:/   \  \::/        \__\::/   \  \:\/:/    \  \::/ /:/   \  \::/~~~~      |~~|:|~~
   \__\/ /:/     \  \:\        /__/:/     \  \::/      \  \:\/:/     \  \:\          |  |:|
     /__/:/       \  \:\       \__\/       \__\/        \  \::/       \  \:\         |  |:|
     \__\/         \__\/                                 \__\/         \__\/         |__|/
		`,
	)
}
