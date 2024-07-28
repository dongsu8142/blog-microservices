package main

import (
	"errors"

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
	if err := s.store.Register(user.Username, user.Email, password); err != nil {
		return nil, err
	}
	res := &pb.RegisterUserResponse{
		Success: true,
		Message: "User registered",
	}
	return res, nil
}

func (s *service) LoginUser(req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := s.store.Login(req.Username)
	if err != nil {
		return nil, err
	}
	if !common.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid password")
	}
	secret := []byte(common.EnvString("SECRET", "q1w2e3r4"))
	token, err := common.CreateJWT(secret, int(user.ID))
	if err != nil {
		return nil, err
	}
	return &pb.LoginUserResponse{Token: token}, nil
}
