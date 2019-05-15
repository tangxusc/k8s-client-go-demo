package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestK8sClient(t *testing.T) {
	//kubelet.kubeconfig  是文件对应地址
	kubeconfig := flag.String("kubeconfig", "/home/tangxu/.kube/config", "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	// 解析到config
	config, err := clientcmd.BuildConfigFromFlags("10.30.21.238:6443", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建连接
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//deploymentsClient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	deploymentsClient := clientset.AppsV1beta1().Deployments("kube-system")
	deploymentList, err := deploymentsClient.List(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for key, value := range deploymentList.Items {
		fmt.Println("deployment", key)
		bytes, err := json.Marshal(value)
		if err != nil {
			println(err.Error())
			return
		}
		fmt.Println("内容为", string(bytes))
	}
	fmt.Println("============获取pod==============")

	pods := clientset.CoreV1().Pods("kube-system")
	podList, err := pods.List(v1.ListOptions{})
	if err != nil {
		println(err.Error())
		return
	}

	for key, value := range podList.Items {
		fmt.Println("第", key, "个pod.................")
		bytes, err := json.Marshal(value)
		if err != nil {
			return
		}
		fmt.Println(string(bytes))
	}

}
