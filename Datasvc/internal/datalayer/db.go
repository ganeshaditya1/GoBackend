package datalayer

import (
    "fmt"
    "log/slog"
	"os"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(dbname string, dbusername string, 
						  dbportno int) {
	datasvc_pword := os.Getenv("DATASVC_PWORD")
	if datasvc_pword == "" {
		slog.Error("Postgres user Datasvc's password not specified.")
		os.Exit(1)
		return
	}

	dbobj, err := sqlx.Connect("postgres", 
			fmt.Sprintf("user=%s dbname=%s password=%s port=%d sslmode=disable", 
						dbusername, dbname, datasvc_pword, dbportno))
    if err != nil {
        slog.Error("Cannot connect to the DB.", "error", err)
		os.Exit(1)
		return
    }
	db = dbobj
}


func GetDB() *sqlx.DB {
	if db == nil {
		slog.Error("DB not initialized yet.")
		return nil
	}
	return db
}

func SetDB(dbobj *sqlx.DB) {
	db = dbobj
}