package k8s

import (

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/pkg/api/v1"
	"fmt"
	informercorev1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
	apicorev1 "k8s.io/client-go/pkg/api/v1"
)


type KubeClient struct {
	Client *kubernetes.Clientset
}

func NewKubeClient(kubeconfig string) (*KubeClient, error){

	var (
		config *rest.Config
		err    error
	)

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, err
	}

	client := kubernetes.NewForConfigOrDie(config)

	kc := &KubeClient{
		Client: client,
	}

	return  kc, nil

}


func(kc *KubeClient) GetConfigMap(namespace string, name string) (*v1.ConfigMap, error){
	cfgMap, err := kc.Client.CoreV1().ConfigMaps(namespace).Get(name,metav1.GetOptions{})
	return cfgMap, err

}


func(kc *KubeClient) Watch(keys []string, namespace string, ch chan string, secretInformer informercorev1.ConfigMapInformer){

	secretInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			UpdateFunc: func(oldObj, newObj interface{}) {
				data := newObj.(*apicorev1.ConfigMap)
				key, _ := cache.MetaNamespaceKeyFunc(newObj)

				for _ ,v := range keys{
					if fmt.Sprintf("%s/%s", namespace, v) == key{
						ch <- data.Name
					}
				}


			},
		},
	)
}