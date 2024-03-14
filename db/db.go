package db

import (
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connConfig, err := pgx.ParseConfig(os.Getenv("PG_URI"))
	if err != nil {
		log.Fatal(err)
		return
	}
	pgxdb := stdlib.OpenDB(*connConfig)
	DB = sqlx.NewDb(pgxdb, "pgx")
	DB.SetMaxIdleConns(4)
	DB.SetMaxOpenConns(8)
	DB.SetConnMaxLifetime(time.Duration(30) * time.Minute)

	createTables()

	log.Println("DB connected !")
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
	    email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Println("Could not create users table")
		log.Fatal(err)
		return
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
    	id SERIAL PRIMARY KEY,
    	name TEXT NOT NULL,
    	description TEXT NOT NULL,
    	location TEXT NOT NULL,
    	dateTime TIMESTAMP NOT NULL,
    	user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Println("Could not create events table")
		log.Fatal(err)
		return
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	    id SERIAL PRIMARY KEY,
	    event_id INTEGER,
	    user_id INTEGER,
	    FOREIGN KEY(event_id) REFERENCES events(id),
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		log.Println("Could not create registrations table")
		log.Fatal(err)
		return
	}
}
