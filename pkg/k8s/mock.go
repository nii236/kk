package k8s

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/manveru/faker"

	// "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/nii236/k"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockClientSet contains an embedded Kubernetes client set
type MockClientSet struct {
	faker     *faker.Faker
	clientSet *fake.Clientset
}

// New returns a new clientset
func NewMock(flags *k.ParsedFlags) (*MockClientSet, error) {
	mockClientSet := fake.NewSimpleClientset()
	mocker, err := faker.New("en")
	if err != nil {
		return nil, err
	}

	cs := &MockClientSet{
		faker:     mocker,
		clientSet: mockClientSet,
	}

	cs.seed()

	return cs, nil
}

func (cs *MockClientSet) seed() {
	for i := 0; i < 5; i++ {
		_, err := cs.clientSet.CoreV1().Pods("default").Create(&v1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: cs.faker.Name(),
				Labels: map[string]string{
					"tag": "mockdata",
				},
				CreationTimestamp: metav1.Time{
					Time: time.Now().Add(-48 * time.Hour),
				},
			},
		})
		if err != nil {
			panic(err)
		}
	}
}

// Get pods (use namespace)
func (cs *MockClientSet) GetPods(namespace string) (*v1.PodList, error) {
	return cs.clientSet.CoreV1().Pods("default").List(metav1.ListOptions{})
}

// Get namespaces
func (cs *MockClientSet) GetNamespaces() (*v1.NamespaceList, error) {
	return cs.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
}

// Get the pod containers
func (cs *MockClientSet) GetPodContainers(podName string, namespace string) []string {
	var pc []string

	pod, _ := cs.clientSet.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	for _, c := range pod.Spec.Containers {
		pc = append(pc, c.Name)
	}

	return pc
}

// Delete pod
func (cs *MockClientSet) DeletePod(podName string, namespace string) error {
	return cs.clientSet.CoreV1().Pods(namespace).Delete(podName, &metav1.DeleteOptions{})
}

// Get pod container logs
func (cs *MockClientSet) GetPodContainerLogs(podName string, containerName string, namespace string, o io.Writer) error {
	tl := int64(50)

	opts := &v1.PodLogOptions{
		Container: containerName,
		TailLines: &tl,
	}

	req := cs.clientSet.CoreV1().Pods(namespace).GetLogs(podName, opts)

	readCloser, err := req.Stream()
	if err != nil {
		return err
	}

	_, err = io.Copy(o, readCloser)

	readCloser.Close()

	return err
}

// Column helper: Restarts
func (cs *MockClientSet) ColumnHelperRestarts(containerStatuses []v1.ContainerStatus) string {
	r := 0
	for _, c := range containerStatuses {
		r = r + int(c.RestartCount)
	}
	return strconv.Itoa(r)
}

// Column helper: Age
func (cs *MockClientSet) ColumnHelperAge(t metav1.Time) string {
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
func (cs *MockClientSet) ColumnHelperStatus(s v1.PodStatus) string {
	return fmt.Sprintf("%s", s.Phase)
}

// Column helper: Ready
func (cs *MockClientSet) ColumnHelperReady(containerStatuses []v1.ContainerStatus) string {
	cr := 0
	for _, c := range containerStatuses {
		if c.Ready {
			cr = cr + 1
		}
	}
	return fmt.Sprintf("%d/%d", cr, len(containerStatuses))
}
