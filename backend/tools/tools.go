package tools

import (
	"fmt"
	"os/exec"
	v1 "k8s.io/api/core/v1"
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

func getDeprecatedPluto(ns v1.Namespace) string {
	nsName := ns.ObjectMeta.Name
	fmt.Println("\n[Checking resources in ns " + nsName + "]")

	// runCmd(export KUBECONFIG=\"$(k3d get-kubeconfig --name='k3s-default')\")
	// fmt.Println(runCmd("kubectl version"))

	var version = runCmd("pluto version | awk -F \"[: ]\" '{print $2}'")
	fmt.Println(version)

	// var plutoResult = runCmd("pluto detect-helm -n " + nsName)
	// fmt.Println(plutoResult)

	return version
}

func AnalyzeDeprecatedResources(ns *v1.NamespaceList) []ExtendedNs {
	fmt.Println("Analyzing deprecated resources in namespaces")
	
	var extendedNs []ExtendedNs

	for _, element := range ns.Items {
		current := ExtendedNs{
			Namespace:    element.String(),
			State:  getDeprecatedPluto(element), 
		}

		extendedNs = append(extendedNs, current)
	}

	return extendedNs
}