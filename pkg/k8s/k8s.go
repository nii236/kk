package k8s

import (
	"io"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	appsv1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

type ClientSet interface {
	GetPods(namespace string) (*corev1.PodList, error)
	GetDeployments(namespace string) (*appsv1.DeploymentList, error)
	GetNamespaces() (*corev1.NamespaceList, error)
	GetPodContainers(podName string, namespace string) []string
	DeletePod(podName string, namespace string) error
	GetPodContainerLogs(podName string, containerName string, namespace string, o io.Writer) error
}
