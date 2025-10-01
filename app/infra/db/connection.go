package db

import (
	"ecommerce/app/config"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)


func GetConnectionString() string {
	return config.ENV.Db_User_Name + ":" + config.ENV.Db_Password + "@localhost:5432/ecommerce?sslmode=disable"
}


func NewConnection() (*sqlx.DB, error) {
	connStr := GetConnectionString()
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
		os.Exit(1);
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
		os.Exit(1);
	}

	return db, nil
}