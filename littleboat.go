package littleboat

import (
	"sync"
	"errors"
)

var once sync.Once

type ConfigurationProvider interface {
	Load(cfg interface{}, key string, format string) (interface{}, error)
	Listen(keys []string, ch chan string) error
	SimpleLoad(key string) ([]byte, error)
}


type ConfigManager struct {
	provider ConfigurationProvider
}

func NewConfigManager(provider ConfigurationProvider) *ConfigManager{
	return &ConfigManager{
		provider: provider,
	}
}

func(cm *ConfigManager) Get(key string, cfg interface{}, format string) (error) {

	if len(format)==0{
		return  errors.New("Format is empty")
	}

	if len(key)==0{
		return  errors.New("Key is empty")
	}

	_, err := cm.provider.Load(cfg, key, format)

	return  err
}

func(cm *ConfigManager) SimpleGet(key string) ([]byte,error) {
	data, err := cm.provider.SimpleLoad(key)
	return  data, err
}

func(cm *ConfigManager) Listen(keys []string, ch chan string){
	cm.provider.Listen(keys, ch)
}
