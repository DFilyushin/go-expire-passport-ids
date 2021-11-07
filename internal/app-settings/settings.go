package app_settings

import (
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strconv"
)

type SettingKey string

type Settings struct {
	ServerHost     string `app:"APP_SERVER_HOST"`
	ServerPort     int    `app:"APP_SERVER_PORT"`
	UrlMVD         string `app:"APP_PUBLIC_URL"`
	MatrixDataFile string `app:"APP_MATRIX_FILE"`
}

const (
	DefaultDotEnvFile     = ".env"
	ApplicationTag        = "app"
	DefaultFilePermission = 0644
)

func (s *Settings) loadFromDotEnv() error {
	/*Загрузить переменные окружения из файла .env*/
	return godotenv.Load(DefaultDotEnvFile)
}

func (s *Settings) LoadSettings() {
	/*Прочитать настройки из переменных окружения*/
	file, err := os.OpenFile(DefaultDotEnvFile, os.O_RDONLY, DefaultFilePermission)
	defer file.Close()
	if !os.IsNotExist(err) {
		err = s.loadFromDotEnv()
	}

	v := reflect.ValueOf(s)
	vPtr := reflect.Indirect(v)

	for i := 0; i < vPtr.NumField(); i++ {
		tagName := vPtr.Type().Field(i).Tag
		envName := tagName.Get(ApplicationTag)

		switch vPtr.Field(i).Kind() {
		case reflect.Int:
			intValue, _ := strconv.Atoi(os.Getenv(envName))
			vPtr.Field(i).SetInt(int64(intValue))
		case reflect.String:
			value := os.Getenv(envName)
			vPtr.Field(i).SetString(value)
		case reflect.Bool:
			value := os.Getenv(envName)
			boolValue, _ := strconv.ParseBool(value)
			vPtr.Field(i).SetBool(boolValue)
		}
	}
}
