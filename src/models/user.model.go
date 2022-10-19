package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email     string `gorm:"not null;unique"`
	FirstName string
	LastName  string
	LastLogin time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Password  string    `gorm:"not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

	passwordHash, hashError := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashError != nil {
		err = errors.New(err.Error())
	}

	user.Password = string(passwordHash)

	return
}

func (user *User) ValidatePassword(providedPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
