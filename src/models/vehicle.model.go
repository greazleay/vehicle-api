package models

import "github.com/google/uuid"

type Engine struct {
	Cylinder   int
	HorsePower float32
	Torque     int
}

type Vehicle struct {
	Base
	MakeID       uuid.UUID `gorm:"index"`
	Model        string
	Category     string
	Year         int
	NumberOfSeat int
	Price        int
	Engine       Engine `gorm:"embedded"`
}

type Make struct {
	Base
	Name     string
	Vehicles []Vehicle
}
