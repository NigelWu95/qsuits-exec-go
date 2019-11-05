#!/usr/bin/env bash
GOOS="darwin" GOARCH="amd64" go build -o "bin/qsuits_darwin_amd64" src/entry.go
#GOOS="darwin" GOARCH="386" go build -o "bin/qsuits_darwin_386" src/entry.go
GOOS="windows" GOARCH="amd64" go build -o "bin/qsuits_windows_amd64.exe" src/entry.go
GOOS="windows" GOARCH="386" go build -o "bin/qsuits_windows_386.exe" src/entry.go
GOOS="linux" GOARCH="amd64" go build -o "bin/qsuits_linux_amd64" src/entry.go
GOOS="linux" GOARCH="386" go build -o "bin/qsuits_linux_386" src/entry.go
chmod +x bin/*
qsuits -path=bin -process=qupload -a=tswork -bucket=qsuits -keep-path=false -save-path=logs