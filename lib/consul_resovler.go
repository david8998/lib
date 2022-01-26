package lib

import (
	"fmt"
	"github.com/david8998/lib/config"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

type consulBuilder struct {
}

type consulResolver struct {
	cc                   resolver.ClientConn
	service              string
	disableServiceConfig bool
	lastIndex            uint64
}

func NewBuilder() resolver.Builder {
	return &consulBuilder{}
}

func (cb *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	cr := &consulResolver{
		service:              target.Endpoint,
		cc:                   cc,
		disableServiceConfig: opts.DisableServiceConfig,
		lastIndex:            0,
	}

	go cr.watcher()
	return cr, nil

}

func (cb *consulBuilder) Scheme() string {
	return "consul"
}
func (cr *consulResolver) watcher() {
	client, err := config.NewConsulAPI()
	if err != nil {
		log.Fatalf("error create consul client: %v\n", err)
		return
	}

	for {
		services, metaInfo, err := client.Health().Service(cr.service, "", true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			fmt.Printf("error retrieving instances from Consul: %v", err)
		}

		cr.lastIndex = metaInfo.LastIndex
		var state resolver.State
		for _, service := range services {
			addr := fmt.Sprintf("%v:%v", service.Service.Address, service.Service.Port)
			state.Addresses = append(state.Addresses, resolver.Address{Addr: addr})
		}
		if err := cr.cc.UpdateState(state); err != nil {
			log.Println("update grpc server err ", err)
			time.Sleep(time.Second)
		}
	}

}

func (cr *consulResolver) ResolveNow(opt resolver.ResolveNowOptions) {
}

func (cr *consulResolver) Close() {
}
