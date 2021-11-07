package main

import (
	"context"
	"expired-passport-checker/internal/api"
	"expired-passport-checker/internal/app-settings"
	"expired-passport-checker/internal/service"
	"expired-passport-checker/pkg/fintech-logger"
	"os"
	"os/signal"
	"syscall"
)

// @title Expired passport checker
// @version 1.0
// @description API server for checking expired passport

// @contact.name Митя Филюшин
// @contact.email dvfilyushin@fintechiq.ru

// @BasePatch /api
func main() {
	logger := fintech_logger.NewFintechLogger(fintech_logger.JsonFormatter)

	settings := app_settings.Settings{}
	settings.LoadSettings()

	passportService := new(service.PassportIdService)
	err := passportService.Init(settings.MatrixDataFile)
	if err != nil {
		logger.Fatalf("Error loading matrix data file %s", err.Error())
	}
	logger.Info("Service initialized")

	server := new(api.HttpServer)
	server.Initialize(settings.ServerHost, settings.ServerPort, passportService)

	go func() {
		if err := server.RunServer(); err != nil {
			logger.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logger.Info("Http server started at %s:%d", server.Host, server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Server shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error("Error occurred on server shutting down: %s", err.Error())
	}
}
