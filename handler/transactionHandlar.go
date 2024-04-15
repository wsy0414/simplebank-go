package handler

import (
	"context"
	"database/sql"
	"simplebank/db/sqlc"
)

//ype fn func(*ddd.Queries) error

func HandleTransaction(db *sql.DB, f func(sqlc.Queries) error) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = f(*q)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
