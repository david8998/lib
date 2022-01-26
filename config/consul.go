package config

import (
	"github.com/hashicorp/consul/api"
	"log"
)

type ConsulConf struct {
	Addr  string
	Token string
}

var apiInstance *api.Client

func NewConsulAPI() (*api.Client, error) {
	if apiInstance == nil {
		conf := GetConsulConf()

		defaultConfig := api.DefaultConfig()
		defaultConfig.Address = conf.Addr
		defaultConfig.Token = conf.Token
		var err error
		apiInstance, err = api.NewClient(defaultConfig)
		if err != nil {
			log.Fatalf("new consul api err %v", err)
		}

	}
	return apiInstance, nil
}
func GetConsulConf() ConsulConf {
	env := GetEnv()
	if env == "test" {
		return ConsulConf{
			Addr:  "",
			Token: "",
		}
	} else if env == "prod" {
		return ConsulConf{
			Addr:  "",
			Token: "",
		}
	} else {
		return ConsulConf{
			Addr:  "127.0.0.1:8500",
			Token: "",
		}
	}
}
