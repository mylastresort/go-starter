package services

import (
	"server/internal/models"
)

func GetUsers() ([]*models.User, error) {
	var users []*models.User
	db := DB()

	if res := db.Find(&users); res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}
