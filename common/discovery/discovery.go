package discovery

import (
	"fmt"
	"math/rand"
	"time"
)

type Registry interface {
	Register(instanceID, serverName, hostPort string) error
	Deregister(instanceID, serviceName string) error
	Discover(serviceName string) ([]string, error)
	HealthCheck(instanceID, serviceName string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}