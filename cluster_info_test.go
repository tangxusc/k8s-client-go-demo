package main

import (
	"encoding/json"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestGetAll(t *testing.T) {

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

	nss, e := getAllNamespace(client)
	if e != nil {
		println(e.Error())
		return
	}

	for _, value := range nss {
		fmt.Println("开始获取命名空间:", value.Name, "中的所有信息")
		result, e := getAllPod(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnPod(result)
		depts, e := getAllDeployment(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnDepts(depts)
		results, e := getAllStatefulSet(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnStatefuls(results)
		daemonsetList, e := getAllDaemonSet(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnDaemonSetList(daemonsetList)
		jobs, e := getAllJob(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnJobs(jobs)
		cronjobs, e := getAllCronJob(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnCronjobs(cronjobs)
		//getAllConfigMap(value, client)
		secrets, e := getAllSecret(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnSecrets(secrets)
		//getAllPv(value)
		//getAllPvc(value)
		//getAllStorageClass(value)
		//
		services, e := getAllService(value, client)
		if e != nil {
			println(e.Error())
			return
		}
		printlnServices(services)
		//getAllServiceAccount(value)

	}
}

func printlnServices(services []corev1.Service) {
	for _, value := range services {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("services:%s,\t %s \n", value.Name, string(bytes))
	}
}

func getAllService(namespace corev1.Namespace, client *kubernetes.Clientset) (services []corev1.Service, err error) {
	serviceList, e := client.CoreV1().Services(namespace.Name).List(v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	return serviceList.Items, nil
}

func printlnSecrets(secrets []corev1.Secret) {
	for _, value := range secrets {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("secrets:%s,\t %s \n", value.Name, string(bytes))
	}
}

func printlnCronjobs(jobs []batchv1beta1.CronJob) {
	for _, value := range jobs {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("cronJob:%s,\t %s \n", value.Name, string(bytes))
	}
}

func getAllSecret(namespace corev1.Namespace, client *kubernetes.Clientset) (secrets []corev1.Secret, err error) {
	secretList, e := client.CoreV1().Secrets(namespace.Name).List(v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	return secretList.Items, nil
}

func getAllCronJob(namespace corev1.Namespace, client *kubernetes.Clientset) (cronjobs []batchv1beta1.CronJob, err error) {
	job, e := client.BatchV1beta1().CronJobs(namespace.Name).List(v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	return job.Items, nil
}

func printlnJobs(jobs []batchv1.Job) {
	for _, value := range jobs {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("job:%s,\t %s \n", value.Name, string(bytes))
	}
}

func getAllJob(namespace corev1.Namespace, client *kubernetes.Clientset) (jobs []batchv1.Job, err error) {
	jobList, e := client.BatchV1().Jobs(namespace.Name).List(v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	return jobList.Items, nil
}

func printlnDaemonSetList(sets []appv1.DaemonSet) {
	for _, value := range sets {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("statefulset:%s,\t %s \n", value.Name, string(bytes))
	}
}

func getAllDaemonSet(namespace corev1.Namespace, client *kubernetes.Clientset) (daemonsetList []appv1.DaemonSet, err error) {
	daemonSetList, e := client.AppsV1().DaemonSets(namespace.Name).List(v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	return daemonSetList.Items, nil
}

func printlnStatefuls(sets []appv1.StatefulSet) {
	for _, value := range sets {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("statefulset:%s,\t %s \n", value.Name, string(bytes))
	}
}

func getAllStatefulSet(namespace corev1.Namespace, clientset *kubernetes.Clientset) (result []appv1.StatefulSet, err error) {
	setList, e := clientset.AppsV1().StatefulSets(namespace.Name).List(v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	return setList.Items, nil
}

func printlnDepts(deployments []appv1.Deployment) {
	for _, value := range deployments {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("dept:%s,\t %s \n", value.Name, string(bytes))
	}
}

func getAllDeployment(namespace corev1.Namespace, client *kubernetes.Clientset) (depts []appv1.Deployment, err error) {
	deploymentList, err := client.AppsV1().Deployments(namespace.Name).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return deploymentList.Items, nil
}

func printlnPod(pods []corev1.Pod) {
	for _, value := range pods {
		bytes, e := json.Marshal(value)
		if e != nil {
			println(e.Error())
			continue
		}
		fmt.Printf("pod:%s,image:%s \t %s \n", value.Name, value.Spec.Containers[0].Image, string(bytes))
	}
}

func getAllPod(namespace corev1.Namespace, clientset *kubernetes.Clientset) (result []corev1.Pod, err error) {
	pods := clientset.CoreV1().Pods(namespace.Name)
	podList, e := pods.List(v1.ListOptions{})
	if err != nil {
		return nil, e
	}
	return podList.Items, nil
}

func getAllNamespace(client *kubernetes.Clientset) (result []corev1.Namespace, error error) {
	namespaceList, e := client.CoreV1().Namespaces().List(v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	return namespaceList.Items, nil
}
