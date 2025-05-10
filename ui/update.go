package ui

import (
	"fmt"
	"strings"

	//"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oalabsi4/goitch/utils"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Viewport.Width = msg.Width - 7
		m.Viewport.Height = msg.Height - 6
		m.TextInput.Width = msg.Width -10 
		m.List.SetWidth(msg.Width)
		m.List.DisableQuitKeybindings()
		m.List.SetHeight(msg.Height)
		itemStyle.Width(msg.Width)

		if len(m.Messages) > 0 {
			// Wrap content before setting it.
			m.Viewport.SetContent(lipgloss.NewStyle().Width(m.Viewport.Width).Render(strings.Join(m.Messages, "\n")))
		}
		// get the channels from the database
		return m, getChannelsCommand()

	case responseMsg:
		if m.State == ModelStatePlaying {
			m.Messages = append(m.Messages, msg.msg.text)
			m.Viewport.SetContent(lipgloss.NewStyle().Width(m.Viewport.Width).Render(strings.Join(m.Messages, "\n")))
			m.Viewport.GotoBottom()
			return m, waitForChatMessage(m.sub)
		}
		return m, nil

	case []list.Item:
		if m.State == ModelStateLoading {
			m.State = ModelStateMain
		}
		m.List.SetItems(msg)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "esc":
			if m.State == ModelStatePlaying ||m.State == ModelStatePlayNoneFollowed {
				m.State = ModelStateMain
				m.List.ResetSelected()
				m.SelectedChannel = ""

				m.TextInput.Reset()
				CloseChatConn = true
				utils.StopTwitchChannel()
				return m, nil
			}
			return m, tea.Quit
		// ? testing switching application state
		case "ctrl+l":
			m.State = ModelStateLoading
		case "enter":

			if m.State == ModelStateError {
				m.State = ModelStateMain
				m.Err = nil
				m.SelectedChannel = ""
				m.List.ResetSelected()
				return m, nil
			}

			if m.State == ModelStatePlaying {
				//todo: sent message to chat
				if m.TextInput.Value() != "" {
					m.ChatMessage = m.TextInput.Value()
					m.TextInput.Reset()
					return m, sendChatMessageCmd(m.ChatMessage, m.SelectedChannel)
				}

				return m, nil
			}
			if m.State == ModelStatePlayNoneFollowed {
				if m.TextInput.Value() != "" {
					existCheck, err :=utils.GetTwitchChannels([]string{m.TextInput.Value()})
					// check if channel exists or not 
					if len(existCheck.Data) == 0 || err != nil {
						m.State = ModelStateError
						m.Err = fmt.Errorf("channel %s does not exist", m.TextInput.Value())
						return m, nil
					}
					m.List.ResetSelected()
					m.SelectedChannel = m.TextInput.Value()
					m.TextInput.Reset()
					m.State = ModelStatePlaying
					return m, tea.Batch(
						waitForChatMessage(m.sub),
						ConnectToChannel(m.SelectedChannel, m.sub),
						playStream(m.SelectedChannel),
					)
				}
				return m, nil
			}
			i, ok := m.List.SelectedItem().(item)
			if ok {
				m.SelectedChannel = string(i.Name)
				m.State = ModelStatePlaying
				CloseChatConn = false
				m.Messages = []string{}
				m.Viewport.SetContent(strings.Join(m.Messages, "\n"))
				return m, tea.Batch(
					waitForChatMessage(m.sub),
					ConnectToChannel(m.SelectedChannel, m.sub),
					playStream(m.SelectedChannel),
				)
			}
			return m, nil

		case "ctrl+p":
		m.State = ModelStatePlayNoneFollowed
			return m, nil
		}

	case errMsg:
		m.Err = msg
		m.State = ModelStateError	
		return m, nil
	}
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)

	m.List, cmd = m.List.Update(msg)
	cmds = append(cmds, cmd)

	m.TextInput, cmd = m.TextInput.Update(msg)
	cmds = append(cmds, cmd)

	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
