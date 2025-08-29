package models

import (
	"github.com/lib/pq"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Tokens pq.StringArray `gorm:"type:text[]"`
	Name   string
	Email  string `gorm:"uniqueIndex"`
}
