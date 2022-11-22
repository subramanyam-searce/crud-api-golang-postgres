package dbconnect

import (
	"github.com/subramanyam-searce/banking/helpers"
	"github.com/subramanyam-searce/banking/typedefs"
)

func UpdateCustomer(id string, key string, val any) *typedefs.Customer {
	db := ConnectToDB()

	stmt, err := db.Prepare("UPDATE customers SET " + key + "=$1 WHERE id=$2;")
	helpers.HandleError("dbPrepareError", err)

	defer stmt.Close()

	_, err = stmt.Query(val, id)
	helpers.HandleError("stmtExecError", err)

	return GetCustomer(id)

}
