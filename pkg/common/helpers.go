package common

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/jroimartin/gocui"

	"github.com/nii236/k"
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
