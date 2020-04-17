module github.com/wayoos/kubedash

require (
	github.com/labstack/echo/v4 v4.1.16
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v12.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator

go 1.14
