package kube

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeManager struct {
	config    *rest.Config
	clientset *kubernetes.Clientset
}

// SpecKind is used to represent a Kubernetes object and its type.
type SpecKind struct {
	Spec runtime.Object
	Kind string
}

// NewManager creates a new KubeManager instance using the provided kubeconfig file path.
func NewManager(kubeConfigPath string) (*KubeManager, error) {
	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read kubeconfig file at %s: %w", kubeConfigPath, err)
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("unable to create clientset: %w", err)
	}

	ctx := context.TODO()
	if _, err = clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{}); err != nil {
		return nil, fmt.Errorf("unable to access cluster: %w", err)
	}

	return &KubeManager{
		config:    config,
		clientset: clientset,
	}, nil
}

// GetAllNodeDetails retrieves details of all nodes in the cluster.
func (km *KubeManager) GetAllNodeDetails() (*corev1.NodeList, error) {
	nodeClient := km.clientset.CoreV1().Nodes()
	return nodeClient.List(context.TODO(), metav1.ListOptions{})
}

// CreateService creates a new Service in the specified namespace.
func (km *KubeManager) CreateService(namespace string, svc *corev1.Service) (*corev1.Service, error) {
	svcClient := km.clientset.CoreV1().Services(namespace)
	return svcClient.Create(context.TODO(), svc, metav1.CreateOptions{})
}

// DeleteService deletes the specified Service in the given namespace.
func (km *KubeManager) DeleteService(namespace, svcName string) error {
	svcClient := km.clientset.CoreV1().Services(namespace)
	return svcClient.Delete(context.TODO(), svcName, metav1.DeleteOptions{})
}

// CreateDeployment creates a new Deployment in the specified namespace.
func (km *KubeManager) CreateDeployment(namespace string, dep *appsv1.Deployment) (*appsv1.Deployment, error) {
	depClient := km.clientset.AppsV1().Deployments(namespace)
	return depClient.Create(context.TODO(), dep, metav1.CreateOptions{})
}

// DeleteDeployment deletes the specified Deployment in the given namespace.
func (km *KubeManager) DeleteDeployment(namespace, depName string) error {
	depClient := km.clientset.AppsV1().Deployments(namespace)
	return depClient.Delete(context.TODO(), depName, metav1.DeleteOptions{})
}

// CreateReplicaSet creates a new ReplicaSet in the specified namespace.
func (km *KubeManager) CreateReplicaSet(namespace string, rs *appsv1.ReplicaSet) (*appsv1.ReplicaSet, error) {
	rsClient := km.clientset.AppsV1().ReplicaSets(namespace)
	return rsClient.Create(context.TODO(), rs, metav1.CreateOptions{})
}

// ScaleReplicaSet scales the specified ReplicaSet in the given namespace.
func (km *KubeManager) ScaleReplicaSet(namespace, rsName string, replicas int32) error {
	rsClient := km.clientset.AppsV1().ReplicaSets(namespace)
	scale := &autoscalingv1.Scale{
		TypeMeta: metav1.TypeMeta{
			Kind: "ReplicaSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      rsName,
			Namespace: namespace,
		},
		Spec: autoscalingv1.ScaleSpec{
			Replicas: replicas,
		},
	}
	_, err := rsClient.UpdateScale(context.TODO(), rsName, scale, metav1.UpdateOptions{})
	return err
}

// DeleteReplicaSet deletes the specified ReplicaSet in the given namespace.
func (km *KubeManager) DeleteReplicaSet(namespace, rsName string) error {
	rsClient := km.clientset.AppsV1().ReplicaSets(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	err := rsClient.Delete(context.TODO(), rsName, metav1.DeleteOptions{PropagationPolicy: &deletePolicy})
	if err != nil {
		return fmt.Errorf("error deleting replicaset: %w", err)
	}
	return nil
}

// MonitorReplicaSet monitors the status of ReplicaSets based on a label selector.
func (km *KubeManager) MonitorReplicaSet(ctx context.Context, namespace, label string) error {
	opts := metav1.ListOptions{
		LabelSelector: label,
	}
	watcher, err := km.clientset.CoreV1().Pods(namespace).Watch(ctx, opts)
	if err != nil {
		return err
	}
	for {
		select {
		case event := <-watcher.ResultChan():
			pod, ok := event.Object.(*corev1.Pod)
			if !ok {
				log.Fatal("unexpected type")
			}
			switch event.Type {
			case watch.Error:
				req := km.clientset.CoreV1().Pods(namespace).GetLogs(pod.Name, &corev1.PodLogOptions{})
				res, err := req.DoRaw(context.TODO())
				if err != nil {
					return fmt.Errorf("could not get logs from failed pod %v: %w", pod.Name, err)
				}
				return fmt.Errorf("log from failed pod %v:\n%s", pod.Name, res)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// ForceDeletePods forcefully deletes pods based on a label selector.
func (km *KubeManager) ForceDeletePods(namespace, label string) error {
	podClient := km.clientset.CoreV1().Pods(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	grace := int64(0)
	err := podClient.DeleteCollection(context.TODO(), metav1.DeleteOptions{
		GracePeriodSeconds: &grace,
		PropagationPolicy:  &deletePolicy,
	}, metav1.ListOptions{
		LabelSelector: label,
	})
	if err != nil {
		return fmt.Errorf("error force deleting pods: %w", err)
	}
	return nil
}
