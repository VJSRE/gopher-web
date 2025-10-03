package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (conf PostgresConfig) Connect() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.Host, conf.Port, conf.User, conf.Password, conf.Database, conf.SSLMode)
}

func main() {
	dbConfig := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", dbConfig.Connect())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to db")

	// Create tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		      id serial PRIMARY KEY,
		      name text, 
		      email text unique not null
		);
		CREATE TABLE IF NOT EXISTS orders (
		       id serial PRIMARY KEY,
		       user_id integer not null,
		       amount INT,
		       description text
		)
 	`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Tables are created")

	name := "Vijay Koneru"
	email := "Vijay.Koneru@mykoneru.com"

	// Insert records
	_, err = db.Exec(`
		INSERT INTO users (name, email)
		VALUES ($1,$2);`, name, email)

	if err != nil {
		panic(err)
	}
	fmt.Println("Record got inserted into the users table")

}
