package deprecated

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/wayoos/kubedash/backend/kublink"
	"github.com/wayoos/kubedash/backend/tools"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Person struct {
    Fn string
    Ln string
}

// Routing function
func Routing(g echo.Group) {
	g.GET("/deprecated", deprecated)
}

func deprecated(c echo.Context) error {

	var ns *v1.NamespaceList
	ns, _ = kublink.ClientSet().CoreV1().Namespaces().List(metav1.ListOptions{})

	tools.AnalyzeDeprecatedResources(ns)

	return c.JSON(http.StatusOK, ns.Items)
}
