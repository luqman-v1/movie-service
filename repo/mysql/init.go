package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var sqlxDB *sqlx.DB

func Conn() {
	usernameAndPassword := fmt.Sprint(os.Getenv("db_user")) + ":" + fmt.Sprint(os.Getenv("db_password"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("db_database")) + "?charset=utf8&parseTime=true&loc=UTC"
	log.Println("url", urlConnection)
	db, err := sql.Open(os.Getenv("driver"), urlConnection)
	if err != nil {
		log.Fatalf("⇨ %s Data source %s:%s , Failed : %s \n", os.Getenv("driver"), os.Getenv("db_host"), os.Getenv("db_port"), err.Error())
	}

	dbx := sqlx.NewDb(db, os.Getenv("driver"))

	fmt.Printf("⇨ Connect MYSQL to Server %s ... \n", hostName)

	sqlxDB = dbx
}
