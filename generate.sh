#!/bin/sh
set -eux

cd internal/generate
go build
cd -

VERSION=${1:-""}
go get github.com/aws/aws-sdk-go
internal/generate/generate service "$VERSION"

go test ./...
