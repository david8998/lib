package config

import (
	"encoding/json"
	"fmt"
)

var env string

func SetEnv(e string) (err error) {
	switch e {
	case "dev", "test", "prod":
		env = e
		break
	default:
		err = fmt.Errorf("error env %s", e)
		return
	}
	return
}

func GetEnv() string {
	return env
}

func getConsulConf(path string, conf interface{}) (err error) {
	consulApi, err := NewConsulAPI()
	if err != nil {
		return
	}
	kv := consulApi.KV()
	pair, _, err := kv.Get(path, nil)
	if err != nil {
		return err
	}
	if pair == nil || len(pair.Value) == 0 {
		err = fmt.Errorf("get config value nil")
		return
	}

	err = json.Unmarshal(pair.Value, &conf)
	return
}
