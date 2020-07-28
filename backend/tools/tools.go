package tools

import (
	"fmt"
	"os"
	"os/exec"
	"encoding/json"
	v1 "k8s.io/api/core/v1"
	"github.com/wayoos/kubedash/backend/kublink"
	// "github.com/wayoos/kubedash/tools/struct"
)

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

func runPlutoHelm(ns string) PlutoItems {

	// var plutoVersion = runCmd("pluto version | awk -F \"[: ]\" '{print $2}'")

	// cmd := "pluto detect-files -o json"
	cmd := "pluto detect-helm -f tmp_deployments/policies/pluto-custom-deprecated.yaml -o json -n " + ns
	cmd_config := exec.Command("/bin/bash", "-c", cmd)
	cmd_config.Env = append(cmd_config.Env, "KUBECONFIG="+ kublink.KUBECONFIG)
	cmd_config.Env = append(cmd_config.Env, "USER="+ kublink.USER)
	cmd_config.Env = append(cmd_config.Env, "HOME="+ kublink.HOME)
	out, err := cmd_config.Output()

	if err != nil {
			fmt.Printf("Failed to execute pluto command: [%s]\n", cmd)
			fmt.Println(err)
	} 

	fmt.Println("Pluto executed")
	fmt.Println(string(out))

	data := &PlutoItems{}
	err = json.Unmarshal(out, data)

	fmt.Println("Struct created")
	fmt.Println(data) 

	return *data
}

func getDeprecatedPluto(ns v1.Namespace) PlutoItems {
	nsName := ns.ObjectMeta.Name
	fmt.Println("\n[Checking resources in ns " + nsName + "]")
	
	var plutoResult = runPlutoHelm(nsName)

	return plutoResult
}

func AnalyzeDeprecatedResources(ns *v1.NamespaceList) []ExtendedNs {
	fmt.Println("Analyzing deprecated resources in namespaces")
	
	var extendedNs []ExtendedNs

	for _, element := range ns.Items {
		current := ExtendedNs{
			Namespace:    element.GetName(),
			StateHelm:  getDeprecatedPluto(element), 
		}

		extendedNs = append(extendedNs, current)
	}

	return extendedNs
}