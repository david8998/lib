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
			Addr:  "47.100.85.33:8500",
			Token: "4d85aeb4-a4e2-f584-dc16-a33d4151bdc8",
		}
	} else if env == "prod" {
		return ConsulConf{
			Addr:  "internal-consul-elb-466801138.ap-northeast-1.elb.amazonaws.com:8500",
			Token: "bd792f9f-20d9-cea8-a5c7-f1dad30d9f33",
		}
	} else {
		return ConsulConf{
			Addr:  "127.0.0.1:8500",
			Token: "",
		}
	}
}
