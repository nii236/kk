
# K
[![](https://godoc.org/github.com/nii236/k?status.svg)](http://godoc.org/github.com/nii236/k)
[![Go Report Card](https://goreportcard.com/badge/github.com/nii236/k)](https://goreportcard.com/report/github.com/nii236/k)

I got sick of typing the same `kubectl` commands over and over again so here is my simplified TUI wrapper, K.

# Installation

```
go get github.com/nii236/k
cd $GOPATH/src/github.com/nii236/k
dep ensure
go install cmd/k/main.go
k
```

```
NAME:
   k - Terminal User Interface (TUI) for Kubernetes

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   For when you are sick of typing namespaces over and over again

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --kubeconfig-path value, -c value  Kubeconfig path (Uses $HOME) (default: "/Users/nii236/.kube/admin.conf") [$KUBECONFIG_PATH]
   --refresh-interval value           Seconds between updates (default: 1) [$REFRESH_INTERVAL]
   --auto-refresh, -a                 Automatic refresh [$AUTO_REFRESH]
   --production, -p                   Production mode [$PRODUCTION]
   --debug, -d                        Debug logging [$DEBUG]
   --log-to-file, -f                  Log to file [$LOG_TO_FILE]
   --log-file-path value              File to log to (default: "/tmp/debug.log") [$LOG_FILE_PATH]
   --test, -t                         Use the K8S mock client [$TEST]
   --help, -h                         show help
   --version, -v                      print the version
```

# Screenshots

![](/static/screenshot.png)
