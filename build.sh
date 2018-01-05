#!/bin/bash

gox -arch="amd64" -os="darwin linux windows" -output="./dist/{{.Dir}}_{{.OS}}_{{.Arch}}" ./cmd/kk/