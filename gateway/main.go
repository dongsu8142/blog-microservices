package main

import (
	"log"
	"net/http"
	"time"

	common "github.com/dongsu8142/blog-common"
	"github.com/dongsu8142/blog-common/discovery"
	"github.com/dongsu8142/blog-common/discovery/consul"
	"github.com/dongsu8142/blog-gateway/gateway"
)

var (
	serviceName = "gateway"
	httpAddr        = common.EnvString("HTTP_ADDR", ":8000")
	consulAddr = common.EnvString("CONSUL_ADDR", "localhost:8500")
)

func main() {
	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(instanceID, serviceName, httpAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.Deregister(instanceID, serviceName)

	mux := http.NewServeMux()

	userGateway := gateway.NewGRPCGateway(registry)

	handler := NewHandler(userGateway)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
