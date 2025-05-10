package ui

import (

	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Spinner         spinner.Model
	Viewport        viewport.Model
	Textarea        textarea.Model
	SenderStyle     lipgloss.Style
	Quitting        bool
	CurrentChannel  string
	Messages        []string
	State           ModelState
	List            list.Model
	SelectedChannel string
	TextInput       textinput.Model
	Err             error
	LoadingString   string

	sub           chan chatMessage
	ChatMessage   string
}

var CloseChatConn bool = false
type errMsg struct{ err error }

func (e errMsg) Error() string {
	return e.err.Error()
}

type TokenPass bool

type chatMessage struct {
	text string
}

type responseMsg struct {
	msg chatMessage
}

const ListHeight = 14

type ModelState int

const (
	ModelStateMain ModelState = iota
	ModelStatePlaying
	ModelStateLoading
	ModelStateError
	ModelStatePlayNoneFollowed
)

type item struct {
	Name        string
	description string
}

func (i item) FilterValue() string { return "" }
func (i item) Description() string { return i.description }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 4 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) ShowDescription() bool                   { return true }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	dsc := fmt.Sprintf("\n%s", descriptionStyle.Render(i.description))
	str := fmt.Sprintf("%d. %s", index+1, i.Name)

	//fn := itemStyle.Render
	//on := onlineStyle.Render
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render(s...)
		}
	}
	// fmt.Fprint(w, fn(str)+ dsc)
	fmt.Fprint(w, fn(str+dsc))

}



func itemsToListItems(items []item) []list.Item {
	listItems := []list.Item{}
	for _, i := range items {
		listItems = append(listItems, i)
	}
	return listItems
}
