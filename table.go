package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	containerSelected   string
	containersTable     table.Model
	containerStatsTable table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.containerSelected = ""
			m.containerStatsTable.Focus()
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.containerSelected == "" {
				selectedRow := m.containersTable.SelectedRow()
				containerID := selectedRow[0] // Assuming the first column contains the container ID.
				m.containerSelected = containerID

				// Fetch stats for the selected container and update containerStatsTable.
				stats, err := fetchContainerStats(containerID)
				if err != nil {
					fmt.Println("Error fetching container stats:", err)
					return m, nil
				}
				statsRows := convertStatsToRows(stats)
				statsColumns := []table.Column{
					{Title: "Metric", Width: 20},
					{Title: "Value", Width: 30},
				}
				m.containerStatsTable = SetTableState(statsColumns, statsRows)
				m.containerStatsTable.Focus()
			}
		}
	}
	m.containersTable, cmd = m.containersTable.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.containerSelected != "" {
		return baseStyle.Render(m.containerStatsTable.View()) + "\n"
	}
	return baseStyle.Render(m.containersTable.View()) + "\n"
}

func SetTableState(columns []table.Column, rows []table.Row) table.Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}
