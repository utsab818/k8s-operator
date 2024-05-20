package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	klient "github.com/utsab818/kluster/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// "github.com/utsab818/kluster/pkg/apis/utsab.dev/v1alpha1"

func main() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Printf("Building config from flags failed, %s, trying to build inclusterconfig", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Printf("error %s building inclusterconfig", err.Error())
		}
	}

	klientset, err := klient.NewForConfig(config)
	if err != nil {
		log.Printf("getting klient set %s\n", err.Error())
	}

	// fmt.Println(klientset)

	klusters, err := klientset.UtsabV1alpha1().Klusters("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("listing klusters %s\n", err.Error())
	}

	fmt.Printf("length of kluster is %d and name is %s\n", len(klusters.Items), klusters.Items[0].Name)
}