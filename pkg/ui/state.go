package ui

import (
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"

	"github.com/imdario/mergo"
	"github.com/jroimartin/gocui"
	"github.com/manveru/faker"
	"github.com/nii236/k"
)

func Debug(g *gocui.Gui, msg string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("Debug")
		if err != nil {
			panic(err)
		}
		fmt.Fprint(v, msg)
		return nil
	})
}

func StateLoader(g *gocui.Gui) {
	for {
		select {
		case newState := <-stateChannel:
			err := mergo.MergeWithOverwrite(newState, store)
			if err != nil {
				panic(err)
			}
			stateUpdater := SetState(newState)
			g.Update(stateUpdater)
		}
	}
}

func SetState(newState *k.State) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		store = newState
		return nil
	}
}

func LoadNamespaces(g *gocui.Gui) error {
	f, err := faker.New("en")
	if err != nil {
		panic(err)
	}
	v, err := g.View("Namespaces")
	if err != nil {
		return err
	}
	v.Clear()
	fmt.Fprint(v, strings.Join(f.Words(4, true), "\n"))
	return nil
}

func LoadResources(g *gocui.Gui) error {
	f, err := faker.New("en")
	if err != nil {
		panic(err)
	}
	v, err := g.View("Resources")
	if err != nil {
		return err
	}
	v.Clear()
	fmt.Fprint(v, strings.Join(f.Words(4, true), "\n"))
	return nil
}

func LoadPods(g *gocui.Gui) error {
	f, err := faker.New("en")
	if err != nil {
		panic(err)
	}

	result := [][]string{}
	for i := 0; i < 5; i++ {
		result = append(result, f.Words(4, true))
	}

	v, err := g.View("Pods")
	if err != nil {
		return err
	}

	v.Clear()
	t := tablewriter.NewWriter(v)
	t.SetBorder(false)
	t.SetColumnSeparator("")
	t.AppendBulk(result)
	t.Render()
	return nil
}
