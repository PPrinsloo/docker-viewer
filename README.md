# Docker list containers

This is a simple terminal program to list all containers in a docker host and a basic status in a table.
It's mainly meant to try out the docker client for a different project I want to work on. I also ended up using this project to learn how to make github actions build binaries when I make a realease and to do some basic checks from a pull request.

## Usage
Install go from [here](https://golang.org/doc/install) and run the following command in the root of the project.
```bash
go run main.go
```

## Install using go
If you want the command to be generally available, you can install it using go from the root of the project.
```bash
go install
```

## Note
This is a simple program and is not meant to be used in production. It's just a simple program to list all containers in a docker host and a basic status in a table.
