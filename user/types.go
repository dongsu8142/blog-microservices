package main

import (
	pb "github.com/dongsu8142/blog-common/api"
	"github.com/dongsu8142/blog-common/database"
)

type UserService interface {
	RegisterUser(*pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	LoginUser(*pb.LoginUserRequest) (*pb.LoginUserResponse, error)
}

type UserStore interface {
	Register(string, string, string) error
	Login(string) (*database.User, error)
}
