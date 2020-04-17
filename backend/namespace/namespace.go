package namespace

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wayoos/kubedash/backend/kublink"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Routing function
func Routing(g echo.Group) {
	g.GET("/namespaces", namespaces)
}

func namespaces(c echo.Context) error {
	var ns *v1.NamespaceList
	ns, _ = kublink.ClientSet().CoreV1().Namespaces().List(metav1.ListOptions{})

	return c.JSON(http.StatusOK, ns.Items)
}
