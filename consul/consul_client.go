package consul

import (
	"errors"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/mojo-zd/wise2c-components/tool"
	"github.com/robfig/cron"
)

var (
	EXPRESSION = "*/5 * * * * *"
)

type ConsulClient struct {
	*ConsulParam
	AutoRegistry bool
	CheckExpress string //自动注册开启时, 此参数为用于检查当前service的健康状态的corn表达式 AutoRegistry开启时不设置此参数 "EXPRESSION"则为默认间隔调用表达式
	c            *api.Client
}

func NewConsulClient(params *ConsulParam) (client *ConsulClient, err error) {
	c, err := api.NewClient(api.DefaultConfig())

	if err != nil {
		fmt.Printf("new consul c error, error info is %s", err.Error())
		return
	}

	params.Default()
	client = &ConsulClient{ConsulParam: params, c: c}
	err = client.validate()
	return
}

func (client *ConsulClient) SetCronExpression(expression string) *ConsulClient {
	client.CheckExpress = expression
	return client
}

//自动注册开启的情况下, 可以调用此方法
func (client *ConsulClient) AutoAgentRegistry() {
	client.AgentRegistry()
	if client.AutoRegistry {
		c := cron.New()
		if len(client.CheckExpress) > 0 {
			EXPRESSION = client.CheckExpress
		}

		c.AddFunc(EXPRESSION, func() {
			services, _, err := client.c.Health().Service(client.RegistryName, "", true, &api.QueryOptions{})
			if err != nil || len(services) == 0 {
				if err = client.AgentRegistry(); err != nil {
					fmt.Println("==== registry error info is %s " + err.Error())
					return
				}
				fmt.Println("=== registry consul again, registry id is " + client.RegistryID)
			}
		})
		c.Start()
	}
}

//只注册一次, 断开后不会自动注册consul
func (client *ConsulClient) AgentRegistry() (err error) {
	err = client.c.Agent().ServiceRegister(client.newAgentServiceRegistration())
	if err != nil {
		fmt.Printf("registry consul failed, error info is %s", err.Error())
		return
	}
	return
}

//获取指定service在consul上的注册地址
func (client *ConsulClient) GetServiceAddress(serviceName string) (address string, err error) {
	services, _, err := client.c.Health().Service(serviceName, "", true, nil)
	if err != nil {
		fmt.Errorf("get services from consul failed, error info is %s", err.Error())
		return
	}

	if len(services) == 0 {
		fmt.Errorf("no available service %s!!!", serviceName)
		err = errors.New(fmt.Sprintf("no available service %s", serviceName))
		return
	}

	agentService := services[0].Service
	address = fmt.Sprintf("%s:%d", agentService.Address, agentService.Port)
	return

}

func (client *ConsulClient) newAgentServiceRegistration() (agentService *api.AgentServiceRegistration) {
	return &api.AgentServiceRegistration{
		ID:                fmt.Sprintf("%s-%s", client.RegistryID, client.RegistryIp),
		Name:              client.RegistryName,
		EnableTagOverride: client.EnableTagOverride,
		Address:           client.RegistryIp,
		Port:              client.RegistryPort,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: client.DeRegisterCriticalServiceAfter,
			HTTP:     fmt.Sprintf("http://%s:%d/%s", client.RegistryIp, client.RegistryPort, client.HealthCheckURL),
			Interval: client.Interval,
			Timeout:  client.TTL,
		},
	}
	return
}

func (client *ConsulClient) validate() (err error) {
	if len(tool.Trim(client.HealthCheckURL)) == 0 {
		err = errors.New("health check url must be set!")
		return
	}

	if len(tool.Trim(client.RegistryName)) == 0 {
		err = errors.New("registry name must be set!")
		return
	}

	if len(tool.Trim(client.ServerURL)) == 0 {
		err = errors.New("consul server url must be set!")
	}

	if client.RegistryPort == 0 {
		err = errors.New("registry port must be set!")
		return
	}

	if len(tool.Trim(client.RegistryIp)) == 0 {
		err = errors.New("registry ip must be set!")
		return
	}

	if len(tool.Trim(client.RegistryID)) == 0 {
		client.RegistryID = tool.UUID()
		fmt.Printf("registry id is not set and generate a random string, value is %s \n", client.RegistryID)
	}
	return
}
