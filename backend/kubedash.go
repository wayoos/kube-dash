package main

import (
	"github.com/labstack/echo/v4"

	"github.com/wayoos/kubedash/backend/kublink"
	"github.com/wayoos/kubedash/backend/namespace"
)

func main() {

	// access the API to list pods
	//	pods, _ := clientset.CoreV1().Pods("").List(v1.ListOptions{})
	//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	kublink.Connect()

	e := echo.New()
	e.Static("/", "frontend/dist")

	apiGroup := e.Group("/api")

	namespace.Routing(*apiGroup)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
