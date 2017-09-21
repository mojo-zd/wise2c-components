### wise2c公用组件
#### consul组件
- 注册consul
```
client, err := consul.NewConsulClient(&consul.ConsulParam{ServerURL: "localhost:8500", RegistryName: "wise2c", RegistryPort: 8001, HealthCheckURL: "localhost:8001/health/check"})
if err != nil {
    debug.Display("new consul client failed, error info is ", err.Error())
}
client.AgentRegistry()
```

- 自动注册consul
```
client, err := consul.NewConsulClient(&consul.ConsulParam{ServerURL: "localhost:8500", RegistryName: "wise2c", RegistryPort: 8001, HealthCheckURL: "localhost:8001/health/check"})
if err != nil {
    debug.Display("new consul client failed, error info is ", err.Error())
}
client.AutoRegistry = true
client.AutoAgentRegistry()
```