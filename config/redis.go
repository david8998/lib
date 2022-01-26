package config

import (
	"log"
	"sync"
	"time"
)

type Conf struct {
	Redis redisCfg `json:"redis"`
}

var C Conf

var once sync.Once

func InitRedisConfig() {
	confPath := "ft_trade/data_center/redis.json"
	once.Do(func() {
		if err := getConsulConf(confPath, &C); err != nil {
			log.Fatalf("init redis conf err %v", err)
		}
	})
}

type redisCfg struct {
	Addr         string        `json:"addr"`
	Password     string        `json:"password,omitempty"`
	DB           int           `json:"db,omitempty"`
	DialTimeout  time.Duration `json:"dial_timeout,omitempty"`  //建连接的时间
	ReadTimeout  time.Duration `json:"read_timeout,omitempty"`  //读超时, 等待服务响应的时间
	WriteTimeout time.Duration `json:"write_timeout,omitempty"` //写超时, 发送命令的时间
	MaxConnAge   time.Duration `json:"max_conn_age,omitempty"`  //连接存活时间
	MaxCons      int           `json:"max_cons,omitempty"`      //最大连接数
	MinIdleCons  int           `json:"min_idle_cons,omitempty"` //最小连接数
}
