package k

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/jroimartin/gocui"
)

func JSONToState(g *gocui.Gui) (*State, error) {
	stateBuf, err := g.View("State")
	if err != nil && err != gocui.ErrUnknownView {
		return nil, errors.Wrap(err, "Could not get view")
	}
	if err == gocui.ErrUnknownView {
		return nil, errors.Wrap(err, "state not initialized")
	}
	s := &State{
		Entities: &EntitiesReducer{
			Pods: &PodEntities{},
		},
		UI: &UIReducer{},
	}
	err = json.Unmarshal([]byte(stateBuf.Buffer()), s)
	if err != nil {
		return nil, errors.Wrap(err, "Could not decode JSON")
	}
	return s, nil
}
