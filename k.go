package k

import (
	"errors"

	"github.com/urfave/cli"
)

var PodListHeaders = []string{"Namespace", "Name", "Restarts", "Age", "Ready", "Status"}
var NamespaceListHeaders = []string{"Name"}

type Kind string

func (k Kind) String() string {
	return string(k)
}

const (
	KindPods       Kind = "Pods"
	KindNamespaces Kind = "Namespaces"
	KindResources  Kind = "Resources"
)

type Screen string

const (
	ScreenTable Screen = "Table"
	ScreenModal Screen = "Modal"
	ScreenState Screen = "State"
	ScreenDebug Screen = "Debug"
)

func (s Screen) String() string {
	return string(s)
}

type State struct {
	UI       *UIReducer
	Entities *EntitiesReducer
}

// ParsedFlags will contain the config for the app
type ParsedFlags struct {
	KubeConfigPath   string
	RefreshFrequency int
	PROD             bool
	DEBUG            bool
	TEST             bool
}

// Parse will parse the flags into a struct
func (flags *ParsedFlags) Parse(c *cli.Context) error {
	flags.KubeConfigPath = c.String("kubeconfig-path")
	flags.RefreshFrequency = c.Int("refresh-frequency")
	flags.PROD = c.Bool("production")
	flags.DEBUG = c.Bool("debug")
	flags.TEST = c.Bool("test")

	if flags.KubeConfigPath == "" || flags.RefreshFrequency == 0 {
		return errors.New("Error parsing flags")
	}
	return nil
}
