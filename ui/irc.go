package ui

import (
	"bufio"
	database "github.com/oalabsi4/goitch/database"
	"math/rand"
	"time"

	"fmt"
	"log"
	"net"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
var chatColors = []lipgloss.Color{
	lipgloss.Color("#fb0d03"), 
	lipgloss.Color("#0422fc"), 
	lipgloss.Color("##00801b"), 
	lipgloss.Color("#b52324"),
	lipgloss.Color("##ff7e56"),
	lipgloss.Color("#99cb46"),
	lipgloss.Color("#fc480d"),
	lipgloss.Color("#2d8a5b"),
	lipgloss.Color("#dca236"),
	lipgloss.Color("#5e9ea0"),
	lipgloss.Color("#1c93fb"),
	lipgloss.Color("#ff6cb2"),
	lipgloss.Color("#8c37da"),
	lipgloss.Color("#04fb89"),
}
var seed = time.Now().Unix()
var s = rand.NewSource(seed)
func pickRandomColor() lipgloss.Color {
	// Pick a random color from the list
	r := rand.New(s) // initialize local pseudorandom generator \
	colorCount := len(chatColors)
	randomIndex := r.Intn(colorCount)
	return chatColors[randomIndex]
	
}

var (
	ChatUserNameStyle = lipgloss.NewStyle().Bold(true)
	msgTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
)
func ReadMessages(channelName string, sub chan chatMessage) error {
	user, err := database.GetUsers()
	if err != nil {
		return err
	}
	// Twitch IRC server details
	server := "irc.chat.twitch.tv:6667"
	nickname := "goitchUser"                   // Replace with your Twitch username
	oauthToken := "oauth:" + user[0].Token     // Replace with your Twitch OAuth token
	channel := fmt.Sprintf("#%s", channelName) //"#zackrawrr" // Replace with the channel you want to join

	// Connect to the Twitch IRC server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatalf("Failed to connect to Twitch IRC: %v", err)
	}
	//defer conn.Close()

	// Authenticate with the server
	fmt.Fprintf(conn, "PASS %s\r\n", oauthToken)
	fmt.Fprintf(conn, "NICK %s\r\n", nickname)
	fmt.Fprintf(conn, "JOIN %s\r\n", channel)

	//reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(conn)
	//fmt.Println("Reading messages from Twitch IRC...")

	go func() {

		//check if the sub channel is closed
		defer conn.Close()
		
		for scanner.Scan() {
			if CloseChatConn  {
				return
			}
			// Read a line from the connection
			line := scanner.Text()

			// Handle PING messages
			if strings.HasPrefix(line, "PING") {
				// Extract the content of the PING message and reply with PONG
				pongMessage := strings.TrimSpace(line)
				fmt.Fprintf(conn, "%s\r\n", pongMessage)

				continue
			}

			// Parse and display chat messages
			if strings.Contains(line, "PRIVMSG") {
				parts := strings.Split(line, "!")
				username := parts[0][1:] // Extract username
				message := strings.Split(line, "PRIVMSG")[1]
				message = strings.SplitN(message, ":", 2)[1]
				userNameStyle := ChatUserNameStyle.Foreground(pickRandomColor()).Render(username)// Extract message content
			

				msg := chatMessage{text: fmt.Sprintf("%s: %s", userNameStyle, msgTextStyle.Render(message))}
				
				// fmt.Println("printing: ", msg.text) // Print the message to the console
				sub <- msg

				continue
			}

			// fmt.Printf("Unknown message: %s\n", line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from server:", err)
		}

	}()
	return nil
}
func sendChatMessage(message string, channelName string)error {
	user, err := database.GetUsers()
	if err != nil {
		return err
	}
	// Twitch IRC server details
	server := "irc.chat.twitch.tv:6667"
	nickname := "goitchUser"                   // Replace with your Twitch username
	oauthToken := "oauth:" + user[0].Token     // Replace with your Twitch OAuth token
	channel := fmt.Sprintf("#%s", channelName) //"#zackrawrr" // Replace with the channel you want to join

	// Connect to the Twitch IRC server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatalf("Failed to connect to Twitch IRC: %v", err)
	}
	defer conn.Close()

	// Authenticate with the server
	fmt.Fprintf(conn, "PASS %s\r\n", oauthToken)
	fmt.Fprintf(conn, "NICK %s\r\n", nickname)
	fmt.Fprintf(conn, "JOIN %s\r\n", channel)

	msgSting := fmt.Sprintf("PRIVMSG %s :%s\r\n", channel, message)
	fmt.Fprintf(conn,"%s", msgSting)

 return nil
}
func sendChatMessageCmd(message string, channelName string) tea.Cmd {
	return func() tea.Msg {
		err := sendChatMessage(message, channelName)
		if err != nil {
			return errMsg{err}
		}
		return nil
	}
}
func waitForChatMessage(sub chan chatMessage) tea.Cmd {
	return func() tea.Msg {
		msg := <-sub
		return responseMsg{msg: msg}

	}
}

func ConnectToChannel(channelName string, sub chan chatMessage) tea.Cmd {
	return func() tea.Msg {
		err := ReadMessages(channelName, sub)
		if err != nil {
			return errMsg{err}
		}
		return nil
	}
}

