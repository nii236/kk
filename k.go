package k

import (
	"errors"

	"github.com/urfave/cli"
)

type State struct {
	Resource    string
	Namespace   string
	CurrentView string
	Debug       string
}

// ParsedFlags will contain the config for the app
type ParsedFlags struct {
	KubeConfigPath   string
	RefreshFrequency int
	PROD             bool
	DEBUG            bool
}

// Parse will parse the flags into a struct
func (flags *ParsedFlags) Parse(c *cli.Context) error {
	flags.KubeConfigPath = c.String("kubeconfig-path")
	flags.RefreshFrequency = c.Int("refresh-frequency")
	flags.PROD = c.Bool("production")
	flags.DEBUG = c.Bool("debug")

	if flags.KubeConfigPath == "" || flags.RefreshFrequency == 0 {
		return errors.New("Error parsing flags")
	}
	return nil
}
