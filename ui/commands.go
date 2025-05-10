package ui

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	auth "github.com/oalabsi4/goitch/Auth"
	"github.com/oalabsi4/goitch/initialize"
	"github.com/oalabsi4/goitch/utils"
)

func GetChannels()  ([]list.Item, error) {
	// change the loading string
	channels,err:=utils.GetTwitchFollows(initialize.User.ID)
	if err != nil {
		return []list.Item{}, err
	}

	var items []item

	for _, channel := range channels.Data.Streams {
		descrption := fmt.Sprintf("%s | %s | %d viewers", utils.TruncateWithEllipsis(channel.Title, 35), channel.GameName, channel.ViewerCount)
		name := channel.UserLogin
		descrption = removeEmojis(descrption)
		items = append(items, item{Name: utils.ToCapsLock(name), description: utils.ToSmallCaps(descrption)})
	}
	// make the descrition string in equal length 
	items = addWhitespaces(items)
	i := itemsToListItems(items)
	
	return i,nil
}


func playStream(selectedChannel string) tea.Cmd {
	return func() tea.Msg {
		err := utils.PlayTwitchChannel(selectedChannel)
		if err != nil {
			return errMsg{err}
		}
		return nil
	}
}


func CheckToken() error {
	fmt.Println("Checking Token...")
	check, err := utils.TokenCheck()
	if err != nil {
		return err
	}
	if !check.Exists {
		log.Println("No user found, starting the OAuth flow...")
		err := auth.StartOAuthFlow()
		if err != nil {
			return err
		}
	}
	if check.Expired && check.Exists {
		log.Println("Token expired, starting the OAuth flow...")

		err := auth.UpdateToken()
		if err != nil {
			return err
		}
	}
	if check.Exists && !check.Expired && !check.Valid {
		log.Println("Token not valid, starting the OAuth flow...")

		err := auth.UpdateToken()
		if err != nil {
			return err
		}

	}
	return nil
}
func getChannelsCommand() tea.Cmd {
	return func() tea.Msg {
		channels, err := GetChannels()
		if err != nil {
			return errMsg{err}
		}
		return channels
	}
}

// a function to make the description equal in length
// we are using utf8.RuneCountInString to get the number of runes in the string since the description uses asci characters to make the text smaller
func addWhitespaces(strs []item) []item {
    maxLen := 0
    for _, str := range strs {
        runeCount := utf8.RuneCountInString(str.description)
        if runeCount > maxLen {
            maxLen = runeCount
        }
    }
    for i, str := range strs {
        runeCount := utf8.RuneCountInString(str.description)
        strs[i].description = str.description + strings.Repeat(" ", maxLen-runeCount)
    }
    return strs
}
// a function to remove emojis
func removeEmojis(input string) string {
	// Regular expression to match common emoji Unicode ranges
	re := regexp.MustCompile(`[\x{1F300}-\x{1F5FF}\x{1F600}-\x{1F64F}\x{1F680}-\x{1F6FF}\x{1F900}-\x{1F9FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}\x{1F1E6}-\x{1F1FF}\x{1F191}-\x{1F251}\x{1F004}\x{1F0CF}\x{1F170}-\x{1F171}\x{1F17E}-\x{1F17F}\x{1F18E}\x{3030}\x{2B50}\x{2B55}\x{2934}-\x{2935}\x{2B05}-\x{2B07}\x{2B1B}-\x{2B1C}\x{2B5E}\x{303D}\x{00A9}\x{00AE}\x{2122}\x{23F0}\x{23F3}\x{24C2}\x{25AA}-\x{25AB}\x{25B6}\x{25C0}\x{25FB}-\x{25FE}]`)
	// Replace emojis with an empty string
	return re.ReplaceAllString(input, "")
}
