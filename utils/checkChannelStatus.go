package utils

import (
	"encoding/json"
	database "github.com/oalabsi4/goitch/database"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/oalabsi4/goitch/models"
)


func GetTwitchChannels(channelNames []string) (models.ChannelResponse, error) {
	clientID := os.Getenv("TWCLIENT")
	user, err := database.GetUsers()
	if err != nil {
		return models.ChannelResponse{}, err
	}
	authToken := user[0].Token
	var b strings.Builder

	for i, name := range channelNames {
		if i > 0 {
			b.WriteString("&")
		}
		b.WriteString("user_login=")
		b.WriteString(name)
	}

	query := b.String()

	url := "https://api.twitch.tv/helix/streams?"

	req, err := http.NewRequest("GET", url+query, nil)
	if err != nil {
		return models.ChannelResponse{}, err
	}

	req.Header.Set("Client-ID", clientID)
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.ChannelResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	data := models.ChannelResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("failed to unmarshal response body: %v", err)
	}

	return data, nil
}


func Contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
