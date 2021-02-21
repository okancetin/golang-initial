package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace")
	var (
		kubeconfig = flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		namespace  = flag.String("namespace", "default", "(optional) namespace of the deployment")
	)
	flag.Parse()
	fmt.Printf("config file : %s", kubeconfig)
	config, error := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if error != nil {
		fmt.Printf("error running BuildConfigFromFlags %s ", error)
	}
	clientset, error := kubernetes.NewForConfig(config)
	if error != nil {
		fmt.Printf("error running NewForConfig %s ", error)
	}

	pods, error := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if error != nil {
		fmt.Printf("error running CoreV1 %s ", error)
	}

	for i, pod := range pods.Items {
		fmt.Println(pod.GetName)
	}
}
