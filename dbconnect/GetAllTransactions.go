package dbconnect

import (
	"github.com/subramanyam-searce/banking/helpers"
	"github.com/subramanyam-searce/banking/typedefs"
)

func GetAllTransactions() []typedefs.Transaction {
	transactions := []typedefs.Transaction{}
	db := ConnectToDB()

	stmt, err := db.Prepare("SELECT * FROM transactions;")
	helpers.HandleError("dbPrepareError", err)

	rows, err := stmt.Query()
	helpers.HandleError("stmtQueryError", err)

	for rows.Next() {
		transaction := typedefs.Transaction{}
		rows.Scan(&transaction.From, &transaction.To, &transaction.Amount, &transaction.TimeStamp)
		transactions = append(transactions, transaction)
	}

	return transactions
}
