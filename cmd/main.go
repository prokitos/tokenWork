package main

import (
	"mymod/internal/app"
	"mymod/internal/config"
	"mymod/internal/database"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var CONFIG_TYPE string = "local"

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debug("log is loaded")

	var cfg config.MainConfig
	cfg.ConfigMustLoad(CONFIG_TYPE)
	log.Debug("config is loaded")

	databaseLoad(cfg)
	log.Debug("databases is loaded")

	var application app.App
	go application.NewServer(cfg.Server.Port)
	log.Debug("server is loaded")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop()

}

// загрузка базы данных и настройка дао.
func databaseLoad(cfg config.MainConfig) {
	database.GlobalPostgres.Run(cfg)
	database.GlobalPostgres.DaoToken.New()
	database.GlobalPostgres.DaoAuth.New()
}
