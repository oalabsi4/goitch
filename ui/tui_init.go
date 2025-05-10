package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	auth "github.com/oalabsi4/goitch/Auth"
	database "github.com/oalabsi4/goitch/database"
	"github.com/oalabsi4/goitch/initialize"
)

func InitialModel() Model {

	s := spinner.New()
	s.Spinner = spinner.Meter
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	err1 := initialize.InitDB()

	if err1 != nil {

		return Model{Spinner: s,
			Quitting:      true,
			State:         ModelStateLoading,
			Err:           err1,
			LoadingString: "failed to initialize DB",
		}
	}
	err := CheckToken()

	if err != nil {
		return Model{Spinner: s,
			Quitting:      true,
			State:         ModelStateLoading,
			Err:           err,
			LoadingString: "failed to check token",
		}
	}
	// getting the time to live form the user db
	user, errGetingUsers := database.GetUsers()
	if errGetingUsers != nil {
		return Model{Spinner: s,
			Quitting:      true,
			State:         ModelStateLoading,
			Err:           errGetingUsers,
			LoadingString: "failed to get users",
		}
	}
	
	timeToLive := user[0].TimeToLive
	timeUntilRefresh := timeToLive - 5*60 // 5 minutes before the token expires
	useToken := user[0].Token
	initialize.InitTwitch(useToken)
	// ticker := time.NewTicker(time.Duration(timeUntilRefresh) * time.Second) // creating a ticker to refresh the token
	ticker := time.NewTicker(time.Duration(timeUntilRefresh) * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				// ?? a bit iffy should keep an eye on this
				auth.UpdateToken()
			}
		}
	}()

	l := list.New( nil, itemDelegate{}, 20, 14)
	l.Title = "What do you want for dinner?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	

	vp := viewport.New(30, 20)

	ti := textinput.New()
	ti.Placeholder = "pokimane"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	sub := make(chan chatMessage)
	messages := []string{}
	return Model{
		Spinner:       s,
		Quitting:      false,
		State:         ModelStateLoading,
		LoadingString: "loading channels",
		List:          l,
		Viewport:      vp,
		TextInput:     ti,
		sub:           sub,
		Messages:      messages,
	}
}
