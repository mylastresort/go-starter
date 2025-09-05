package users

import (
	"server/internal/models"
	"server/internal/services"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type CreateUserType struct {
	Name       string `validate:"required,min=4"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,min=8,containsany=!@#?*"`
	Provider   string `validate:"required"`
	ProviderId string
}

func CreateUser(newUser CreateUserType) (models.User, error) {
	db := services.PostgresDB()
	var user models.User
	user.Email = newUser.Email
	user.Name = newUser.Name
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 8)
	if err != nil {
		return user, err
	}
	user.Password = string(hashed)
	user.Provider = newUser.Provider
	user.ProviderId = newUser.ProviderId
	res := db.Clauses(clause.Returning{}).Create(&user)
	if res.Error != nil {
		return user, res.Error
	}
	db.First(&user, user.ID)
	return user, err
}
