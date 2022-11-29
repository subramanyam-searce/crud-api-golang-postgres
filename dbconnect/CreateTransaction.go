package dbconnect

import (
	"github.com/subramanyam-searce/banking/helpers"
	"github.com/subramanyam-searce/banking/typedefs"
)

func CreateTransaction(t typedefs.Transaction) error {
	db = ConnectToDB()

	stmt, err := db.Prepare("INSERT INTO transactions (from_user, to_user, amount, timestamp) VALUES($1, $2, $3, $4)")
	helpers.HandleError("dbPrepareError", err)

	_, err = stmt.Query(t.From, t.To, t.Amount, t.TimeStamp)
	helpers.HandleError("stmtQueryError", err)

	return err
}
