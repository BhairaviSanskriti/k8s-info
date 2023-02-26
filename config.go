package main

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	//check config file path in $KUBECONFIG variable
	kubeconfig = os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	//check config file in ~/.kube/config
	kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	//check config file in the current directory
	kubeconfig = filepath.Join(os.Getenv("PWD"), "kubeconfig")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	// set KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT in case the function uses in-cluster info
	return rest.InClusterConfig()
}
