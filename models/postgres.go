package models

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// This method will open the connection with the provided postgres database. The caller has to ensure to eventually close the connection.
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.Connect())
	if err != nil {
		return nil, fmt.Errorf("error while opening postgres connection: %v", err)
	}
	return db, nil
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

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
