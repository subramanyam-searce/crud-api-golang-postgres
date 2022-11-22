package dbconnect

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/subramanyam-searce/banking/helpers"
)

var db *sql.DB

func init() {
	connectionString := "user=service-banking dbname=banking password=crudapi host=localhost sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connectionString)
	helpers.HandleError("sqlOpenError", err)
}

func ConnectToDB() *sql.DB {
	return db
}
