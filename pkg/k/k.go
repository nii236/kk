package k

import (
	"encoding/json"
	"errors"

	"github.com/urfave/cli"
)

// ModalKind is the type of Modal
type ModalKind string

// String returns the string representation of the ModalKind
func (k ModalKind) String() string {
	return string(k)
}

// TableKind is the type of Table
type TableKind string

// String returns the string representation of the TableKind
func (k TableKind) String() string {
	return string(k)
}

// Screen represents a main window
type Screen string

const (
	// ScreenTable is the Table screen
	ScreenTable Screen = "Table"
	// ScreenModal is the Modal screen
	ScreenModal Screen = "Modal"
	// ScreenState is the State screen
	ScreenState Screen = "State"
	// ScreenDebug is the Debug screen
	ScreenDebug Screen = "Debug"
)

func (s Screen) String() string {
	return string(s)
}

// State is the top level reducer
type State struct {
	UI       *UIReducer
	Entities *EntitiesReducer
}

// ParsedFlags will contain the config for the app
type ParsedFlags struct {
	KubeConfigPath  string
	RefreshInterval int
	AutoRefresh     bool
	LogFilePath     string
	LogToFile       bool
	Prod            bool
	Debug           bool
	Test            bool
}

// Parse will parse the flags into a struct
func (flags *ParsedFlags) Parse(c *cli.Context) error {
	flags.KubeConfigPath = c.String("kubeconfig-path")
	flags.RefreshInterval = c.Int("refresh-interval")
	flags.Prod = c.Bool("production")
	flags.Debug = c.Bool("debug")
	flags.Test = c.Bool("test")
	flags.LogToFile = c.Bool("log-to-file")
	flags.LogFilePath = c.String("log-file-path")
	flags.AutoRefresh = c.Bool("auto-refresh")

	if flags.KubeConfigPath == "" {
		return errors.New("Error parsing flags")
	}
	return nil
}

// JSONString returns a string representation of the top level reducer for debugging purposes
func (s *State) JSONString() (string, error) {
	b, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
