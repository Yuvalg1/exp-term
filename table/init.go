package table

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	table table.Model
}

func InitialModel() Model {
	columns := []table.Column{
		{Title: "Name", Width: 16},
		{Title: "Type", Width: 10},
		{Title: "Size", Width: 10},
		{Title: "Date Modified", Width: 16},
	}

	rows := getDirectoryContent()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
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

	return Model{t}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func getDirectoryContent() []table.Row {
	currentwd, errwd := os.Getwd()
	if errwd != nil {
		log.Fatal(errwd)
	}

	entries, errread := os.ReadDir(currentwd)
	if errread != nil {
		log.Fatal(errread)
	}

	var names []table.Row

	for _, entry := range entries {
		names = append(names, GetEntryContent(entry, currentwd))
	}
	return names
}
