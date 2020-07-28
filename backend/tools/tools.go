package tools

import (
	"fmt"
	"os"
	"os/exec"
	v1 "k8s.io/api/core/v1"
	"github.com/wayoos/kubedash/backend/kublink"
)

type ExtendedNs struct {
    Namespace   string
    State		string
}

func runCmd(cmd string) string {

	fmt.Println("Running cmd [" + cmd + "]")
	out, err := exec.Command("/bin/bash","-c", cmd).Output()
	if err != nil {
			fmt.Printf("Failed to execute command: [%s]\n", cmd)
			return ""
	}

	result := string(out)
	return result
}

func runKubectl(cmd string) string {

	fmt.Println("Running kubectl cmd [" + cmd + "]")

	cmd_config := exec.Command("/bin/bash","-c", cmd)
	cmd_config.Env = os.Environ()
	cmd_config.Env = append(cmd_config.Env, "KUBECONFIG="+ kublink.KUBECONFIG)
	out, err := cmd_config.Output()

	if err != nil {
			fmt.Printf("Failed to execute kubectl command: [%s]\n", cmd)
			return ""
	}

	result := string(out)
	return result
}

func getDeprecatedPluto(ns v1.Namespace) string {
	nsName := ns.ObjectMeta.Name
	fmt.Println("\n[Checking resources in ns " + nsName + "]")

	// var plutoVersion = runCmd("pluto version | awk -F \"[: ]\" '{print $2}'")
	// fmt.Println(plutoVersion)

	// TODO support pluto command
	var plutoResult = runKubectl("pluto detect-helm -n " + nsName)

	return plutoResult
	// return plutoVersion
}

func AnalyzeDeprecatedResources(ns *v1.NamespaceList) []ExtendedNs {
	fmt.Println("Analyzing deprecated resources in namespaces")
	
	var extendedNs []ExtendedNs

	for _, element := range ns.Items {
		current := ExtendedNs{
			Namespace:    element.GetName(),
			State:  getDeprecatedPluto(element), 
		}

		extendedNs = append(extendedNs, current)
	}

	return extendedNs
}