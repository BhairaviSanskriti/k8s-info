package main

import (
	"log"

	"github.com/spf13/cobra"
)

var namespace string
var kubeconfig string

var rootCmd = &cobra.Command{
	Use:   "k8s-info",
	Short: "Retrieve deployment or service data from a Kubernetes cluster",
}
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display deployent or service data",
}
var serviceCmd = &cobra.Command{
	Use:   "svc",
	Short: "Retrieve service data from a Kubernetes cluster",
	Run:   runService,
}
var deploymentCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Retrieve deployment data from a Kubernetes cluster",
	Run:   runDeploy,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "the namespace to retrieve deployments from")
	rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "absolute path to the kubeconfig file")

	getCmd.AddCommand(deploymentCmd)
	getCmd.AddCommand(serviceCmd)

	rootCmd.AddCommand(getCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
