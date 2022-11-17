package dbconnect

import "fmt"

func DeleteCustomer(id int) error {
	db := ConnectToDB()
	stmt, err := db.Prepare("DELETE FROM customers WHERE id=$1")
	if err != nil {
		fmt.Println("stmtPrepareError:", err)
	} else {
		_, err = stmt.Query(id)
		if err != nil {
			fmt.Println("deleteError:", err)
		}
	}
	return err
}
