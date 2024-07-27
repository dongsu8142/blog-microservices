package gateway

import (
	"context"

	pb "github.com/dongsu8142/blog-common/api"
)

type UserGateway interface {
	RegisterUser(context.Context, *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
}