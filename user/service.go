package main

import (
	common "github.com/dongsu8142/blog-common"
	pb "github.com/dongsu8142/blog-common/api"
)

type service struct {
	store UserStore
}

func NewService(store UserStore) *service {
	return &service{store}
}

func (s *service) RegisterUser(user *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	password, err := common.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	s.store.Register(user.Username, user.Email, password)
	res := &pb.RegisterUserResponse{
		Success: true,
		Message: "User registered",
	}
	return res, nil
}
