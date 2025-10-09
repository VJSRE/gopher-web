package main

import (
	"database/sql"
	"fmt"
	"github.com/VJSRE/lenslocked/models"
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
	dbConfig := models.DefaultPostgresConfig()
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

	/*
		name := "VJSRE"
		email := "Vk@mykoneru.com"

		// Insert records
		record := db.QueryRow(`
			INSERT INTO users (name, email)
			VALUES ($1,$2) RETURNING id; `, name, email)

		var id int
		err = record.Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Record got inserted into the users table; id=", id)


	*/

	/*
		// Insert records into order tables.
		userID := 1

		for i := 1; i <= 5; i++ {
			amount := i * 100
			description := fmt.Sprintf("Fake order  %d", amount)
			_, err := db.Exec(`
			INSERT INTO orders (user_id, amount, description)
			values ($1, $2, $3)`, userID, amount, description)

			if err != nil {
				panic(err)
			}
		}

	*/

	/*
		// Fetch data from Postgresql

		type Order struct {
			ID          int
			userID      int
			Amount      int
			Description string
		}

		var orders []Order
		userID := 1
		rows, err := db.Query(`
		SELECT id, amount, description FROM orders WHERE user_id=$1;`, userID)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var order Order
			order.userID = userID
			err = rows.Scan(&order.ID, &order.Amount, &order.Description)
			if err != nil {
				panic(err)
			}
			orders = append(orders, order)
		}
		if rows.Err() != nil {
			panic(err)
		}
		fmt.Println("Order details ", orders)
	*/
}
