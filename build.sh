#!/usr/bin/env bash

_BIN_NAME=esm

rm ./bin/*
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/${_BIN_NAME}_linux_amd64
upx ./bin/${_BIN_NAME}_linux_amd64
GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/${_BIN_NAME}_windows_amd64.exe
upx ./bin/${_BIN_NAME}_windows_amd64.exe
GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/${_BIN_NAME}_darwin_amd64
upx ./bin/${_BIN_NAME}_darwin_amd64
