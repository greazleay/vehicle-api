package dtos

type CreateMake struct {
	Name string `json:"name" binding:"required"`
}
