package k8s

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/nii236/k"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClientSet contains an embedded Kubernetes client set
type ClientSet struct {
	clientset *kubernetes.Clientset
}

// New returns a new clientset
func New(flags *k.ParsedFlags) (*ClientSet, error) {
	// Use the current context in kubeconfig
	cc, err := clientcmd.BuildConfigFromFlags("", flags.KubeConfigPath)
	if err != nil {
		return nil, err
	}

	// Create the client set
	clientSet, err := kubernetes.NewForConfig(cc)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		clientSet,
	}, nil
}

// Get pods (use namespace)
func (cs *ClientSet) GetPods(namespace string) (*v1.PodList, error) {
	return cs.clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}

// Get namespaces
func (cs *ClientSet) GetNamespaces() (*v1.NamespaceList, error) {
	return cs.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
}

// Get the pod containers
func (cs *ClientSet) GetPodContainers(podName string, namespace string) []string {
	var pc []string

	pod, _ := cs.clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	for _, c := range pod.Spec.Containers {
		pc = append(pc, c.Name)
	}

	return pc
}

// Delete pod
func (cs *ClientSet) DeletePod(podName string, namespace string) error {
	return cs.clientset.CoreV1().Pods(namespace).Delete(podName, &metav1.DeleteOptions{})
}

// Get pod container logs
func (cs *ClientSet) GetPodContainerLogs(podName string, containerName string, namespace string, o io.Writer) error {
	tl := int64(50)

	opts := &v1.PodLogOptions{
		Container: containerName,
		TailLines: &tl,
	}

	req := cs.clientset.CoreV1().Pods(namespace).GetLogs(podName, opts)

	readCloser, err := req.Stream()
	if err != nil {
		return err
	}

	_, err = io.Copy(o, readCloser)

	readCloser.Close()

	return err
}

// Column helper: Restarts
func (cs *ClientSet) ColumnHelperRestarts(containerStatuses []v1.ContainerStatus) string {
	r := 0
	for _, c := range containerStatuses {
		r = r + int(c.RestartCount)
	}
	return strconv.Itoa(r)
}

// Column helper: Age
func (cs *ClientSet) ColumnHelperAge(t metav1.Time) string {
	d := time.Now().Sub(t.Time)

	if d.Hours() > 1 {
		if d.Hours() > 24 {
			ds := float64(d.Hours() / 24)
			return fmt.Sprintf("%.0fd", ds)
		} else {
			return fmt.Sprintf("%.0fh", d.Hours())
		}
	} else if d.Minutes() > 1 {
		return fmt.Sprintf("%.0fm", d.Minutes())
	} else if d.Seconds() > 1 {
		return fmt.Sprintf("%.0fs", d.Seconds())
	}

	return "?"
}

// Column helper: Status
func (cs *ClientSet) ColumnHelperStatus(s v1.PodStatus) string {
	return fmt.Sprintf("%s", s.Phase)
}

// Column helper: Ready
func (cs *ClientSet) ColumnHelperReady(containerStatuses []v1.ContainerStatus) string {
	cr := 0
	for _, c := range containerStatuses {
		if c.Ready {
			cr = cr + 1
		}
	}
	return fmt.Sprintf("%d/%d", cr, len(containerStatuses))
}
