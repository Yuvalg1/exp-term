package main

import (
	"fmt"
	"os"

	ls "github.com/Yuvalg1/exp-term/ls"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ls.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
