package utils

import (
	"github.com/nicklaw5/helix"
	"github.com/oalabsi4/goitch/initialize"
)

// GetTwitchFollows retrieves the list of followed streams for a given user ID
// return only the first 100 followed streams that are live
func GetTwitchFollows(userId string) (*helix.StreamsResponse, error) {
	
	client := initialize.HelixClient

	// Get the user's followed streams
	channels, err := client.GetFollowedStream(&helix.FollowedStreamsParams{
		UserID: userId,
		First: 100,
	})
	if err != nil {
		
		return &helix.StreamsResponse{}, err
	}
	return channels, nil
}
