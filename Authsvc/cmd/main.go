package main

import (
	"fmt"
	"github.com/ganeshaditya1/GoBackend/Authsvc/internal/api"
	. "github.com/ganeshaditya1/GoBackend/Authsvc/internal/config"
	"github.com/ganeshaditya1/GoBackend/Authsvc/internal/datalayer"
	. "github.com/ganeshaditya1/GoBackend/Authsvc/internal/logging"
	"github.com/ganeshaditya1/GoBackend/Authsvc/internal/middleware"
	"github.com/ganeshaditya1/GoBackend/Authsvc/internal/util"
	"net/http"
	"strings"
)


func start_server() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer(util.NewJWTHandler())

	r := http.NewServeMux()
	
	// get an `http.Handler` that we can use
	allowed_origins := strings.Split(Config.SvcConfig.AllowedOrigins, ",")
	cors := middleware.NewCORSFilter(allowed_origins)
	ssi := api.NewStrictHandler(server, 
							  []api.StrictMiddlewareFunc{middleware.Validator, cors})

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(ssi, r)


	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	fmt.Println(s.ListenAndServe())
}

func main() {
	InitConfig()
	InitLogger()

	dbusername := Config.DBConfig.Username
	dbname := Config.DBConfig.DBname
	dbportno := Config.DBConfig.Portno
	datalayer.InitDB(dbname, dbusername, int(dbportno))


	fmt.Println("Started Server")
	start_server()


	//err := datalayer.CreateUser("Hello", "Hello@gmail.com", "World", 69)
	//fmt.Println(err)

	//user, err := datalayer.LoginUser("Hello", "World")
	//fmt.Println(user, err)
}