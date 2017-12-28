package main

import (
	"fmt"
	"os"

	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
	"github.com/nii236/k/pkg/logger"
	"github.com/nii236/k/pkg/ui"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var log *logger.Log

func run(c *cli.Context) error {
	flags := &k.ParsedFlags{}
	err := flags.Parse(c)
	if err != nil {
		return err
	}

	logger.New(flags.LogToFile, flags.Debug)
	log = logger.Get()

	if flags.Test {
		mockClientSet, err := k8s.NewMock(flags)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "main.go: Could not initialise k8s client"))
		}
		app, err := ui.New(flags, mockClientSet)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "main.go: Could not initialise k8s client"))
		}
		return app.Run()
	}

	clientSet, err := k8s.New(flags)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "main.go: Could not initialise k8s client"))
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
			Name:   "refresh-interval",
			Usage:  "Seconds between updates",
			EnvVar: "REFRESH_INTERVAL",
			Value:  1,
		},
		cli.BoolFlag{
			Name:   "auto-refresh, a",
			Usage:  "Automatic refresh",
			EnvVar: "AUTO_REFRESH",
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
		cli.BoolFlag{
			Name:   "log-to-file, f",
			Usage:  "Log to file",
			EnvVar: "LOG_TO_FILE",
		},
		cli.StringFlag{
			Name:   "log-file-path",
			Usage:  "File to log to",
			Value:  "/tmp/debug.log",
			EnvVar: "LOG_FILE_PATH",
		},
		cli.BoolFlag{
			Name:   "test, t",
			Usage:  "Use the K8S mock client",
			EnvVar: "TEST",
		},
	}
	app.Run(os.Args)
}
