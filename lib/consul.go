package lib

import (
	"fmt"
	"github.com/david8998/lib/config"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
)

type ConsulAgent struct {
	client *api.Client
	agent  *api.Agent
}

func NewConsulAgent() *ConsulAgent {
	client, err := config.NewConsulAPI()
	if err != nil {
		log.Fatalf("new consul api err %v", err)
	}
	ret := &ConsulAgent{client: client, agent: client.Agent()}
	return ret
}
func (c *ConsulAgent) Register(param *api.AgentServiceRegistration) error {
	return c.agent.ServiceRegister(param)
}
func (c *ConsulAgent) Deregister(service, serviceID string) error {
	services, _, err := c.client.Health().Service(service, "", false, nil)
	if err != nil {
		return err
	}
	conf := config.GetConsulConf()
	for _, s := range services {
		if s.Service.ID == serviceID {
			defaultConfig := api.DefaultConfig()
			defaultConfig.Address = s.Node.Address + ":8500"
			defaultConfig.Token = conf.Token
			if apiInstance, err := api.NewClient(defaultConfig); err != nil {
				return err
			} else {
				return apiInstance.Agent().ServiceDeregister(serviceID)
			}

		}
	}
	return fmt.Errorf("service %s with ID %s is empty", service, serviceID)
}
func (c *ConsulAgent) GetAvailable(service string) ([]api.AgentServiceChecksInfo, error) {
	state, available, err := c.agent.AgentHealthServiceByName(service)
	if state != api.HealthPassing {
		return nil, fmt.Errorf("state %s have not available endpoint", state)
	}
	if err != nil {
		return nil, err
	}
	if len(available) > 0 {
		return available, nil
	} else {
		return nil, fmt.Errorf("have not available endpoint")
	}
}

func LocalIP() string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range address {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
