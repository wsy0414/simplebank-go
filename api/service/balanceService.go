package service

import (
	"context"
	"database/sql"
	"errors"
	"simplebank/api/model"
	"simplebank/customError"
	"simplebank/db/sqlc"
	"strconv"

	"github.com/lib/pq"
)

type BalanceService interface {
	CreateBalance(context.Context, model.CreateBalanceRequest) (model.CreateBalanceResponse, error)
	GetBalance(context.Context, string) (model.GetBalanceResponse, error)
	ListBalance(context.Context) ([]model.GetBalanceResponse, error)
	Deposite(context.Context, model.DepositeRequestParam) (model.DepositeResponse, error)
	Withdraw(context.Context, model.WithdrawRequestParam) (model.WithdrawResponse, error)
	Transfer(context.Context, model.TransferRequestParam) (model.TransferResponse, error)
	GetActivity(context.Context) ([]model.ListActivityResponse, error)
}

type balanceService struct {
	repo sqlc.Store
}

func NewBalanceService(store sqlc.Store) BalanceService {
	return &balanceService{
		repo: store,
	}
}

func (service *balanceService) CreateBalance(
	ctx context.Context, param model.CreateBalanceRequest,
) (response model.CreateBalanceResponse, err error) {
	userId := getUserId(ctx)

	createBalanceParams := sqlc.CreateBalanceParams{
		UserID:   int32(userId),
		Currency: param.Currency,
		Balance:  strconv.FormatFloat(param.Amount, 'e', -1, 64),
	}
	balance, err := service.repo.CreateBalance(ctx, createBalanceParams)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			// unique valid
			if pgErr.Code.Name() == "unique_violation" {
				return response, customError.NewBadRequestError(errors.New("the Balance already be created"))
			}
		}
		return response, customError.NewInternalError(err)
	}

	response = model.CreateBalanceResponse{
		BalanceResponse: model.BalanceResponse{
			Currency: balance.Currency,
			Balance:  balance.Balance,
		},
	}

	return
}

func (service *balanceService) GetBalance(ctx context.Context, currency string) (model.GetBalanceResponse, error) {
	userId := getUserId(ctx)

	balance, err := service.repo.GetSpecifyCurrencyBalanceByUser(ctx,
		sqlc.GetSpecifyCurrencyBalanceByUserParams{
			UserID:   int32(userId),
			Currency: currency,
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.GetBalanceResponse{}, customError.NewBadRequestError(errors.New("the Balance is not yet be created"))
		}
		return model.GetBalanceResponse{}, customError.NewInternalError(err)
	}

	return model.GetBalanceResponse{
		BalanceResponse: model.BalanceResponse{
			Currency: balance.Currency,
			Balance:  balance.Balance,
		},
	}, nil
}

func (service *balanceService) ListBalance(ctx context.Context) (response []model.GetBalanceResponse, err error) {
	userId := getUserId(ctx)
	response = make([]model.GetBalanceResponse, 0)

	balanceList, err := service.repo.GetBalanceByUser(ctx, int32(userId))
	if err != nil {
		return response, customError.NewInternalError(err)
	}

	for _, balance := range balanceList {
		response = append(response, model.GetBalanceResponse{
			BalanceResponse: model.BalanceResponse{
				Currency: balance.Currency,
				Balance:  balance.Balance,
			},
		})
	}

	return
}

func (service *balanceService) Deposite(
	ctx context.Context,
	param model.DepositeRequestParam) (
	response model.DepositeResponse,
	err error,
) {
	userId := getUserId(ctx)
	_, err = service.repo.GetSpecifyCurrencyBalanceByUser(ctx, sqlc.GetSpecifyCurrencyBalanceByUserParams{
		UserID:   int32(userId),
		Currency: param.Currency,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return response, customError.NewBadRequestError(errors.New("the balance not yet be created"))
		}
		return response, customError.NewInternalError(err)
	}

	balance, err := service.repo.AddBalance(ctx, sqlc.AddBalanceParams{
		UserID:   int32(userId),
		Currency: param.Currency,
		Balance:  strconv.FormatFloat(param.Amount, 'e', -1, 64),
	})
	if err != nil {
		return response, customError.NewInternalError(err)
	}

	response = model.DepositeResponse{
		BalanceResponse: model.BalanceResponse{
			Currency: balance.Currency,
			Balance:  balance.Balance,
		},
	}

	return
}
func (service *balanceService) Withdraw(ctx context.Context, param model.WithdrawRequestParam) (response model.WithdrawResponse, err error) {
	userId := getUserId(ctx)

	balance, err := service.repo.GetSpecifyCurrencyBalanceByUser(ctx, sqlc.GetSpecifyCurrencyBalanceByUserParams{
		UserID:   int32(userId),
		Currency: param.Currency,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return response, customError.NewBadRequestError(errors.New("the balance not yet be created"))
		}
		return
	}

	userBalance, err := strconv.ParseFloat(balance.Balance, 64)
	if err != nil {
		return response, customError.NewInternalError(err)
	}
	if userBalance < param.Amount {
		return response, customError.NewBadRequestError(errors.New("the balance is not enough"))
	}

	balance, err = service.repo.SubBalance(ctx, sqlc.SubBalanceParams{
		UserID:   balance.UserID,
		Currency: balance.Currency,
		Balance:  strconv.FormatFloat(param.Amount, 'e', -1, 64),
	})
	if err != nil {
		return response, customError.NewInternalError(err)
	}

	response = model.WithdrawResponse{
		BalanceResponse: model.BalanceResponse{
			Currency: balance.Currency,
			Balance:  balance.Balance,
		},
	}

	return
}
func (service *balanceService) Transfer(ctx context.Context, param model.TransferRequestParam) (response model.TransferResponse, err error) {
	userId := getUserId(ctx)

	balance, err := service.repo.GetSpecifyCurrencyBalanceByUser(ctx, sqlc.GetSpecifyCurrencyBalanceByUserParams{
		UserID:   int32(userId),
		Currency: param.Currency,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return response, customError.NewBadRequestError(errors.New("the balance not yet be created"))
		}
		return
	}

	userBalance, err := strconv.ParseFloat(balance.Balance, 64)
	if err != nil {
		return response, customError.NewInternalError(err)
	}
	if userBalance < param.Amount {
		return response, customError.NewBadRequestError(errors.New("the balance is not enough"))
	}

	_, err = service.repo.GetSpecifyCurrencyBalanceByUser(ctx, sqlc.GetSpecifyCurrencyBalanceByUserParams{
		UserID:   int32(param.ToUserId),
		Currency: param.Currency,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return response, customError.NewBadRequestError(errors.New("the to user's balance not yet be created"))
		}
		return
	}

	var result sqlc.TransferTxResponse
	result, err = service.repo.TransferTx(ctx, sqlc.TransferTxParam{
		FromUserID: userId,
		ToUserID:   param.ToUserId,
		Currency:   param.Currency,
		Amount:     param.Amount,
	})
	if err != nil {
		return response, customError.NewInternalError(err)
	}

	response = model.TransferResponse{
		BalanceResponse: model.BalanceResponse{
			Currency: result.FromBalance.Currency,
			Balance:  result.FromBalance.Balance,
		},
	}

	return
}

func (service *balanceService) GetActivity(ctx context.Context) (response []model.ListActivityResponse, err error) {
	userId := getUserId(ctx)

	activityList, err := service.repo.GetActivity(ctx, int32(userId))
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}
		return response, customError.NewInternalError(err)
	}

	for _, activity := range activityList {
		response = append(response, model.ListActivityResponse{
			UserId:   int(activity.UserID),
			Amount:   activity.Amount,
			CreateAt: activity.CreatedAt,
		})
	}

	return
}

func getUserId(ctx context.Context) int {
	userId, _ := ctx.Value("userId").(int)
	return userId
}
