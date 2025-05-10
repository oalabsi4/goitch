package initialize

import (
	"log"
	"os"

	"github.com/nicklaw5/helix"
)


var HelixClient *helix.Client
var User *helix.User
func InitTwitch(token string) {

	// Initialize the Helix client globally
	var err error
	HelixClient, err = helix.NewClient(&helix.Options{
        ClientID:     os.Getenv("TWCLIENT"),
        ClientSecret: os.Getenv("TWSECRET"),
		UserAccessToken:token ,
    })
    if err != nil {
        log.Fatal("Error creating helix client")
        return 
    }


	usersResp, err := HelixClient.GetUsers(&helix.UsersParams{})
	if err != nil || len(usersResp.Data.Users) == 0 {
		log.Fatalf("Failed to get user: %v", err)
	}

	user := usersResp.Data.Users[0]
	User = &user
	 
}