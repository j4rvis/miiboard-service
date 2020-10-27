package sql

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migration "github.com/golang-migrate/migrate/v4/database/mysql"
)

const MIGRATIONS_PATH = "file://repository/sql/migrations"

func NewDB(host string, user string, password string, databaseName string) *sql.DB {
	config := mysql.NewConfig()
	config.User = user
	config.Passwd = password
	config.Net = "tcp"
	config.DBName = databaseName
	config.Addr = host
	config.ParseTime = true
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXIST %v", databaseName))
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(fmt.Sprintf("USE %v", databaseName))
	if err != nil {
		panic(err.Error())
	}

	driver, err := migration.WithInstance(db, &migration.Config{})
	if err != nil {
		panic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(MIGRATIONS_PATH, "mysql", driver)
	if err != nil {
		panic(err.Error())
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err.Error())
	}

	return db
}
