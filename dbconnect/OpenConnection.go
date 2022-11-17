package dbconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	connectionString := "user=service-banking dbname=banking password=crudapi host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return db
}
