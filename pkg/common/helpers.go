package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/jroimartin/gocui"

	"github.com/nii236/k"
	v1 "k8s.io/api/core/v1"
)

func JSONToState(g *gocui.Gui) (*k.State, error) {
	stateBuf, err := g.View("State")
	if err != nil && err != gocui.ErrUnknownView {
		return nil, errors.Wrap(err, "Could not get view")
	}
	if err == gocui.ErrUnknownView {
		return nil, errors.Wrap(err, "state not initialized")
	}
	s := &k.State{
		Entities: &k.EntitiesReducer{
			Pods: &k.PodEntities{},
		},
		UI: &k.UIReducer{},
	}
	err = json.Unmarshal([]byte(stateBuf.Buffer()), s)
	if err != nil {
		return nil, errors.Wrap(err, "Could not decode JSON")
	}
	return s, nil
}

// Column helper: Restarts
func ColumnHelperRestarts(pod v1.Pod) string {
	cs := pod.Status.ContainerStatuses
	r := 0
	for _, c := range cs {
		r = r + int(c.RestartCount)
	}
	return strconv.Itoa(r)
}

// Column helper: Age
func ColumnHelperAge(pod v1.Pod) string {
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
func ColumnHelperStatus(pod v1.Pod) string {
	s := pod.Status
	return fmt.Sprintf("%s", s.Phase)
}

// Column helper: Ready
func ColumnHelperReady(pod v1.Pod) string {
	cs := pod.Status.ContainerStatuses
	cr := 0
	for _, c := range cs {
		if c.Ready {
			cr = cr + 1
		}
	}
	return fmt.Sprintf("%d/%d", cr, len(cs))
}
