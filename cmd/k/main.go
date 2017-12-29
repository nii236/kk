package main

import (
	"fmt"
	"os"

	"github.com/nii236/k"
	"github.com/nii236/k/pkg/k8s"
	"github.com/nii236/k/pkg/logger"
	"github.com/nii236/k/pkg/ui"
	"github.com/urfave/cli"
)

func run(c *cli.Context) error {
	flags := &k.ParsedFlags{}
	err := flags.Parse(c)
	if err != nil {
		return err
	}

	logger.New(flags.PROD, flags.DEBUG)
	clientSet, err := k8s.New(flags)
	if err != nil {
		fmt.Println(err)
	}
	app, err := ui.New(flags, clientSet)
	if err != nil {
		fmt.Println(err)
	}
	return app.Run()
}

func main() {

	app := cli.NewApp()
	app.Name = "k"
	app.Usage = "Terminal User Interface (TUI) for Kubernetes"
	app.Description = "For when you are sick of typing namespaces over and over again"
	app.Version = "0.0.1"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "kubeconfig-path, c",
			Usage:  "Kubeconfig path (Uses $HOME)",
			EnvVar: "KUBECONFIG_PATH",
			Value:  fmt.Sprintf("%s/.kube/admin.conf", os.Getenv("HOME")),
		},
		cli.IntFlag{
			Name:   "refresh-frequency, f",
			Usage:  "Seconds between updates",
			EnvVar: "REFRESH_FREQUENCY",
			Value:  5,
		},
		cli.BoolFlag{
			Name:   "production, p",
			Usage:  "Production mode",
			EnvVar: "PRODUCTION",
		},
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "Debug logging",
			EnvVar: "DEBUG",
		},
	}
	app.Run(os.Args)
}
