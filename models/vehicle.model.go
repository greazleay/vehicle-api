package models

import "github.com/google/uuid"

type Engine struct {
	Cylinder   int
	HorsePower float32
	Torque     int
}

type Vehicle struct {
	Base
	MakerID      uuid.UUID
	Model        string
	Category     string
	Year         int
	NumberOfSeat int
	Price        int
	Engine       Engine `gorm:"embedded"`
}

type Maker struct {
	Base
	Name     string
	Vehicles []Vehicle
}
