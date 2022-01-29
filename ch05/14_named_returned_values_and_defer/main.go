package main

import (
	"context"
	"database/sql"
)

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			// if everything went well, then commit the changes to the database
			err = tx.Commit()
		}
		if err != nil {
			// if any error occured, don't modify the database
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// use tx to do more database inserts here
	return nil
}
func main() {

}
