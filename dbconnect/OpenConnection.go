package dbconnect

import (
	"database/sql"
	"fmt"
	"runtime"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connectionString := "user=service-banking dbname=banking password=crudapi host=localhost sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func ConnectToDB() *sql.DB {
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	return db
}
