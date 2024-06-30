package main

import (
	"context"
	"github.com/charmbracelet/bubbles/table"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

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
