package dtos

// uri is used for binding request params
type EntityID struct {
	ID string `uri:"id" binding:"required,uuid"`
}
