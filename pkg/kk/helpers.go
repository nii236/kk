package k

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	appsv1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodListHeaders are the headers needed for the Pod list
var PodListHeaders = []string{"Namespace", "Name", "Restarts", "Age", "Ready", "Status"}

// NamespaceListHeaders are the headers needed for the Namespace list
var NamespaceListHeaders = []string{"Status", "Name", "Age"}

// DeploymentListHeaders are the headers needed for the Deployment list
var DeploymentListHeaders = []string{"Namespace", "Name", "Pods", "Age"}

// DeploymentLineHelper is the column helper for Deployments
func DeploymentLineHelper(deployment appsv1.Deployment) []string {
	pods := fmt.Sprintf("%d/%d", deployment.Status.ReadyReplicas, deployment.Status.Replicas)
	return []string{
		deployment.Namespace,
		deployment.Name,
		pods,
		columnHelperAge(deployment.ObjectMeta),
	}
}

// NamespaceLineHelper is the column helper for Namespaces
func NamespaceLineHelper(ns corev1.Namespace) []string {
	status := fmt.Sprintf("%s", ns.Status.Phase)
	return []string{
		status,
		ns.Name,
		columnHelperAge(ns.ObjectMeta),
	}
}

// PodNameFromLine returns a pods name from a line in a table
func PodNameFromLine(line string) (string, error) {
	lines := strings.Fields(line)
	if len(lines) < 1 {
		return "", errors.New("could not extra pod from line")
	}
	return lines[1], nil
}

// PodLineHelper is the column helper for Pods
func PodLineHelper(pod corev1.Pod) []string {
	return []string{
		pod.Namespace,
		pod.Name,
		columnHelperRestarts(pod),
		columnHelperAge(pod.ObjectMeta),
		columnHelperReady(pod),
		columnHelperStatus(pod),
	}
}

func columnHelperRestarts(pod corev1.Pod) string {
	cs := pod.Status.ContainerStatuses
	r := 0
	for _, c := range cs {
		r = r + int(c.RestartCount)
	}
	return strconv.Itoa(r)
}

func columnHelperAge(meta metav1.ObjectMeta) string {
	t := meta.CreationTimestamp
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

func columnHelperStatus(pod corev1.Pod) string {
	s := pod.Status
	return fmt.Sprintf("%s", s.Phase)
}

func columnHelperReady(pod corev1.Pod) string {
	cs := pod.Status.ContainerStatuses
	cr := 0
	for _, c := range cs {
		if c.Ready {
			cr = cr + 1
		}
	}
	return fmt.Sprintf("%d/%d", cr, len(cs))
}
