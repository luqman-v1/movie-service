package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrate() (err error) {
	fmt.Println("⇨ Start Migration")

	urlConnection := fmt.Sprintf(
		"%s://%s:%s@tcp(%s:%s)/%s",
		os.Getenv("driver"),
		os.Getenv("db_user"),
		os.Getenv("db_password"),
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("db_database"),
	)

	fmt.Printf("⇨ URL Migration %s ... \n", urlConnection)

	m, err := migrate.New(
		"file://./repo/mysql/migrations",
		urlConnection,
	)
	if err != nil {
		log.Fatal("Migrate: ", err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal("Error Up Migrate: ", err)
		}
		fmt.Printf("⇨ Migration %s ... \n", err)
	} else {
		fmt.Println("⇨ Migration Success")
	}
	return
}
