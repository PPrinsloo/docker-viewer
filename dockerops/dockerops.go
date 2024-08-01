package dockerops

import (
	"github.com/docker/docker/api/types"
)

type DockerOps interface {
	GetRunningContainers() ([]types.Container, error)
	FetchContainerStats(containerID string) (ContainerStats, error)
}
