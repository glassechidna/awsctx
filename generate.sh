#!/bin/sh
set -eux

cd internal/generate
go build
cd -

go get github.com/aws/aws-sdk-go
internal/generate/generate service
