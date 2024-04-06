package test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"simplebank/db/sqlc"
	"simplebank/handler"
)

// TestCreateAccount
func TestCreateAccount(t *testing.T) {
	createAccountParam := sqlc.CreateAccountParams{
		Name:     "Ivan",
		Password: "asdfadsfasdf",
		Balance:  "0.0",
		Currency: "TWD",
	}

	account, err := testQuery.CreateAccount(context.Background(), createAccountParam)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createAccountParam.Name, account.Name)
	require.NotEmpty(t, account.ID)
}

// transaction test
func TestTx(t *testing.T) {
	err := handler.HandleTransaction(db, func(q sqlc.Queries) error {
		createAccountParam := sqlc.CreateAccountParams{
			Name:     "tttt",
			Password: "asdfadsfasdf",
			Balance:  "0.0",
			Currency: "TWD",
		}

		account, err := q.CreateAccount(context.Background(), createAccountParam)
		if err != nil {
			return err
		}

		updateAccountParam := sqlc.UpdateAccountParams{
			Balance: "0.2",
			ID:      account.ID,
		}

		err = q.UpdateAccount(context.Background(), updateAccountParam)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Println("error")
	}

}
