package examples

import (
	"os"
	"fmt"
	"github.com/andrepinto/littleboat/k8s"
	"github.com/andrepinto/littleboat"
	"github.com/andrepinto/littleboat/utils"
)

func main(){
	os.Setenv("KUBECONFIG", "./kubeconfig")

	provider, err := k8s.NewK8SConfigProvider(
		&k8s.K8SConfigProviderConfig{
			Namespace:"dev",
		},
	)

	if err!=nil{
		panic(err)
	}

	cm := littleboat.NewConfigManager(provider)

	var st Service

	cm.Get("services",&st, utils.YAML)

	fmt.Println(st.Services[0].Code)


	ch := make(chan string)

	cm.Listen([]string{"services"}, ch)

	select {
	case v := <-ch:
		fmt.Println(v)
	}

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