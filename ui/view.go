package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	loading := fmt.Sprintf("\n\n %s %s", m.LoadingString, m.Spinner.View())
	
	if m.Err != nil && !strings.Contains(m.Err.Error(), "failed to run yt-dlp") {
		err := fmt.Sprintf("Error woooo: %v", m.Err)
		return err
	}

	if m.State == ModelStateError {
		if strings.Contains(m.Err.Error(), "failed to run yt-dlp") {
			return fmt.Sprintf("Channel %s is offline\n\npress enter to continue", m.SelectedChannel)
		}
	}
	if m.State == ModelStateLoading {
		return loading
	}
	//? main screen
	if m.State == ModelStateMain {
		keyStyle := helpStyle.Render
		mainScreenHelp := fmt.Sprintf("Enter %s - esc %s - ctrl+p %s", keyStyle("play stream"), keyStyle("quit"), keyStyle("play none followed"))
		m.List.SetShowHelp(false)
		return fmt.Sprintf("%s\n%s", m.List.View(), mainScreenHelp)
	}
	//? play stream
	if m.State == ModelStatePlaying {
		keyStyle := helpStyle.Render
		mainScreenHelp := fmt.Sprintf("Enter %s - esc %s", keyStyle("send message"), keyStyle("return"))
		m.TextInput.Placeholder = "Send a message to chat"
		m.TextInput.Focus()
		// m.TextInput.CharLimit = 200
		ti:= itemStyle.Render(m.TextInput.View())
		
		vp := itemStyle.Render(m.Viewport.View())
		return fmt.Sprintf("%s\n%s\n%s", vp, ti,mainScreenHelp)
	}
	// ?play none followed channel 
	if m.State == ModelStatePlayNoneFollowed{
		keyStyle := helpStyle.Render
		mainScreenHelp := fmt.Sprintf("Enter %s - esc %s", keyStyle("play stream"), keyStyle("return"))

		ti := m.TextInput
		ti.Focus()
		ti.Placeholder = "Enter a channel name"
		return fmt.Sprintf("%s\n%s", inputStyle.Render(ti.View()), mainScreenHelp)
	}

	return "out of state"
}
