package sqlc

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	TransferTx(context.Context, TransferTxParam) (TransferTxResponse, error)
}

type SqlStore struct {
	DB *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SqlStore{
		DB:      db,
		Queries: New(db),
	}
}
