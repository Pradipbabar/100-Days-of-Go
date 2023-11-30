package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/path/to/kubeconfig", "absolute path to the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	// Example: List all pods in the default namespace
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing pods: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Pods in the default namespace:")
	for _, pod := range pods.Items {
		fmt.Printf(" - %s\n", pod.GetName())
	}

	// Add more interactions as needed...

	fmt.Println("Successfully interacted with the Kubernetes cluster!")
}


