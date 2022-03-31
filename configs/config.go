package configs

import (
	"os"
	"sync"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port     int
	Database struct {
		Driver   string
		Name     string
		Address  string
		Port     string
		Username string
		Password string
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	// err := godotenv.Load("local.env")

	// if err != nil {
	// 	log.Info(err)
	// }
	// port, errParse := strconv.Atoi(os.Getenv("DB_PORT"))
	// if errParse != nil {
	// 	log.Warn(errParse)
	// }

	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Address = os.Getenv("DB_HOST")
	defaultConfig.Database.Port = os.Getenv("DB_PORT")
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")

	log.Info(defaultConfig)

	return &defaultConfig
}
