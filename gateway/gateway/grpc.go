package gateway

import (
	"context"
	"log"

	pb "github.com/dongsu8142/blog-common/api"
	"github.com/dongsu8142/blog-common/discovery"
)

type gateway struct {
	registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
	return &gateway{registry}
}

func (g *gateway) RegisterUser(ctx context.Context, user *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	conn, err := discovery.ServiceConnection("user", g.registry)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewUserServiceClient(conn)

	return c.RegisterUser(ctx, user)
}

func (g *gateway) LoginUser(ctx context.Context, user *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	conn, err := discovery.ServiceConnection("user", g.registry)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewUserServiceClient(conn)

	return c.LoginUser(ctx, user)
}