package main

import (
	"fmt"
	"os"

	ls "github.com/Yuvalg1/exp-term/dir"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := ls.InitialModel()

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
