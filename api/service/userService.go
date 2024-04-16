package service

import (
	"database/sql"
	"errors"
	"simplebank/api/model"
	"simplebank/db/sqlc"
	"simplebank/util"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	SignUp(*gin.Context, *model.SignUpRequestParam) (model.SignUpResponse, error)
	Login(*gin.Context, *model.LoginRequestParam) (model.LoginResponse, error)
	GetUserInfo(*gin.Context, int) (model.GetUserInfoResponse, error)
}

type userService struct {
	repo *sqlc.Queries
}

func NewUserService(db *sql.DB) UserService {
	return &userService{
		repo: sqlc.New(db),
	}
}

func (us userService) SignUp(ctx *gin.Context, param *model.SignUpRequestParam) (response model.SignUpResponse, err error) {
	// encrypt password
	pwd, err := util.EncryptPassword(param.Password)
	if err != nil {
		return
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
		return
	}

	// generate JWT token
	token, err := util.GenerateToken(int(user.ID), util.TOKEN_DEFAULT_DURATION)
	if err != nil {
		return
	}
	// response user.id and token
	response = model.SignUpResponse{
		ID:    int(user.ID),
		Token: token,
	}

	return
}

func (u *userService) Login(ctx *gin.Context, param *model.LoginRequestParam) (response model.LoginResponse, err error) {
	// check user exist
	user, err := u.repo.GetUserByName(ctx, param.Name)
	if err != nil {

		return
	}

	// check password after encrypt equal database
	if !util.CheckPassword(param.Password, user.Password) {
		return response, errors.New("password isn't correct")
	}

	// generate JWT Token
	token, err := util.GenerateToken(int(user.ID), util.TOKEN_DEFAULT_DURATION)
	if err != nil {
		return
	}

	// respone user.id and token
	response = model.LoginResponse{
		ID:    int(user.ID),
		Token: token,
	}

	return
}

func (u userService) GetUserInfo(ctx *gin.Context, userId int) (response model.GetUserInfoResponse, err error) {
	user, err := u.repo.GetUser(ctx, int32(userId))
	if err != nil {
		return
	}

	balanceList, err := u.repo.GetBalanceByUser(ctx, int32(userId))
	if err != nil {
		return
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
