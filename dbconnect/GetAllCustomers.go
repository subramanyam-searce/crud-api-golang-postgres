package dbconnect

import (
	"fmt"

	"github.com/subramanyam-searce/banking/typedefs"
)

func GetAllCustomers() []typedefs.Customer {
	db := ConnectToDB()
	customers := []typedefs.Customer{}

	stmt, err := db.Prepare("SELECT * FROM customers")
	if err != nil {
		fmt.Println("dbPrepareError:", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("stmtExecError:", err)
	}
	for rows.Next() {
		customer := typedefs.Customer{}
		rows.Scan(&customer.Name, &customer.Id, &customer.Age, &customer.Email, &customer.Balance, &customer.CreatedAt)
		customers = append(customers, customer)
	}

	return customers
}
