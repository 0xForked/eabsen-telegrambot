package config

import (
	"github.com/spf13/viper"
	"log"
)

// AppConfig represent the data-struct for configuration
type AppConfig struct {
	// another stuff, may be needed by configuration
}

func SetupApplicationConfig() *AppConfig {
	return &AppConfig{}
}

func init() {
	loadConfigFile()
}

const DevMode = "development"

const ProdMode = "production"

func loadConfigFile() {
	log.Println("Load configuration file . . . .")
	// find environment file
	viper.SetConfigFile(`.env`)
	// error handling for specific case
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(".env file not found!, please copy .env.example and paste as .env")
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}
	log.Println("configuration file: ready")
}

func (config AppConfig) GetAppName() string {
	return viper.GetString("APP_NAME")
}

func (config AppConfig) GetAppVersion() string {
	return viper.GetString("APP_VERSION")
}

func (config AppConfig) GetAppEnvironment() string {
	return viper.GetString("APP_ENVIRONMENT")
}

func (config AppConfig) GetDatabaseDriver() string {
	return viper.GetString("DATABASE_DRIVER")
}

func (config AppConfig) GetCacheDriver() string {
	return viper.GetString("CACHE_DRIVER")
}

func (config AppConfig) GetMySqlDsnUrl() string {
	return viper.GetString("MYSQL_DSN_URL")
}

func (config AppConfig) GetRedisDsnUrl() string {
	return viper.GetString("REDIS_DSN_URL")
}

func (config AppConfig) GetSqlitePath() string {
	return viper.GetString("SQLITE_PATH")
}

func (config AppConfig) GetTelegramKey() string {
	return viper.GetString("TELEGRAM_KEY")
}