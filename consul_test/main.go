package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register (address string ,port int, name string, tags []string, id string) error{

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", address, port)

	client , err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Address = address
	registration.Port = port
	registration.Tags = tags

	check := &api.AgentServiceCheck{
		HTTP: "http://172.20.0.204:8021/health",
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	return nil
}

func AllService(address string ,port int) {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", address, port)

	client , err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data ,err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for key,_ := range data {
		fmt.Println(key)
	}

}

func FilterService(address string ,port int) {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", address, port)

	client , err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data ,err := client.Agent().ServicesWithFilter(`Service == "mxshop_svr"`)
	if err != nil {
		panic(err)
	}

	for key,_ := range data {
		fmt.Println(key)
	}
}

func main() {
	// Register("172.20.0.204", 8500, "mxshop_svr", []string{"mxshop"}, "mxshop_test1")
	// AllService("172.20.0.204", 8500)
	FilterService("172.20.0.204", 8500)
}
