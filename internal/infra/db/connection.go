package db

import (
	"ecommerce/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.ENV.Db_User_Name,
		config.ENV.Db_Password,
		config.ENV.Db_Host,
		config.ENV.Db_Port,
		config.ENV.Db_Name,
	)
}

func NewConnection() (*sqlx.DB, error) {
	connStr := GetConnectionString()

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
