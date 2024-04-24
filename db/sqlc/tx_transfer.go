package sqlc

import (
	"context"
	"strconv"
)

type TransferTxParam struct {
	FromUserID int
	ToUserID   int
	Currency   string
	Amount     float64
}

type TransferTxResponse struct {
	FromBalance Balance
	ToBalance   Balance
}

func (store *SqlStore) TransferTx(ctx context.Context, param TransferTxParam) (TransferTxResponse, error) {
	var response TransferTxResponse
	err := store.execTx(ctx, func(q *Queries) error {
		err := insertActivity(q, ctx, param.FromUserID, -param.Amount)
		if err != nil {
			return err
		}

		err = insertActivity(q, ctx, param.ToUserID, param.Amount)
		if err != nil {
			return err
		}

		_, err = q.CreateTransferLog(ctx, CreateTransferLogParams{
			FromUserID: int32(param.FromUserID),
			ToUserID:   int32(param.ToUserID),
			Amount:     strconv.FormatFloat(param.Amount, 'e', -1, 64),
		})
		if err != nil {
			return err
		}

		amountStr := strconv.FormatFloat(param.Amount, 'e', -1, 64)
		if param.FromUserID > param.ToUserID {
			response.ToBalance, err = q.AddBalance(ctx, AddBalanceParams{
				UserID:   int32(param.ToUserID),
				Currency: param.Currency,
				Balance:  amountStr,
			})
			if err != nil {
				return err
			}
			response.FromBalance, err = q.SubBalance(ctx, SubBalanceParams{
				UserID:   int32(param.FromUserID),
				Currency: param.Currency,
				Balance:  amountStr,
			})
			if err != nil {
				return err
			}
		} else {
			response.FromBalance, err = q.SubBalance(ctx, SubBalanceParams{
				UserID:   int32(param.FromUserID),
				Currency: param.Currency,
				Balance:  amountStr,
			})
			if err != nil {
				return err
			}
			response.ToBalance, err = q.AddBalance(ctx, AddBalanceParams{
				UserID:   int32(param.ToUserID),
				Currency: param.Currency,
				Balance:  amountStr,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return response, err
}

func insertActivity(q *Queries, ctx context.Context, userId int, amount float64) error {
	_, err := q.CreateActivity(ctx, CreateActivityParams{
		UserID: int32(userId),
		Amount: strconv.FormatFloat(amount, 'e', -1, 64),
	})

	return err
}
