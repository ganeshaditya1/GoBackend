package config

import (
	"fmt"
	"os"

    "github.com/spf13/pflag"
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
	PortNo int `mapstructure:"SVC_PORTNO"`
}

type Configuration struct {
	DBConfig dbconfig `mapstructure:",squash"`
	AppConfig appconfig `mapstructure:",squash"`
	SvcConfig svcconfig `mapstructure:",squash"`
}

var Config *Configuration

func InitConfig() {
	var configuration *Configuration

	// Command line arguments
    pflag.String("config", "", "config file")
    pflag.Parse()

	// bind pflag to viper
    viper.BindPFlags(pflag.CommandLine)

    viper.SetConfigFile(viper.GetString("config"))
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
