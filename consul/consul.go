package consul

import (
	"gopkg.in/yaml.v2"

	"github.com/andrepinto/littleboat/utils"
	"github.com/hashicorp/consul/api"

	"encoding/json"

)


type ConsulConfigOptions struct {

	Schema string
	Endpoint string
}

type ConsulConfig struct {
	Api *api.Config
	Key string
	Format string
}

func NewConsulConfig(options *ConsulConfigOptions) *ConsulConfig{
	return &ConsulConfig{
		Api: &api.Config{
			Address: options.Endpoint,
			Scheme:  options.Schema,
		},
	}
}

func (cf *ConsulConfig) Load(cfg interface{},  key string, format string) (interface{}, error){

	cf.Format = format
	cf.Key = key

	switch cf.Format {
	case utils.YAML:
		return cf.LoadYAMLFromRemote(cfg)
	case utils.JSON:
		return cf.LoadJSONFromRemote(cfg)
	default:
		panic("Format not supported")
	}
}

func (cf *ConsulConfig)LoadJSONFromRemote( cfg interface{}) (interface{}, error) {


	configKey := cf.Key
	client, err := api.NewClient(cf.Api)
	if err != nil {
		return nil, err
	}
	kv := client.KV()
	kvPair, _, err := kv.Get(configKey, nil)
	if err != nil {
		return nil, err
	}
	if kvPair == nil {
		return nil, err
	}
	if len(kvPair.Value) == 0 {
		return nil, err
	}
	if err = json.Unmarshal(kvPair.Value, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}


func (cf *ConsulConfig)LoadYAMLFromRemote( cfg interface{}) (interface{}, error) {
	configKey := cf.Key
	client, err := api.NewClient(cf.Api)
	if err != nil {
		return nil, err
	}
	kv := client.KV()
	kvPair, _, err := kv.Get(configKey, nil)
	if err != nil {
		return nil, err
	}
	if kvPair == nil {
		return nil, err
	}
	if len(kvPair.Value) == 0 {
		return nil, err
	}

	if err =  yaml.Unmarshal(kvPair.Value, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (cf *ConsulConfig) Listen(keys []string, ch chan string) error{


	return nil
}