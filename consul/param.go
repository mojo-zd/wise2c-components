package consul

import (
	"fmt"

	"github.com/mojo-zd/wise2c-components/network"
	"github.com/mojo-zd/wise2c-components/tool"
)

var (
	DeRegisterCriticalServiceAfter = "30m"
	Interval                       = "15s"
	TTL                            = "10s"
)

type ConsulParam struct {
	ServerURL                      string //eg: consul:8500
	RegistryName                   string
	RegistryIp                     string
	RegistryID                     string
	RegistryPort                   int
	HealthCheckURL                 string
	DeRegisterCriticalServiceAfter string
	Interval                       string
	TTL                            string
	EnableTagOverride              bool
}

func (param *ConsulParam) Default() {
	if param == nil {
		fmt.Printf("consul param is nil!")
		return
	}

	if len(tool.Trim(param.DeRegisterCriticalServiceAfter)) == 0 {
		param.DeRegisterCriticalServiceAfter = DeRegisterCriticalServiceAfter
	}

	if len(tool.Trim(param.Interval)) == 0 {
		param.Interval = Interval
	}

	if len(tool.Trim(param.TTL)) == 0 {
		param.TTL = TTL
	}

	if len(tool.Trim(param.RegistryIp)) == 0 {
		param.RegistryIp = network.GetRegistryIp()
	}
	return
}
