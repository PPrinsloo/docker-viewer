package dockerops

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
)

type DockerClient struct{}

type ContainerStats struct {
	CPUStats struct {
		CPUUsage struct {
			TotalUsage uint64 `json:"total_usage"`
		} `json:"cpu_usage"`
	} `json:"cpu_stats"`
	MemoryStats struct {
		Usage uint64 `json:"usage"`
		Limit uint64 `json:"limit"`
	} `json:"memory_stats"`
}

func (d *DockerClient) GetRunningContainers() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (d *DockerClient) FetchContainerStats(containerID string) (ContainerStats, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return ContainerStats{}, err
	}
	stats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return ContainerStats{}, err
	}
	defer stats.Body.Close()

	var containerStats ContainerStats
	if err := json.NewDecoder(stats.Body).Decode(&containerStats); err != nil {
		if err == io.EOF {
			return ContainerStats{}, nil // No stats, but not an error
		}
		return ContainerStats{}, err
	}

	return containerStats, nil
}
