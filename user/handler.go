package main

import (
	"context"
	"log"

	pb "github.com/dongsu8142/blog-common/api"
	"google.golang.org/grpc"
)

type Handler struct {
	pb.UnimplementedUserServiceServer

	service UserService
}

func NewHandler(grpcServer *grpc.Server, service UserService) {
	handler := &Handler{
		service: service,
	}
	pb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *Handler) RegisterUser(ctx context.Context, p *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Println("Register user received!")
	return h.service.RegisterUser(p)
}

func (h *Handler) LoginUser(ctx context.Context, p *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	log.Println("Login user recevied!")
	return h.service.LoginUser(p)
}