package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/wayoos/kubedash/backend/deprecated"
	"github.com/wayoos/kubedash/backend/kublink"
	"github.com/wayoos/kubedash/backend/namespace"
)

var KUBECONFIG = "/Users/ralbasini/.config/k3d/k3s-default/tmp.yaml"

// var KUBECONFIG = "./kubeconfig.yaml" // For docker-compose

func main() {

	fmt.Println("KUBECONFIG from main: " + KUBECONFIG)

	// access the API to list pods
	//	pods, _ := clientset.CoreV1().Pods("").List(v1.ListOptions{})
	//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	kublink.Connect(KUBECONFIG)

	e := echo.New()
	e.Static("/", "frontend/dist")

	apiGroup := e.Group("/api")

	namespace.Routing(*apiGroup)
	deprecated.Routing(*apiGroup)

	e.Logger.Fatal(e.Start("127.0.0.1:8008"))
}
