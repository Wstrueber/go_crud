package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "billy"
	password = "password"
	dbname   = "tododb"
)

// DB connection
var DB *sql.DB

// InitDB initialize the database
func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to %s:%d", host, port)
	DB = conn
}

func Close() {
	DB.Close()
}
