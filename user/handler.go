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
	log.Printf("Register user received! User %v", p)
	return h.service.RegisterUser(p)
}
