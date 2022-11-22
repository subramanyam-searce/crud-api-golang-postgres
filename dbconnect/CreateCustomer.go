package dbconnect

import (
	"time"

	"github.com/subramanyam-searce/banking/helpers"
	"github.com/subramanyam-searce/banking/typedefs"
)

func CreateCustomer(data *typedefs.Customer) error {
	db := ConnectToDB()

	stmt, err := db.Prepare("INSERT INTO customers (name, id, age, email, balance, created_at) VALUES($1, $2, $3, $4, $5, $6);")
	helpers.HandleError("dbPrepareError", err)

	defer stmt.Close()

	data.CreatedAt = time.Now().Format(time.RFC3339)
	_, err = stmt.Query(data.Name, data.Id, data.Age, data.Email, data.Balance, data.CreatedAt)
	helpers.HandleError("stmtExecError", err)

	return err
}
