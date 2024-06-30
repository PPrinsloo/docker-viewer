package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	containers := getRunningContainers()

	columns := dockerContainerListHeaders()

	rows := getContainerTableRows(containers)

	t := SetTableState(columns, rows)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
