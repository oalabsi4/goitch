package ui

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// an empty function to satisfy the interface

var (
	titleStyle = lipgloss.NewStyle().MarginLeft(2)
	//itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	//selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle  = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	descriptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#707070")).PaddingLeft(2).Faint(true).Italic(true)

	itemStyle         = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#707070")).Padding(0, 1).Align(lipgloss.Left)
	selectedItemStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#6441a5")).Padding(0, 1).Align(lipgloss.Left).Foreground(lipgloss.Color("#6441a5"))
	helpStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#707070")).Faint(true)
	inputStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#6441a5")).Padding(0, 1).Align(lipgloss.Center, lipgloss.Center)
)

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return width - 10
}
