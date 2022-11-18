package dbconnect

import (
	"fmt"

	"github.com/subramanyam-searce/banking/typedefs"
)

func UpdateCustomer(id int, key string, val any) *typedefs.Customer {
	db := ConnectToDB()

	stmt, err := db.Prepare("UPDATE customers SET " + key + "=$1 WHERE id=$2;")
	if err != nil {
		fmt.Println("dbPrepareError:", err)
	}
	defer stmt.Close()

	_, err = stmt.Query(val, id)
	if err != nil {
		fmt.Println("stmtQueryError:", err)
	}

	return GetCustomer(id)

}
