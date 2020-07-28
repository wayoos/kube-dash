package tools

type ExtendedNs struct {
    Namespace   string
    State		PlutoItems
}

type API struct {
	Component       string `json:"component"`
	Deprecated_in   string `json:"deprecated-in"`
	Kind            string `json:"kind"`
	Removed_in      string `json:"removed-in"`
	Replacement_api string `json:"replacement-api"`
	Version         string `json:"version"`
}

type PlutoItems struct {
	Items []struct {
		API struct {
			Component       string `json:"component"`
			Deprecated_in   string `json:"deprecated-in"`
			Kind            string `json:"kind"`
			Removed_in      string `json:"removed-in"`
			Replacement_api string `json:"replacement-api"`
			Version         string `json:"version"`
		} `json:"api"`
		Deprecated bool   `json:"deprecated"`
		Name       string `json:"name"`
		Removed    bool   `json:"removed"`
	} `json:"items"`
	Target_versions struct {
		Cert_manager string `json:"cert-manager"`
		Istio        string `json:"istio"`
		K8s          string `json:"k8s"`
	} `json:"target-versions"`
}
