package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Base
	Email     string `gorm:"not null;unique"`
	FirstName string
	LastName  string
	LastLogin time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	Password  string    `gorm:"not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) {

	user.Password = "rrrrrrrrr"

	return
}
