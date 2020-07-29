package tools

import (
	"fmt"
	"errors"
	"strings"
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

func analyzeError(out string, err error, cmd string) string {

	switch errMsg := err.Error(); errMsg {
	case errors.New("exit status 127").Error():
		fmt.Printf("Failed to execute command: [%s]\n", cmd)
		fmt.Printf("Error. Result : \n" + out)
		return "Error. Command not found"
	case errors.New("exit status 3").Error():
		fmt.Printf("Warning. Result : \n" + out)
		return "Exit status 3"
	default:
		fmt.Printf("Failed to execute command: [%s]\n", cmd)
		fmt.Println("Error : " + err.Error())
		return "Unknown error"
	}
}

func runPlutoHelm(ns string) PlutoItems {

	// Computing plutoVersion
	// var plutoVersion = runCmd("pluto version | awk -F \"[: ]\" '{print $2}'")

	cmd := "pluto detect-helm -f tmp_deployments/policies/pluto-custom-deprecated.yaml -o json -n " + ns
	// cmd_config := exec.Command("/bin/bash", "-c", cmd)

	app := "plutos"
	args := []string{"detect-helm", "-f", "tmp_deployments/policies/pluto-custom-deprecated.yaml", "-o", "json", "-n", ns}
	cmd = app + " " + strings.Join(args, " ")
	cmd_config := exec.Command(app, args...)
	cmd_config.Env = append(cmd_config.Env, "KUBECONFIG="+ kublink.KUBECONFIG)
	cmd_config.Env = append(cmd_config.Env, "USER="+ kublink.USER)
	cmd_config.Env = append(cmd_config.Env, "HOME="+ kublink.HOME)
	out, err := cmd_config.Output()

	if err != nil {
		fmt.Println(analyzeError(string(out), err, cmd))
	} 

	data := &PlutoItems{}
	err = json.Unmarshal(out, data)
	fmt.Printf("%+v\n", data)

	return *data
}

func getDeprecatedPluto(ns v1.Namespace) PlutoItems {
	nsName := ns.ObjectMeta.Name
	fmt.Println("[Checking resources in ns " + nsName + "]")
	
	var plutoResult = runPlutoHelm(nsName)

	fmt.Println()
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