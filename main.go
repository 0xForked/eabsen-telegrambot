package main

import (
	"github.com/aasumitro/svc-tgbotgo/config"
	"github.com/aasumitro/svc-tgbotgo/src/handlers"
	"github.com/aasumitro/svc-tgbotgo/src/repositories"
	"log"
	"runtime"
)

var appConfig *config.AppConfig

func init() {
	log.Println("=====================================")
	log.Println("|   Trying to run the application   |")
	log.Println("=====================================")
	appConfig = config.SetupApplicationConfig()
	appConfig.InitDatabaseConnection()
	appConfig.InitTelegramBot()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	db := appConfig.GetDatabaseConnection()
	bot := appConfig.GetTelegramBot()

	userRepo := repositories.NewUserRepository(db)

	handlers.NewBotHandler(bot, userRepo).Watch()
}