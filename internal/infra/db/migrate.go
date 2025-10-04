package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func RunMigrations(db *sqlx.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	log.Println("Running migrations...")
	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Printf("Applied %d migrations!\n", n)
}
