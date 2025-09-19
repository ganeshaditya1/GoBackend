package datalayer

import (
    "fmt"
    "log"
	"os"
    
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(dbname string, dbusername string, 
						  dbportno int) {
	authsvc_pword := os.Getenv("AUTHSVC_PWORD")
	if authsvc_pword == "" {
		log.Fatalln("Postgres user Authsvc's password not specified.")
		return
	}

	dbobj, err := sqlx.Connect("postgres", 
			fmt.Sprintf("user=%s dbname=%s password=%s port=%d sslmode=disable", 
						dbusername, dbname, authsvc_pword, dbportno))
    if err != nil {
        log.Fatalln(err)
		return
    }
	db = dbobj
}


func GetDB() *sqlx.DB {
	if db == nil {
		log.Fatalln("DB not initialized yet.")
		return nil
	}
	return db
}