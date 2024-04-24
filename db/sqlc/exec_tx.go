package sqlc

import "context"

func (store *SqlStore) execTx(ctx context.Context, f func(*Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = f(q)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
