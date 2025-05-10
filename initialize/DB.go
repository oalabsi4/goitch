package initialize

import (
	"github.com/glebarez/sqlite"
	"github.com/oalabsi4/goitch/models"
	"gorm.io/gorm"
)




var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("goitch.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	m1err := DB.AutoMigrate(&models.User{}); if m1err != nil {
		return m1err
	}

	return nil
}
