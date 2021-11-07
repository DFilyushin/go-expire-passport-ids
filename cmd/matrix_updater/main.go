package main

import (
	app_settings "expired-passport-checker/internal/app-settings"
	"expired-passport-checker/internal/service"
	"expired-passport-checker/pkg/fget"
	_ "expired-passport-checker/pkg/fget"
	fintech_logger "expired-passport-checker/pkg/fintech-logger"
	"fmt"
	"os"
)

func main() {
	logger := fintech_logger.NewFintechLogger(fintech_logger.JsonFormatter)

	settings := app_settings.Settings{}
	settings.LoadSettings()

	logger.Info("Matrix updater started")
	client := new(fget.FGet)
	zipFile, err := client.DownloadFile(settings.UrlMVD)
	if err != nil {
		logger.Fatalf("Error downloading file. %s", err)
	}

	logger.Info(fmt.Sprintf("File saved at %s", zipFile))
	passportService := service.NewPassportIdService()

	//load from bz-archive
	countLoaded, err := passportService.LoadDataFromArchive(zipFile)
	if err != nil {
		_ = os.Remove(zipFile)
		logger.Fatalf("Error loading file in matrix. %s", err)
	}
	logger.Info(fmt.Sprintf("Loaded %d lines", countLoaded))

	//save to matrix file
	err = passportService.SaveDataFile(settings.MatrixDataFile)
	if err != nil {
		logger.Fatalf("Error saving matrix file. %s", err)
	}
	logger.Info("Matrix updater finished")
}
