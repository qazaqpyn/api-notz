package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/qazaqpyn/api-notz/pkg/handler"
	"github.com/qazaqpyn/api-notz/pkg/repository"
	"github.com/qazaqpyn/api-notz/pkg/service"
	apinotz "github.com/qazaqpyn/api-notz/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	initConfig()

	db, err := repository.NewPostgresDB(viper.GetString("db.url"))
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	repos := repository.NewRepository(db)

	// Initialize services
	services := service.NewService(repos)

	// Initialize handlers
	handlers := handler.NewHandler(services)

	srv := new(apinotz.Server)

	// Graceful shutdown
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err)
		}
	}()

	logrus.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error reading config file: %s", err)
	}
}
