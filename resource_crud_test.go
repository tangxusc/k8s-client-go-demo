package main

import (
	"encoding/json"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
	"time"
)

func TestResource(t *testing.T) {
	config, e := clientcmd.BuildConfigFromFlags("10.30.21.238:6443", "/home/tangxu/.kube/config")
	if e != nil {
		println(e.Error())
		return
	}
	client, e := kubernetes.NewForConfig(config)
	if e != nil {
		println(e.Error())
		return
	}

	//podInterface := client.CoreV1().Pods("kube-system")
	//podList, e := podInterface.List(metav1.ListOptions{
	//	LabelSelector: "app=helm",
	//})
	//if e != nil {
	//	println(e.Error())
	//}
	//for _, value := range podList.Items {
	//	bytes, e := json.Marshal(value)
	//	if e != nil {
	//		println(e.Error())
	//		continue
	//	}
	//	fmt.Println("pod", value.Name, string(bytes))
	//}
	//pod := podList.Items[0]
	//logs := podInterface.GetLogs(pod.Name, &corev1.PodLogOptions{
	//	Follow:   true,
	//	Previous: true,
	//})
	//do := logs.Do()
	//e = do.Error()
	//fmt.Println("do.Error", e)
	//bytes, e := do.Raw()
	//fmt.Println(string(bytes), e)
	//
	//return

	//client.RESTClient().Get().Do().Into()
	//newClient, e := dynamic.NewClient(config)
	//newClient.Resource(nil, "test").Get()
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	fmt.Println("开始创建命名空间...")
	create, e := client.CoreV1().Namespaces().Create(namespace)
	//statusError := e.(*errors.StatusError)
	//if errors.IsAlreadyExists(e) {
	//	fmt.Println("exists")
	//}
	var rep int32 = 1
	fmt.Println("创建命名空间完成,返回结果:", create, e)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test",
			Name:      "test",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"test": "test",
				},
			},
			Replicas: &rep,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
					Name:      "test",
					Labels: map[string]string{
						"test": "test",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Image:           "nginx",
							Name:            "test",
							ImagePullPolicy: corev1.PullIfNotPresent,
						},
					},
				},
			},
		},
	}
	fmt.Println("开始创建deployment...")
	dept, e := client.AppsV1().Deployments("test").Create(deployment)
	fmt.Println(dept, e)

	deploymentList, e := client.AppsV1().Deployments("test").List(metav1.ListOptions{})
	if e != nil {
		println(e.Error())
	}
	for _, value := range deploymentList.Items {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Println("dept", value.Name, string(bytes))
	}

	time.Sleep(20 * time.Second)
	fmt.Println("清理ns,deployment...")
}

//https://github.com/kubernetes/client-go/blob/master/examples/workqueue/main.go#L174
//https://github.com/kubernetes/sample-controller/blob/master/controller.go#L87:6
//https://yq.aliyun.com/articles/688485#
//https://blog.csdn.net/weixin_42663840/article/details/81980022
//https://blog.csdn.net/weixin_42663840/article/details/81530606
//https://blog.csdn.net/weixin_42663840/article/details/81699303
func TestInformer(t *testing.T) {
	config, e := clientcmd.BuildConfigFromFlags("10.30.21.238:6443", "/home/tangxu/.kube/config")
	if e != nil {
		println(e.Error())
		return
	}
	client, e := kubernetes.NewForConfig(config)
	if e != nil {
		println(e.Error())
		return
	}
	stopChan := make(chan struct{})
	factory := informers.NewSharedInformerFactory(client, 30*time.Second)
	podInformer := factory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(&EventHandler{})
	//ret, e := podInformer.Lister().List(labels.Nothing())

	//informer, e := factory.ForResource(schema.GroupVersionResource{
	//	Group:    "extensions",
	//	Version:  "v1beta1",
	//	Resource: "Deployment",
	//})
	//if e != nil {
	//	println(e.Error())
	//}
	//ret, e := informer.Lister().ByNamespace("kube-system").List(labels.Nothing())
	//fmt.Println(ret,e)

	//informer.Informer().AddEventHandler()
	//factory.InformerFor(&corev1.Pod{}, func(inter kubernetes.Interface, duration time.Duration) cache.SharedIndexInformer {
	//	return nil
	//})
	factory.Start(stopChan)
	TestResource(t)

	stopChan <- EventHandler{}

}

type EventHandler struct {
}

func (*EventHandler) OnAdd(obj interface{}) {
	fmt.Println("OnAdd")
	bytes, _ := json.Marshal(obj)
	fmt.Println(string(bytes))
}

func (*EventHandler) OnUpdate(oldObj, newObj interface{}) {
	fmt.Println("OnUpdate")
	bytes1, _ := json.Marshal(oldObj)
	bytes2, _ := json.Marshal(newObj)
	fmt.Println(string(bytes1))
	fmt.Println(string(bytes2))
}

func (*EventHandler) OnDelete(obj interface{}) {
	fmt.Println("OnDelete")
	bytes, _ := json.Marshal(obj)
	fmt.Println(string(bytes))
}
