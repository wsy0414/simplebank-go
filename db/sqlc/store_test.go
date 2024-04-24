package sqlc

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	account1, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   1,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   3,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	n := 20
	amount := 1.0
	errs := make(chan error)
	results := make(chan TransferTxResponse)

	for i := 1; i <= n; i++ {
		go func() {
			result, err := testStore.TransferTx(context.Background(), TransferTxParam{
				FromUserID: int(account1.UserID),
				ToUserID:   int(account2.UserID),
				Currency:   account1.Currency,
				Amount:     amount,
			})

			errs <- err
			results <- result
		}()
	}

	for j := 1; j <= n; j++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.NotEmpty(t, result.FromBalance)
		require.NotEmpty(t, result.ToBalance)

		fmt.Println("after transfer", result.FromBalance.Balance, result.ToBalance.Balance)
	}

	updateAccount1, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   1,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, updateAccount1)
	updateAccount2, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   3,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, updateAccount2)

	balance1, err := strconv.ParseFloat(account1.Balance, 64)
	require.NoError(t, err)
	balance2, err := strconv.ParseFloat(account2.Balance, 64)
	require.NoError(t, err)

	updateBalance1, err := strconv.ParseFloat(updateAccount1.Balance, 64)
	require.NoError(t, err)
	updateBalance2, err := strconv.ParseFloat(updateAccount2.Balance, 64)
	require.NoError(t, err)

	require.Equal(t, balance1-float64(n)*amount, updateBalance1)
	require.Equal(t, balance2+float64(n)*amount, updateBalance2)
}

func TestTransferTxDeadLock(t *testing.T) {
	account1, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   1,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   3,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	n := 20
	amount := 1.0
	errs := make(chan error)
	results := make(chan TransferTxResponse)

	for i := 0; i < n; i++ {
		fromAccountID := account1.UserID
		toAccountID := account2.UserID

		if i%2 == 1 {
			fromAccountID = account2.UserID
			toAccountID = account1.UserID
		}

		go func() {
			result, err := testStore.TransferTx(context.Background(), TransferTxParam{
				FromUserID: int(fromAccountID),
				ToUserID:   int(toAccountID),
				Currency:   account1.Currency,
				Amount:     amount,
			})

			errs <- err
			results <- result
		}()
	}

	for j := 1; j <= n; j++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.NotEmpty(t, result.FromBalance)
		require.NotEmpty(t, result.ToBalance)

		fmt.Println("after transfer", result.FromBalance.Balance, result.ToBalance.Balance)
	}

	updateAccount1, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   1,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, updateAccount1)
	updateAccount2, err := testStore.GetSpecifyCurrencyBalanceByUser(context.Background(), GetSpecifyCurrencyBalanceByUserParams{
		UserID:   3,
		Currency: "TWD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, updateAccount2)

	balance1, err := strconv.ParseFloat(account1.Balance, 64)
	require.NoError(t, err)
	balance2, err := strconv.ParseFloat(account2.Balance, 64)
	require.NoError(t, err)

	updateBalance1, err := strconv.ParseFloat(updateAccount1.Balance, 64)
	require.NoError(t, err)
	updateBalance2, err := strconv.ParseFloat(updateAccount2.Balance, 64)
	require.NoError(t, err)

	require.Equal(t, balance1, updateBalance1)
	require.Equal(t, balance2, updateBalance2)
}
