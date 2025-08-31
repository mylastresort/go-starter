package seeds

import (
	"fmt"
	"log"
	"server/internal/models"
	"server/internal/services"

	"gorm.io/gorm/clause"
)

func AddUsersSeeds() {
	db := services.PostgresDB()

	users := []*models.User{
		{
			Name:  "first",
			Email: "example@email.org",
		},
		{
			Name:  "second",
			Email: "second@email.org",
		},
		{
			Name:  "third",
			Email: "third@email.org",
		},
	}

	res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(users)

	if res.Error != nil {
		log.Fatal("Could not seed users")
	}

	msg := fmt.Sprintf("%d users inserted", res.RowsAffected)

	services.Logger.Info(msg)
}
