package dtos

type EntityID struct {
	ID string `uri:"id" binding:"required,uuid"`
}
