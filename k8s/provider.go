package k8s

import (
	"os"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/informers"
	"time"
)

type K8SConfigProviderConfig struct {
	Namespace string
}

type K8SConfigProvider struct{
	Client *KubeClient
	Namespace string
}

func NewK8SConfigProvider(options *K8SConfigProviderConfig)(*K8SConfigProvider, error){
	client, err := NewKubeClient(os.Getenv("KUBECONFIG"))

	if err != nil{
		return nil, err
	}

	if len(options.Namespace)==0{
		options.Namespace = "default"
	}

	return &K8SConfigProvider{
		Client: client,
		Namespace: options.Namespace,
	}, nil

}


func (cf *K8SConfigProvider) Load(cfg interface{},  key string, format string) (interface{}, error){
	data, err := cf.Client.GetConfigMap(cf.Namespace, key)

	if err = yaml.Unmarshal([]byte(data.Data[key]), cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}

func (cf *K8SConfigProvider) Listen(keys []string, ch chan string) error{

	sharedInformers := informers.NewSharedInformerFactory(cf.Client.Client, 10*time.Minute)

	cf.Client.Watch(keys, cf.Namespace, ch, sharedInformers.Core().V1().ConfigMaps())

	sharedInformers.Start(nil)

	return nil
}

