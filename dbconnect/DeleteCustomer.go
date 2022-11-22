package dbconnect

import (
	"github.com/subramanyam-searce/banking/helpers"
)

func DeleteCustomer(id string) error {
	db := ConnectToDB()

	stmt, err := db.Prepare("DELETE FROM customers WHERE id=$1")
	helpers.HandleError("dbPrepareError", err)

	_, err = stmt.Query(id)
	helpers.HandleError("stmtQueryError", err)

	return err
}
