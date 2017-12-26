package k

import (
	"fmt"
	"strconv"
	"time"

	v1 "k8s.io/api/core/v1"
)

func PodLineHelper(pod v1.Pod) []string {
	return []string{
		pod.Namespace,
		pod.Name,
		columnHelperRestarts(pod),
		columnHelperAge(pod),
		columnHelperReady(pod),
		columnHelperStatus(pod),
	}
}

// Column helper: Restarts
func columnHelperRestarts(pod v1.Pod) string {
	cs := pod.Status.ContainerStatuses
	r := 0
	for _, c := range cs {
		r = r + int(c.RestartCount)
	}
	return strconv.Itoa(r)
}

// Column helper: Age
func columnHelperAge(pod v1.Pod) string {
	t := pod.CreationTimestamp
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
func columnHelperStatus(pod v1.Pod) string {
	s := pod.Status
	return fmt.Sprintf("%s", s.Phase)
}

// Column helper: Ready
func columnHelperReady(pod v1.Pod) string {
	cs := pod.Status.ContainerStatuses
	cr := 0
	for _, c := range cs {
		if c.Ready {
			cr = cr + 1
		}
	}
	return fmt.Sprintf("%d/%d", cr, len(cs))
}
