package main

import (
	"fmt"
	"log"
	"github.com/ganeshaditya1/GoBackend/Authsvc/internal/datalayer"
	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	ReadConfig()
	dbusername := viper.Get("database.username").(string)
	dbname := viper.Get("database.dbname").(string)
	dbportno := viper.Get("database.portno").(int64)
	datalayer.InitDB(dbname, dbusername, int(dbportno))



	//err := datalayer.CreateUser("Hello", "Hello@gmail.com", "World", 69)
	//fmt.Println(err)

	user, err := datalayer.LoginUser("Hello", "World")
	fmt.Println(user, err)
}