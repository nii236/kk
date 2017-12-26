package k8s

import (
	"io"
	"strings"
	"time"

	"github.com/manveru/faker"

	// "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/nii236/k/pkg/k"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

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

func (cs *MockClientSet) seedNamespaces() error {
	for i := 0; i < 5; i++ {
		_, err := cs.clientSet.CoreV1().Namespaces().Create(&corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: strings.Join(cs.faker.Words(2, true), "-"),
				Labels: map[string]string{
					"tag": "mockdata",
				},
				CreationTimestamp: metav1.Time{
					Time: time.Now().Add(-48 * time.Hour),
				},
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (cs *MockClientSet) seedPod(ns, name string) error {
	// fmt.Println(ns)
	// fmt.Println(name)
	_, err := cs.clientSet.CoreV1().Pods(ns).Create(&corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"tag": "mockdata",
			},
			CreationTimestamp: metav1.Time{
				Time: time.Now().Add(-48 * time.Hour),
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (cs *MockClientSet) seedPods() error {
	namespaces, err := cs.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, ns := range namespaces.Items {
		for i := 0; i < 5; i++ {
			cs.seedPod(ns.Name, strings.Join(cs.faker.Words(3, true), "-"))
		}
	}
	// os.Exit(1)
	return nil
}

func (cs *MockClientSet) seedDeployments() error {
	for i := 0; i < 5; i++ {
		_, err := cs.clientSet.AppsV1().Deployments("default").Create(&appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
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
			return err
		}
	}
	return nil
}
func (cs *MockClientSet) seed() error {
	err := cs.seedNamespaces()
	if err != nil {
		return err
	}
	err = cs.seedDeployments()
	if err != nil {
		return err
	}
	err = cs.seedPods()
	if err != nil {
		return err
	}
	return nil
}

// Get pods (use namespace)
func (cs *MockClientSet) GetPods(namespace string) (*corev1.PodList, error) {
	return cs.clientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}

// Get namespaces
func (cs *MockClientSet) GetNamespaces() (*corev1.NamespaceList, error) {
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

	opts := &corev1.PodLogOptions{
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
