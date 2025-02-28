package main

import (
	"fmt"
	"os"

	"github.com/Yuvalg1/exp-term/tab"
	"github.com/Yuvalg1/exp-term/table"
	teaTable "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	columns := []teaTable.Column{
		{Title: "Name", Width: 16},
		{Title: "Type", Width: 10},
		{Title: "Size", Width: 10},
		{Title: "Date Modified", Width: 16},
	}
	tabContent := table.InitialModel(columns)
	m := tab.Model{Tabs: columns, TabContent: tabContent}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
