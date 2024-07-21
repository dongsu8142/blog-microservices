package main

import pb "github.com/dongsu8142/blog-common/api"

type UserService interface {
	RegisterUser(*pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
}

type UserStore interface {
	Register() error
}
