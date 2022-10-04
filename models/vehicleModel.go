package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	Make         string
	Mod          string
	Year         int
	NumberOfSeat int
}
