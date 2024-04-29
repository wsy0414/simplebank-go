package gapi

import (
	"context"
	"simplebank/api/model"
	"simplebank/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server Server) SignUp(ctx context.Context, req *pb.SignUpUserRequest) (*pb.SignUpUserResponse, error) {
	param := model.SignUpRequestParam{
		Name:      req.Name,
		Password:  req.Password,
		Email:     req.Password,
		Birthdate: req.Birthdate.AsTime(),
	}
	response, err := server.userService.SignUp(ctx, &param)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create user failed, %s", err.Error())
	}

	return &pb.SignUpUserResponse{
		Id:    int32(response.ID),
		Token: response.Token,
	}, nil
}
