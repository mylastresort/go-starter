package services

import (
	"server/internal/models"
)

func GetUsers() ([]*models.User, error) {
	var users []*models.User
	db := PostgresDB()

	res := db.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

type CreateUserType struct {
	Name     string `validate:"required,min=4"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,containsany=!@#?*"`
}

func CreateUser(newUser CreateUserType) error {
	db := PostgresDB()
	var user models.User
	user.Email = newUser.Email
	user.Name = newUser.Name
	res := db.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	db := PostgresDB()
	var user models.User
	res := db.First(&user, "email = ?", email)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func UpdateUser(user models.User) error {
	db := PostgresDB()
	res := db.Save(&user)
	return res.Error
}
