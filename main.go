package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
	"github.com/oalabsi4/goitch/ui"
)

func main() {
	//read text file from the ascii folder and print it, paths are relative to the main.go file
	asciiLogo := ` ░▒▓██████▓▒░  ░▒▓██████▓▒░ ░▒▓█▓▒░░▒▓████████▓▒░░▒▓██████▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░       ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░   ░▒▓█▓▒░   ░▒▓█▓▒░       ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒▒▓███▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░   ░▒▓█▓▒░   ░▒▓█▓▒░       ░▒▓████████▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░   ░▒▓█▓▒░   ░▒▓█▓▒░       ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ 
 ░▒▓██████▓▒░  ░▒▓██████▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒░    ░▒▓██████▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ 
                                                                             
                                                                             `

	logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#6441a5")).Align(lipgloss.Center, lipgloss.Center)
	fmt.Println(logoStyle.Render(asciiLogo))
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// Initialize the Twitch API client with the user token
	// Initialize the Bubble Tea program with the initial model
	if _, err := tea.NewProgram(ui.InitialModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
