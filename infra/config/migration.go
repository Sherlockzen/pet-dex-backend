package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunUpMigrations(dbconnection *sql.DB) {
	driver, _ := mysql.WithInstance(dbconnection, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://"+env.MIGRATIONS_PATH,
		"mysql",
		driver,
	)
	fmt.Println("Executing UP migrations")
	fmt.Println(env.MIGRATIONS_PATH)
	error := m.Up()

	if error != nil && error != migrate.ErrNoChange {
		fmt.Printf("Error trying execute migrations UP: %v \n", error)
		panic(error)
	}

	fmt.Println("Migrations has been executed")
}

//func RunDownMigrations(dbconnection *sql.DB) {
//	driver, _ := mysql.WithInstance(dbconnection, &mysql.Config{})
//	m, _ := migrate.NewWithDatabaseInstance(
//		"file://"+env.MIGRATIONS_PATH,
//		"mysql",
//		driver,
//	)
//	fmt.Println("Executing DOWN migrations")
//}
