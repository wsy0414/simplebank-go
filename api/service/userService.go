package service

import (
	"context"
	"database/sql"
	"errors"
	"simplebank/api/model"
	"simplebank/customError"
	"simplebank/db/sqlc"
	"simplebank/util"

	"github.com/lib/pq"
)

type UserService interface {
	SignUp(context.Context, *model.SignUpRequestParam) (model.SignUpResponse, error)
	Login(context.Context, *model.LoginRequestParam) (model.LoginResponse, error)
	GetUserInfo(context.Context, int) (model.GetUserInfoResponse, error)
}

type userService struct {
	repo sqlc.Store
}

func NewUserService(store sqlc.Store) UserService {
	return &userService{
		repo: store,
	}
}

func (us userService) SignUp(ctx context.Context, param *model.SignUpRequestParam) (response model.SignUpResponse, err error) {
	// encrypt password
	pwd, err := util.EncryptPassword(param.Password)
	if err != nil {
		return response, customError.NewInternalError(err)
	}

	createUserParam := &sqlc.CreateUserParams{
		Name:     param.Name,
		Password: pwd,
		Birthdate: sql.NullTime{
			Time: param.Birthdate,
		},
		Email: param.Email,
	}
	// insert user
	user, err := us.repo.CreateUser(ctx, *createUserParam)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				return response, customError.NewBadRequestError(err)
			}
		}
		return response, customError.NewInternalError(err)
	}

	// generate JWT token
	token, err := util.GenerateToken(int(user.ID), util.TOKEN_DEFAULT_DURATION)
	if err != nil {
		return response, customError.NewInternalError(err)
	}
	// response user.id and token
	response = model.SignUpResponse{
		ID:    int(user.ID),
		Token: token,
	}

	return
}

func (u *userService) Login(ctx context.Context, param *model.LoginRequestParam) (response model.LoginResponse, err error) {
	// check user exist
	user, err := u.repo.GetUserByName(ctx, param.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, customError.NewBadRequestError(err)
		}
		return response, customError.NewInternalError(err)
	}

	// check password after encrypt equal database
	if !util.CheckPassword(param.Password, user.Password) {
		return response, customError.NewBadRequestError(errors.New("password isn't correct"))
	}

	// generate JWT Token
	token, err := util.GenerateToken(int(user.ID), util.TOKEN_DEFAULT_DURATION)
	if err != nil {
		return response, customError.NewInternalError(err)
	}

	// respone user.id and token
	response = model.LoginResponse{
		ID:    int(user.ID),
		Token: token,
	}

	return
}

func (u userService) GetUserInfo(ctx context.Context, userId int) (response model.GetUserInfoResponse, err error) {
	user, err := u.repo.GetUser(ctx, int32(userId))
	if err != nil {
		if err == sql.ErrNoRows {
			return response, customError.NewBadRequestError(errors.New("this user is not exist"))
		}
		return response, customError.NewInternalError(err)
	}

	balanceList, err := u.repo.GetBalanceByUser(ctx, int32(userId))
	if err != nil {
		if err == sql.ErrNoRows {
			balanceList = []sqlc.Balance{}
		} else {
			return
		}
	}

	var list = make([]model.Balance, 0)
	for _, balance := range balanceList {
		list = append(list, model.Balance{
			Currency: balance.Currency,
			Balance:  balance.Balance,
		})
	}

	response = model.GetUserInfoResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Birthdate: model.CustomTime{user.Birthdate.Time},
		Email:     user.Email,
		Balance:   list,
	}

	return
}
