package dtos

import "github.com/google/uuid"

type CreateVehicleDto struct {
	Model         string    `json:"model" binding:"required"`
	MakeID        uuid.UUID `json:"makeId" binding:"required"`
	Category      string    `json:"category" binding:"required"`
	Year          int       `json:"year" binding:"required"`
	ImageUrl      string    `json:"imageUrl" binding:"required"`
	NumberOfSeats int       `json:"numberOfSeats" binding:"required"`
	Price         int       `json:"price" binding:"required"`
	Cylinder      int       `json:"cylinder" binding:"required"`
	HorsePower    string    `json:"horsePower" binding:"required"`
	Torque        string    `json:"torque" binding:"required"`
	TopSpeed      string    `json:"topSpeed" binding:"required"`
	Acceleration  string    `json:"acceleration" binding:"required"`
	Transmission  int       `json:"transmission" binding:"required"`
}
