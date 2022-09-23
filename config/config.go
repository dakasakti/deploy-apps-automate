package config

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Address        string
	Port           string
	DB_Driver      string
	DB_Name        string
	DB_Address     string
	DB_Port        string
	DB_Username    string
	DB_Password    string
	Secret_JWT     string
	Mongo_URI      string
	Mongo_Database string
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
	var defaultConfig AppConfig
	defaultConfig.Address = getEnv("ADDRESS", "http://localhost")
	defaultConfig.Port = getEnv("PORT", "3000")
	defaultConfig.DB_Driver = getEnv("DB_DRIVER", "mysql")
	defaultConfig.DB_Name = getEnv("DB_NAME", "day6")
	defaultConfig.DB_Address = getEnv("DB_ADDRESS", "localhost")
	defaultConfig.DB_Port = getEnv("DB_PORT", "3306")
	defaultConfig.DB_Username = getEnv("DB_USERNAME", "root")
	defaultConfig.DB_Password = getEnv("DB_PASSWORD", "")
	defaultConfig.Secret_JWT = getEnv("SECRET_JWT", "ini_harus_rahasia")
	defaultConfig.Mongo_URI = getEnv("MONGO_URI", "mongodb+srv://<username>:<password>@cluster0-zzart.mongodb.net/test?retryWrites=true&w=majority")
	defaultConfig.Mongo_Database = getEnv("MONGO_DATABASE", "hexagonal")

	fmt.Println(defaultConfig)
	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
