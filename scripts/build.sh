#!/bin/bash

RUN_NAME="YOUR_BINARY_NAME"
PROJECT_NAME="YOUR_PROJECT_NAME"

go env -w GO111MODULE="on"
go env -w GOPROXY="YOUR_GOPROXY_URL"
go env -w GOPRIVATE=""
#debug 信息
go version
go env

cd ${PROJECT_NAME} || exit

rm -rf output
mkdir -p output/bin output/conf

cp -r configs/* output/conf

cp scripts/bootstrap.sh output/
chmod +x output/bootstrap.sh

go build -o output/bin/${RUN_NAME} cmd/main.go



