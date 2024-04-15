package service

import "simplebank/model"

func SignUp(param *model.SignUpRequestParam) (response model.SignUpResponse, err error) {
	// check userName and email
	// check password is valid
	// encrypt password
	// insert user
	// generate JWT token
	// response user.id and token

	return
}

func Login() {
	// check user exist
	// check password after encrypt equal database
	// generate JWT Token
	// respone user.id and token
}

func GetUser() {
	// check user
	// return user info and balance list
}
