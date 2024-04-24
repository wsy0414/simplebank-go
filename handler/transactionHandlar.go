package handler

import (
	"context"
	"simplebank/db/sqlc"
)

//ype fn func(*ddd.Queries) error

func HandleTransaction(store sqlc.SqlStore, f func(*sqlc.Queries) error) error {
	tx, err := store.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = f(q)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
