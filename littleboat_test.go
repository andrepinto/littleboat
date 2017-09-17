package littleboat

import (
	"github.com/andrepinto/littleboat/consul"
	"fmt"
	"testing"
	"github.com/andrepinto/littleboat/utils"
	"github.com/andrepinto/littleboat/k8s"
	"os"
)

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}


func Test_LoadConsul(t *testing.T){
	provider := consul.NewConsulConfig(&consul.ConsulConfigOptions{
		Schema:"http",
		Endpoint:"localhost:8500",
	})
	var st T

	cm := NewConfigManager(provider)

	err := cm.Get("services/demo",&st, utils.YAML)

	fmt.Println(err, st.A)
}



type ServiceData struct {
	Code string `yaml:"code"`
	Name string `yaml:"name"`
	Schema string `yaml:"schema"`
	Endpoint string `yaml:"endpoint"`
	Port int `yaml:"port"`
	BasePath string `yaml:"basePath"`
	Cluster  string `yaml:"cluster"`
}

type Service struct {
	Services ServicesData
}

type ServicesData []* ServiceData

func Test_LoadK8S(t *testing.T){

	os.Setenv("KUBECONFIG", "./kubeconfig")

	provider, err := k8s.NewK8SConfigProvider(
		&k8s.K8SConfigProviderConfig{
			Namespace:"gateway",
		},
	)

	if err!=nil{
		t.Error(err)
	}

	cm := NewConfigManager(provider)

	var st Service

	 cm.Get("gateway-services.yaml",&st, utils.YAML)


	fmt.Println(st.Services[0].Code)
}


func Test_ListemK8S(t *testing.T){

	os.Setenv("KUBECONFIG", "./kubeconfig")

	provider, err := k8s.NewK8SConfigProvider(
		&k8s.K8SConfigProviderConfig{
			Namespace:"gateway",
		},
	)

	if err!=nil{
		t.Error(err)
	}

	cm := NewConfigManager(provider)

	ch := make(chan string)

	cm.Listen([]string{"gateway-services.yaml"}, ch)

	select {
	case v := <-ch:
		fmt.Println(v)
	}
}
