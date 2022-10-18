package models

import "github.com/google/uuid"

/*
For Has Many Relations, define a reverse Belongs to relation (Optional)
Horsepower determines how fast a Vehicle will run, it provides a measurement of the engine’s overall performance.
Torque determines the performace of a Vehicle's engine when streched, it provides a measurement of the maximum twisting force that engine can generate, when worked hard.
*/

type Vehicle struct {
	Base
	Model         string
	MakeID        uuid.UUID `gorm:"index"`
	Make          Make
	Category      string
	Year          int
	ImageUrl      string
	NumberOfSeats int
	Price         int
	Cylinder      int
	HorsePower    string
	Torque        string
	TopSpeed      string
	Acceleration  string
	Transmission  int
}

type Make struct {
	Base
	Name     string
	Country  string
	Vehicles []Vehicle
}
