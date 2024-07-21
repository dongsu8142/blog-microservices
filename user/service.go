package main

import pb "github.com/dongsu8142/blog-common/api"

type service struct {
	store UserStore
}

func NewService(store UserStore) *service {
	return &service{store}
}

func (s *service) RegisterUser(*pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := &pb.RegisterUserResponse{
		Success: true,
		Message: "User registered",
	}
	return user, nil
}
