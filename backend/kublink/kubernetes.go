package kublink

import (
	"fmt"
	_ "os"
	_ "os/exec"
	_ "strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var KUBECONFIG = ""
var USER = "ralbasini"
var HOME = "/Users/ralbasini"

var (
	config    *restclient.Config
	clientset *kubernetes.Clientset
)

// Connect function
func Connect(Kubeconfig_imported string) {

	KUBECONFIG = Kubeconfig_imported
	// KUBECONFIG = "/Users/ralbasini/Documents/iolab/kubedash/kubeconfig.yaml"

	// c := exec.Command("curl", "https://0.0.0.0:6443")
	// c.Stdout = os.Stdout
	// c.Stderr = os.Stderr
	// err2 := c.Run()
	// if err2 != nil {
	// 	fmt.Println("Error: ", err2)
	// }

	// fmt.Println(string(output))

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

	fmt.Println("Checking connection with host...")
	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Init... There are %d pods in the cluster\n", len(pods.Items))

}

// ClientSet function
func ClientSet() *kubernetes.Clientset {
	return clientset
}
