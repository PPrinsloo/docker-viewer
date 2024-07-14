package main

import (
	"context"
	"encoding/json"
	"github.com/charmbracelet/bubbles/table"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"strconv"
)

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

func getRunningContainers() []types.Container {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}

func dockerContainerListHeaders() []table.Column {
	columns := []table.Column{
		{Title: "ID", Width: 64},
		{Title: "Image", Width: 10},
		{Title: "Status", Width: 16},
		{Title: "State", Width: 10},
	}
	return columns
}

func getContainerTableRows(containers []types.Container) []table.Row {
	var rows []table.Row

	for _, c := range containers {
		row := table.Row{
			c.ID,
			c.Image,
			c.Status,
			c.State,
		}
		rows = append(rows, row)
	}
	return rows
}

func fetchContainerStats(containerID string) (ContainerStats, error) {
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

func convertStatsToRows(stats ContainerStats) []table.Row {
	rows := []table.Row{
		{"CPU Usage", formatCPUUsage(stats.CPUStats.CPUUsage.TotalUsage)},
		{"Memory Usage", formatMemoryUsage(stats.MemoryStats.Usage, stats.MemoryStats.Limit)},
		// Add more stats as needed
	}
	return rows
}

// Helper functions to format stats data
func formatCPUUsage(usage uint64) string {
	// Assuming usage is in nanoseconds, convert to seconds
	return strconv.FormatFloat(float64(usage)/1e9, 'f', 2, 64) + "s"
}

func formatMemoryUsage(usage uint64, limit uint64) string {
	// Convert bytes to MB
	usageMB := float64(usage) / 1024 / 1024
	limitMB := float64(limit) / 1024 / 1024
	return strconv.FormatFloat(usageMB, 'f', 2, 64) + "MB / " + strconv.FormatFloat(limitMB, 'f', 2, 64) + "MB"
}
