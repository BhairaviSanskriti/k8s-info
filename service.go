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

type ServiceData struct {
	Name  string  `json:"service_name"`
	Type  string  `json:"type"`
	Ports []int32 `json:"exposed_ports"`
}

func runService(cmd *cobra.Command, args []string) {
	config, err := getClientConfig(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	
	services, err := clientset.CoreV1().Services(namespace).List(context.Background(), v1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	data := []ServiceData{}

	for _, service := range services.Items {
		ports := service.Spec.Ports
		var p []int32
		for _, port := range ports {
			p = append(p, port.Port)
		}
		data = append(data, ServiceData{
			Name:  service.Name,
			Type:  string(service.Spec.Type),
			Ports: p,
		})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
