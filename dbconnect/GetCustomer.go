package dbconnect

import (
	"github.com/subramanyam-searce/banking/helpers"
	"github.com/subramanyam-searce/banking/typedefs"
)

func GetCustomer(id string) *typedefs.Customer {
	db := ConnectToDB()
	customer := typedefs.Customer{}

	stmt, err := db.Prepare("SELECT * FROM customers WHERE id=$1")
	helpers.HandleError("dbPrepareError", err)
	defer stmt.Close()

	rows, err := stmt.Query(id)
	helpers.HandleError("stmtExecError", err)

	ok := rows.Next()
	if ok {
		rows.Scan(&customer.Name, &customer.Id, &customer.Age, &customer.Email, &customer.Balance, &customer.CreatedAt)
		return &customer
	} else {
		return nil
	}
}
