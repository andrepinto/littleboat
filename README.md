# LittleBoat - Golang Configuration manager

Is a simple configuration solution for golang applications. It works only by consuming remotely configurations and was designed to work with kubernetes configmaps

**Note:** It also supports consul (the watch feature for consul is in development)

## Get config

Get(key string, cfg interface{}, format string)

## Watch

Listen(keys []string, ch chan string)


## Example

### Kubernetes

```go

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

``

### Consul

```go
package examples

import (
	"fmt"
	"github.com/andrepinto/littleboat/consul"
	"github.com/andrepinto/littleboat"
	"github.com/andrepinto/littleboat/utils"
)

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main(){
	provider := consul.NewConsulConfig(&consul.ConsulConfigOptions{
		Schema:"http",
		Endpoint:"localhost:8500",
	})
	var st T

	cm := littleboat.NewConfigManager(provider)

	err := cm.Get("services/demo",&st, utils.YAML)

	fmt.Println(err, st.A)
}
``