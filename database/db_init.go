package database

import (
	"database/sql"
	"fmt"

	"../config"
	"../tools"
	_ "github.com/lib/pq"
)
//the result of an function
var (
	worng       = -1
	scuess      = 1
	enable      = 2
	disable     = -2
	repectname  = -20
	repectemail = -30
	othererror  = -99
)

//pointer of database
var db *sql.DB

//reset the config and connect to the database
func init() {
	var err error
	//if don't add 'sllmod=disable' there may export error "pq: SSL is not enabled on the server"
	confstr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Dbhost,
		config.Dbport,
		config.Dbuser,
		config.Dbpassword,
		config.Dbname,
	)
	db, err = sql.Open("postgres", confstr)
	err = db.Ping()
	tools.HandleError("ping of db :", err, -1)
	fmt.Println("Database Connect Scuess!")
}

func Testdb() {
	fmt.Println("ok")
}
