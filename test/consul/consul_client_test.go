package consul

import (
	"testing"

	"fmt"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/wise2c-components/consul"
)

func Test_Consul(t *testing.T) {
	client, err := consul.NewConsulClient(&consul.ConsulParam{ServerURL: "localhost:8500", RegistryName: "wise2c", RegistryPort: 8001, HealthCheckURL: "localhost:8001/health/check"})
	if err != nil {
		debug.Display("new consul client failed, error info is ", err.Error())
	}
	client.AutoRegistry = true
	client.AutoAgentRegistry()

	for {
		select {
		default:

		}
	}
}

func Test_Println(t *testing.T) {
	fmt.Printf("==== %s", "mojo")
}
