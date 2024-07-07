#!/bin/bash

# Create the build directory if it doesn't exist
mkdir -p build

# Build for Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/container-view-linux-amd64

# Build for Linux ARM
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/container-view-linux-arm64

# Build for Mac Intel
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/container-view-darwin-amd64

# Build for Mac ARM (Apple Silicon)
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/container-view-darwin-arm64

# Build for Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/container-view-windows-amd64.exe

# Build for Windows ARM
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/container-view-windows-arm64.exe