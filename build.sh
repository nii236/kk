#!/bin/bash

gox -output="./dist/{{.Dir}}_{{.OS}}_{{.Arch}}" -os="linux" ./cmd/kk/