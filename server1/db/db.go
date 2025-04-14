package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	dsn := "root:97531@tcp(127.0.0.1:3306)/tododb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database successfully")

}
