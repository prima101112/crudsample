package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// InitDBConn for initial database connection
func InitDBConn() {
	var err error
	db, err = sqlx.Connect("postgres", "host=db user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("fatal init db : ", err.Error())
	}
}

// GetCon for getting db connection
func GetCon() *sqlx.DB {
	return db
}
