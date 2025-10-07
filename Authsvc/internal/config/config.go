package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type dbconfig struct {
	DBname string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"DB_USERNAME"`
	Portno int `mapstructure:"DB_PORTNO"`
}

type appconfig struct {
	Logfile string `mapstructure:"APP_LOGFILE"`
}

type svcconfig struct {
	AllowedOrigins string `mapstructure:"SVC_ALLOWED_ORIGINS"`
}

type Configuration struct {
	DBConfig dbconfig `mapstructure:",squash"`
	AppConfig appconfig `mapstructure:",squash"`
	SvcConfig svcconfig `mapstructure:",squash"`
}

var Config *Configuration

func InitConfig() {
	var configuration *Configuration
    viper.SetConfigFile("../.env")
    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading config file, %s", err)
		os.Exit(1)
    }

    err := viper.Unmarshal(&configuration)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to decode into struct, %v", err)
		os.Exit(1)
    }

    Config = configuration
	fmt.Println(Config.DBConfig.DBname)
}
