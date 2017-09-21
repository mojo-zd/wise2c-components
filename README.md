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
- ConsulParam参数说明
1. ServerURL  consul地址 eg: http://xx.xx.xx.xx:8500 (必设)
2. RegistryName 注册到consul的名称 eg: resource-manager (必设)
3. RegistryIp 注册到consul的ip地址 (该组件已经自动获取了运行环境的ip地址,不需要用户再去设置了)
4. RegistryID 注册到consul的id （不设置会随机生成一个）
5. RegistryPort 当前应用运行的端口（必设）
6. HealthCheckURL 当前应用提供的health check api eg: /api/health (必设)