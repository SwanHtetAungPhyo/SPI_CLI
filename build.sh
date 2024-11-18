#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o spi_window.exe
GOOS=linux GOARCH=amd64 go build -o spi_linux
GOOS=darwin GOARCH=amd64 go build -o spi_darwin


