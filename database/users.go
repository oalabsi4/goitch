package database

import (
	"time"

	"github.com/oalabsi4/goitch/initialize"
	"github.com/oalabsi4/goitch/models"
)

func CreateUser(user models.User) (models.User, error) {
	err := initialize.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func GetUsers() ([]models.User, error) {
	var users []models.User
	err := initialize.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
// func UpdateUser(user models.User) (models.User, error) {
// 	initialize.DB.First(&user, "token = ?", user.Token)

// 	err := initialize.DB.Save(&user).Error
// 	if err != nil {
// 		return models.User{}, err
// 	}
// 	return user, nil
// }
func UpdateUser(user models.User,oldToken string)  error {
    // Attempt to find an existing user by Token
        resp := initialize.DB.Model(&user).Where("token = ?", oldToken).Updates(models.User{
            Token:        user.Token,
            RefreshToken: user.RefreshToken,
            TimeToLive:   user.TimeToLive,
            UpdatedAt:    time.Now(),
        })
    if resp.Error != nil {
        return  resp.Error
    }

    return  nil
}