package database

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var Db *sql.DB

func InitDB() {
	log.Print("Initilazing DB Connection")
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/news")
	if err != nil {
		log.Panic(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
		return
	}
	Db = db
	if Db != nil {
		log.Print("Initialized DB Connection")
	}
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	driver, err := postgres.WithInstance(Db, new(postgres.Config))
	if err != nil {
		log.Fatal(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/postgres",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
