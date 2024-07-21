package main

import (
	"log"
	"net/http"

	common "github.com/dongsu8142/blog-common"
	pb "github.com/dongsu8142/blog-common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr        = common.EnvString("HTTP_ADDR", ":8000")
	userServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing user service at ", userServiceAddr)

	c := pb.NewUserServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
