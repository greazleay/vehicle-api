package dtos

type CreateMakeDto struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
}

// Form is used instead of json for binding request query
type MakeNameDto struct {
	Name string `form:"name" binding:"required"`
}

type MakeCountryDto struct {
	Country string `form:"country" binding:"required"`
}
