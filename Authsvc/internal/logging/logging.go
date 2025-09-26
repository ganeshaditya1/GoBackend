package logging

import (
	"log/slog"
	"os"
	. "github.com/ganeshaditya1/GoBackend/Authsvc/internal/config"
)

func InitLogger() {
	fhandle, err := os.OpenFile(Config.AppConfig.Logfile, 
						   os.O_WRONLY | os.O_APPEND | os.O_CREATE,
						0770)
    if err != nil {
        panic(err)
    }
	logger := slog.New(slog.NewJSONHandler(fhandle, nil))
 	slog.SetDefault(logger)

	slog.Info("Logger initialized")
}