package images

const (
	DefaultImageRegistry       string = "quay.io/skupper"
	RouterImageName            string = "skupper-router:2.7.1"
	ServiceControllerImageName string = "service-controller:1.8.1"
	ControllerPodmanImageName  string = "controller-podman:1.8.1"
	ConfigSyncImageName        string = "config-sync:1.8.1"
	FlowCollectorImageName     string = "flow-collector:1.8.1"
	SiteControllerImageName    string = "site-controller:1.8.1"
	PrometheusImageRegistry    string = "quay.io/prometheus"
	PrometheusServerImageName  string = "prometheus:v2.42.0"
	OauthProxyImageRegistry    string = "quay.io/openshift"
	OauthProxyImageName        string = "origin-oauth-proxy:4.14.0"
)
