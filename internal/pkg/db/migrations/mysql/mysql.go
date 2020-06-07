package database

import (
	"database/sql"
	"log"

	// just for eslint
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"

	// just for eslint
	_ "github.com/golang-migrate/migrate/source/file"
)

// Db - The db handler
var Db *sql.DB

// InitDB - Init the connection with database
func InitDB() {
	db, err := sql.Open("mysql", "root:dbpass_new@(localhost:3306)/hackernews")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	log.Println("Connect to database with success!")
	Db = db
}

// Migrate - Uses migrate to migrate
func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
