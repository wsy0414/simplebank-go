package gapi

import (
	"context"
	"simplebank/api/model"
	"simplebank/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	param := model.LoginRequestParam{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}
	response, err := server.userService.Login(ctx, &param)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error", err.Error())
	}

	return &pb.LoginResponse{
		Id:    int32(response.ID),
		Token: response.Token,
	}, nil
}
