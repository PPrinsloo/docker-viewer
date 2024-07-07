#!/bin/bash

echo "Building container-view for all platforms"

echo "Creating build directory..."
mkdir -p build

if [ ! -d "build" ]; then
  echo "Failed to create build directory"
  exit 1
fi
echo "Build directory created"

echo "Building for Linux AMD64..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/container-view-linux-amd64
# Check if the build was successful
if [ ! -f "build/container-view-linux-amd64" ]; then
  echo "Failed to build for Linux AMD64"
  exit 1
fi
echo "Build for Linux AMD64 done"


echo "Building for Linux ARM..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/container-view-linux-arm64
# Check if the build was successful
if [ ! -f "build/container-view-linux-arm64" ]; then
  echo "Failed to build for Linux ARM"
  exit 1
fi
echo "Build for Linux ARM done"


echo "Building for Mac Intel..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/container-view-darwin-amd64
# Check if the build was successful
if [ ! -f "build/container-view-darwin-amd64" ]; then
  echo "Failed to build for Mac Intel"
  exit 1
fi
echo "Build for Mac Intel done"


echo "Building for Mac ARM..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/container-view-darwin-arm64
# Check if the build was successful
if [ ! -f "build/container-view-darwin-arm64" ]; then
  echo "Failed to build for Mac ARM"
  exit 1
fi
echo "Build for Mac ARM done"


echo "Building for Windows AMD64..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/container-view-windows-amd64.exe
# Check if the build was successful
if [ ! -f "build/container-view-windows-amd64.exe" ]; then
  echo "Failed to build for Windows AMD64"
  exit 1
fi
echo "Build for Windows AMD64 done"

# Build for Windows ARM
echo "Building for Windows ARM..."
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/container-view-windows-arm64.exe
# Check if the build was successful
if [ ! -f "build/container-view-windows-arm64.exe" ]; then
  echo "Failed to build for Windows ARM"
  exit 1
fi
echo "Build for Windows ARM done\n All builds successful!"