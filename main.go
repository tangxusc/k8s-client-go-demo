package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//kubelet.kubeconfig  是文件对应地址
	//kubeconfig := flag.String("kubeconfig", "kubelet.kubeconfig", "(optional) absolute path to the kubeconfig file")
	kubeconfig := flag.String("kubeconfig", "/home/tangxu/.kube/config", "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	//clientcmd.BuildConfigFromFlags()

	// 解析到config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建连接
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(clientset)

	deploymentsClient := clientset.AppsV1beta1().Deployments(v1.NamespaceDefault)
	deploymentList, err := deploymentsClient.List(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for key, value := range deploymentList.Items {
		fmt.Println("deployment", key, ":", value)
	}
}