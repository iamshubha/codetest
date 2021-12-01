package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	hostname     = "localhost"
	host_port    = 5432
	username     = "postgres"
	password     = "1234"
	databasename = "postgres"
)

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, databasename)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		panic(err)

	}
	return db
}

func InitDB(db *sql.DB) {
	sqlQueryforlogintable := `
	CREATE TABLE IF NOT EXISTS filesys (
		id SERIAL PRIMARY KEY,
		filePath TEXT NOT NULL,
		fileName TEXT NOT NULL,
		folderPath TEXT NOT NULL
	);	
	`
	_, err := db.Exec(sqlQueryforlogintable)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

}
