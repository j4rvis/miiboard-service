package sql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migration "github.com/golang-migrate/migrate/v4/database/mysql"

	// Used for sql driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationsPath = "file://repository/sql/migrations"

// NewDB ...
func NewDB(host string, user string, password string, databaseName string) *sql.DB {
	config := mysql.NewConfig()
	config.User = user
	config.Passwd = password
	config.Net = "tcp"
	// config.DBName = databaseName
	config.Addr = host
	config.ParseTime = true

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("USE " + databaseName)
	if err != nil {
		panic(err.Error())
	}

	driver, err := migration.WithInstance(db, &migration.Config{})
	if err != nil {
		panic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "mysql", driver)
	if err != nil {
		panic(err.Error())
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err.Error())
	}

	return db
}
