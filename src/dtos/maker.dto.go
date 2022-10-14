package dtos

type CreateMaker struct {
	Name string `json:"name" binding:"required"`
}
