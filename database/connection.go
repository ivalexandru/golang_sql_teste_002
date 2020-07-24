package database

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE person0 (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place0 (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func ConnectDatabase() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=postgres password=pern123 dbname=golang_sqlx_test002 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)

	// tx := db.MustBegin()
}
