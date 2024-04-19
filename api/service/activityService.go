package service

import (
	"database/sql"
	"simplebank/api/model"
	"simplebank/db/sqlc"
	"simplebank/handler"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityService interface {
	Deposite(*gin.Context, int, model.DepositeRequestParam) (model.DepositeResponse, error)
	Withdraw(*gin.Context)
	Transfer(*gin.Context)
	ListActivities(*gin.Context)
	ListTransfers(*gin.Context)
}

type activityService struct {
	repo *sql.DB
}

func NewActivityService(repo *sql.DB) ActivityService {
	return activityService{
		repo: repo,
	}
}

func (a activityService) Deposite(ctx *gin.Context, userId int, param model.DepositeRequestParam) (response model.DepositeResponse, err error) {
	// check user exist
	// check amount is valid
	// insert activity
	var bb sqlc.Balance
	err = handler.HandleTransaction(a.repo, func(q sqlc.Queries) error {
		createActivityParam := sqlc.CreateActivityParams{
			UserID: int32(userId),
			Amount: strconv.FormatFloat(param.Amount, 'e', -1, 64),
		}

		_, err := q.CreateActivity(ctx, createActivityParam)
		if err != nil {
			return err
		}

		getBalanceParam := sqlc.GetSpecifyCurrencyBalanceByUserParams{
			UserID:   int32(userId),
			Currency: param.Currency,
		}
		balance, err := q.GetSpecifyCurrencyBalanceByUser(ctx, getBalanceParam)
		if err != nil {
			if err == sql.ErrNoRows {
				// if err is noRowErr, insert balance
				createBalanceParam := sqlc.CreateBalanceParams{
					UserID:   int32(userId),
					Currency: param.Currency,
					Balance:  strconv.FormatFloat(param.Amount, 'e', -1, 64),
				}
				balance, err = q.CreateBalance(ctx, createBalanceParam)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			addBalanceParam := sqlc.AddBalanceParams{
				Balance:  strconv.FormatFloat(param.Amount, 'e', -1, 64),
				UserID:   int32(userId),
				Currency: param.Currency,
			}
			balance, err = q.AddBalance(ctx, addBalanceParam)
			if err != nil {
				return err
			}
		}

		bb = balance

		return nil
	})
	if err != nil {
		return
	}
	// insert or update balance
	// maybe send email
	response = model.DepositeResponse{
		Currency: bb.Currency,
		Balance:  bb.Balance,
	}
	// return balance
	return
}

func (a activityService) Withdraw(ctx *gin.Context) {
	// check user exist
	// check balance
	// check amount is valid
	// insert activity
	// update balance
	// maybe send email
	// return balance
}

func (a activityService) Transfer(ctx *gin.Context) {
	// check from and to user.id is exist
	// check from.user.balance is valid
	// check transfer amount is valid
	// insert activity(from and to)
	// insert transfer
	// insert or update balance(from and to)
	// maybe send email to to_user
	// return from_user.balance
}

func (a activityService) ListActivities(ctx *gin.Context) {
	// get a user's activities
}

func (a activityService) ListTransfers(ctx *gin.Context) {
	// get a user's transfer history list
}
