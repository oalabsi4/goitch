package utils

import (
	"fmt"

	"net/http"
	"time"

	database "github.com/oalabsi4/goitch/database"
	"github.com/oalabsi4/goitch/models"
)

func TokenCheck() (models.Check, error) {
	fmt.Println("Checking Token...")
	checks := models.Check{
		Exists: true,
		Expired: false,
		Valid: true,
	}
	fmt.Println("after models")
	user, err := database.GetUsers()

	if len(user) == 0 || err != nil {
		fmt.Println("No user found")
		checks.Exists = false
		return checks,err
	}
	// check if the token is expired
	timeToLive := user[0].TimeToLive
	UpdatedAt := user[0].UpdatedAt
	checks.Expired = time.Since(UpdatedAt).Seconds() >= float64(timeToLive)


	// check if the token is valid
	checks.Valid = ValidateToken(user[0].Token)
	return checks,nil
}

func ValidateToken(token string) bool {

	req, err := http.NewRequest("GET", "https://id.twitch.tv/oauth2/validate", nil)
	if err != nil {
		return false
	}
	req.Header.Set("Authorization", "OAuth "+token)
	client := &http.Client{}

	// Make the HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	
	statusCode := res.StatusCode
	if statusCode != 200 {
		return false
	}

	defer res.Body.Close()

	return true	
}
