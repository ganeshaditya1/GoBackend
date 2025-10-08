package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type appconfig struct {
	Logfile string `mapstructure:"APP_LOGFILE"`
}

type svcconfig struct {
	AllowedOrigins string `mapstructure:"SVC_ALLOWED_ORIGINS"`
	PortNo int `mapstructure:"SVC_PORTNO"`
}

type authsvcconfig struct {
	PortNo int `mapstructure:"AUTHSVC_PORTNO"`
}

type datasvcconfig struct {
	Shard1PortNo int `mapstructure:"DATASVC1_PORTNO"`
	Shard2PortNo int `mapstructure:"DATASVC2_PORTNO"`
}

type Configuration struct {
	AppConfig appconfig `mapstructure:",squash"`
	SvcConfig svcconfig `mapstructure:",squash"`
	Authsvcconfig authsvcconfig `mapstructure:",squash"`
	Datasvcconfig datasvcconfig `mapstructure:",squash"`
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
}

