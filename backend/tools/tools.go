package tools

import (
	"fmt"
	_ "os/exec"
	v1 "k8s.io/api/core/v1"
)

type ExtendedNs struct {
    Namespace   string
    State		string
}

func AnalyzeDeprecatedResources(ns *v1.NamespaceList) []ExtendedNs {
	fmt.Println("Analyzing deprecated resources in namespaces")
	var extendedNs []ExtendedNs

	for _, element := range ns.Items {
		current := ExtendedNs{
			Namespace:    element.String(),
			State:   "alksdjf",
		}

		extendedNs = append(extendedNs, current)
	}

	// Pour chacun des ns, lancer les tools d'analyse de deprecation

	// Execute random command
	// cmd := "printf \"executing pluto v\"; pluto version | awk -F \"[: ]\" '{print $2}'"
	// out, err := exec.Command("/bin/sh","-c",cmd).Output()
	// if err != nil {
	// 		fmt.Println("ERROR during encryption")
	// 		fmt.Println(err)
	// 		fmt.Printf("Failed to execute command: %s\n", cmd)
	// }
	// result := string(out)
	// 	fmt.Println("--------------\n[" + result + "]\n--------------\n")

	return extendedNs
}