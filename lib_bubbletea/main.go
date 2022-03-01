package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	todos := []string{"cleanning", "wash clothes", "write a blog"}
	model := NewModel(todos)
	cmd := tea.NewProgram(model)
	if err := cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		os.Exit(1)
	}
}
