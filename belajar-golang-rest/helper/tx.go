package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	if p := recover(); p != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(p) // re-throw the panic so error handler can catch it
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
