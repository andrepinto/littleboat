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