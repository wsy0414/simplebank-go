package gapi

import (
	"simplebank/api/service"
	"simplebank/db/sqlc"
	"simplebank/pb"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	userService    service.UserService
	balanceService service.BalanceService
}

func NewServer(store sqlc.Store) *Server {
	server := &Server{
		userService:    service.NewUserService(store),
		balanceService: service.NewBalanceService(store),
	}

	return server
}
