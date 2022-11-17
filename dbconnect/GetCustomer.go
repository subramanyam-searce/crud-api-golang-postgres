package dbconnect

import (
	"fmt"

	"github.com/subramanyam-searce/banking/typedefs"
)

func GetCustomer(id int) *typedefs.Customer {
	db := ConnectToDB()
	customer := typedefs.Customer{}

	stmt, err := db.Prepare("SELECT * FROM customers WHERE id=$1")
	if err != nil {
		fmt.Println("dbPrepareError:", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println("stmtExecError:", err)
	}
	ok := rows.Next()
	if ok {
		rows.Scan(&customer.Name, &customer.Id, &customer.Age, &customer.Email, &customer.Balance, &customer.CreatedAt)
		return &customer
	} else {
		return nil
	}
}
