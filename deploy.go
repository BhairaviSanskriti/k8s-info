package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentData struct {
	Name              string `json:"name"`
	HealthyReplicas   int32  `json:"healthy_replicas"`
	UnhealthyReplicas int32  `json:"unhealthy_replicas"`
}

func runDeploy(cmd *cobra.Command, args []string) {
	config, err := getClientConfig(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	data := []DeploymentData{}

	for _, deployment := range deployments.Items {
		healthyReplicas := deployment.Status.AvailableReplicas
		unhealthyReplicas := *deployment.Spec.Replicas - deployment.Status.AvailableReplicas

		data = append(data, DeploymentData{
			Name: deployment.Name,

			UnhealthyReplicas: unhealthyReplicas,
			HealthyReplicas:   healthyReplicas,
		})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
