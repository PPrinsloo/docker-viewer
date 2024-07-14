package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	containers := getRunningContainers()

	columns := dockerContainerListHeaders()
	rows := getContainerTableRows(containers)

	containersTable := SetTableState(columns, rows)
	containerStatsTable := table.New() // Initialize with default settings for now

	m := model{
		containersTable:     containersTable,
		containerStatsTable: containerStatsTable,
	}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
