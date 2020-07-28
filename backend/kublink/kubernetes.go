package kublink

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var KUBECONFIG = "/Users/ralbasini/.config/k3d/k3s-default/kubeconfig.yaml"
var USER = "ralbasini"
var HOME = "/Users/ralbasini"

var (
	config    *restclient.Config
	clientset *kubernetes.Clientset
)

// Connect function
func Connect() {

	// uses the current context in kubeconfig
	// path-to-kubeconfig -- for example, /root/.kube/config
	var err error
	config, err = clientcmd.BuildConfigFromFlags("", KUBECONFIG)
	if err != nil {
		panic(err)
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

}

// ClientSet function
func ClientSet() *kubernetes.Clientset {
	return clientset
}
