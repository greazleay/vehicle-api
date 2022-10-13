package models

import "gorm.io/gorm"

type Engine struct {
	Cylinder   int
	HorsePower float32
	Torque     int
}

type Vehicle struct {
	gorm.Model

	Make         string
	VModel       string
	Category     string
	Year         int
	NumberOfSeat int
	Price        int
	Engine       Engine `gorm:"embedded"`
}

type Maanufacturer struct {
	gorm.Model

	Name string
}
