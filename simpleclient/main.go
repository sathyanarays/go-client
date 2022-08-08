package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig = flag.String("kubeconfig", "/Users/sathyanarays/.kube/config", "kubeconfig file")

func main() {
	fmt.Println("Hello World")

	var config *rest.Config
	config, err := rest.InClusterConfig()

	if err != nil {
		fmt.Println("Error getting incluster config")
		flag.Parse()
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)

		if err != nil {
			fmt.Errorf("Error", err)
		}
	} else {
		fmt.Println("Got incluster config")
	}

	//fmt.Println(config)
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Errorf("Error", err)
	}

	pod, err := clientset.CoreV1().Pods("default").Get(context.Background(), "simple-app-59b56fb7f9-c62sv", metav1.GetOptions{})

	if err != nil {
		fmt.Errorf("Error", err)
	}

	fmt.Println(pod)

	for {

	}
}
