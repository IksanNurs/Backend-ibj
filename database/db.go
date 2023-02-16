package database

import (
	"database/sql"
	"log"
	"os"
)

func Database() *sql.DB {
	db, err := sql.Open(""+os.Getenv("DBMS"), ""+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME")+"?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
