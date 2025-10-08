package main

import (
	"fmt"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/api"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/authsvc"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/datasvc"
	. "github.com/ganeshaditya1/GoBackend/Aggsvc/internal/config"
	. "github.com/ganeshaditya1/GoBackend/Aggsvc/internal/logging"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/middleware"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/requestmodels"
	"net/http"
	"strings"
)


func start_server() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	datasvcShard1 := datasvc.NewServer(Config.Datasvcconfig.Shard1PortNo)
	datasvcShard2 := datasvc.NewServer(Config.Datasvcconfig.Shard2PortNo)
	server := api.NewServer(datasvcShard1, datasvcShard2)

	r := http.NewServeMux()
	
	// get an `http.Handler` that we can use
	allowed_origins := strings.Split(Config.SvcConfig.AllowedOrigins, ",")
	cors := middleware.NewCORSFilter(allowed_origins)

	authserv := authsvc.NewServer(Config.Authsvcconfig.PortNo)
	bearerTokenAuth := middleware.NewBearerTokenAuth(authserv)
	ssi := requestmodels.NewStrictHandler(server, 
							  []requestmodels.StrictMiddlewareFunc{middleware.Validator, cors, bearerTokenAuth})

	// get an `http.Handler` that we can use
	h := requestmodels.HandlerFromMux(ssi, r)


	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("0.0.0.0:%d", Config.SvcConfig.PortNo),
	}

	// And we serve HTTP until the world ends.
	fmt.Println(s.ListenAndServe())
}

func main() {
	InitConfig()
	InitLogger()


	fmt.Println("Started Server")
	start_server()
}
