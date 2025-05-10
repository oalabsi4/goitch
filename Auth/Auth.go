package auth

import (
	"fmt"
	//"go/token"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/nicklaw5/helix"
	"github.com/oalabsi4/goitch/database"
	"github.com/oalabsi4/goitch/models"
	"github.com/oalabsi4/goitch/shared"
	"github.com/skratchdot/open-golang/open"
)

func StartOAuthFlow() error{
    clientID := os.Getenv("TWCLIENT")
    clientSecret := os.Getenv("TWSECRET")
    redirectURI := "http://localhost:8080/oauth/callback"
    scopes := []string{"chat:read", "chat:edit","user:read:broadcast","user:read:follows"}
    // scopes := []string{"user:read:chat","user:write:chat","chat:read","chat:write"}

    // Initialize the Helix client globally
    client, err := helix.NewClient(&helix.Options{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURI:  redirectURI,
    })
    if err != nil {
        return err
    }
   
    shared.HelixClient = client

    // Generate the authorization URL
    authURL := shared.HelixClient.GetAuthorizationURL(&helix.AuthorizationURLParams{
        ResponseType: "code",
        Scopes:       scopes,
        State:        "some-state",
    })

    // Open the browser for authorization
    fmt.Println("Opening browser for authorization...")
    open.Run(authURL)

    // Start the HTTP server to handle the callback
    startServer()

    return nil
}

func startServer() {
    http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")
        if code == "" {
            http.Error(w, "Code not found in callback", http.StatusBadRequest)
            return
        }

        // Exchange the code for a token
        tokenResp, err := shared.HelixClient.RequestUserAccessToken(code)
        if err != nil {
            log.Printf("Error exchanging code for token: %v", err)
            http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
            return
        }

        // Save the token globally
        shared.AccessToken = tokenResp.Data.AccessToken

        database.CreateUser(models.User{Token: shared.AccessToken, RefreshToken: tokenResp.Data.RefreshToken, CreatedAt: time.Now(), UpdatedAt: time.Now(), TimeToLive: tokenResp.Data.ExpiresIn})
        w.Write([]byte("Authorization successful! You can close this window."))
        style := lipgloss.NewStyle().Foreground(lipgloss.Color("#6441a5")).Render
        fmt.Println(style("Authorization successful! Restart the App."))
    })

    log.Println("Starting server at :8080...")
    go log.Fatal(http.ListenAndServe(":8080", nil))
}

// UpdateToken refreshes the access token for the user by using the refresh token.
//
// Updates the user in the database with the new access token, refresh token, and time to live.
func UpdateToken() error {
    // Get the user from the database
    user, err := database.GetUsers()
    if err != nil {
        return err
    }

    // Create a new Helix client with the client ID and secret
    client, err := helix.NewClient(&helix.Options{
        ClientID:     os.Getenv("TWCLIENT"),
        ClientSecret: os.Getenv("TWSECRET"),
    })
    if err != nil {
        fmt.Println("Error creating helix client")
        return err
    }

    // Get the refresh token and access token from the user
    refreshToken := user[0].RefreshToken
    token := user[0].Token

    // Refresh the access token using the refresh token
    resp, err := client.RefreshUserAccessToken(refreshToken)
    if err != nil {
        fmt.Println("Error refreshing the access token")
        return err
    }

    // Update the user in the database with the new access token, refresh token, and time to live
    database.UpdateUser(models.User{Token: resp.Data.AccessToken, RefreshToken: resp.Data.RefreshToken, TimeToLive: resp.Data.ExpiresIn}, token)

    return nil
}
