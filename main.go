package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s\t%s\t%s\t%s\n", ctr.ID, ctr.Image, ctr.Status, ctr.State)
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Image", "Status", "State")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, ctr := range containers {
		fmt.Printf("%s\t%s\t%s\t%s\n", ctr.ID, ctr.Image, ctr.Status, ctr.State)
		tbl.AddRow(ctr.ID, ctr.Image, ctr.Status, ctr.State)
	}

	tbl.Print()
}
