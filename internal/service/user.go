package service

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/wizard-xie/realworld/api/realworld/v1"
)

func (srv *RealworldService) CurrentUser(ctx context.Context, _ *emptypb.Empty) (*v1.UserResponse, error) {
	return &v1.UserResponse{
		User: &v1.User{
			Email:    "986740642@qq.com",
			Username: "xieweizhi",
		},
	}, nil
}

func (srv *RealworldService) Login(context.Context, *v1.LoginRequest) (*v1.UserResponse, error) {
	return &v1.UserResponse{
		User: &v1.User{
			Email:    "986740642@qq.com",
			Username: "xieweizhi",
		},
	}, nil
}
