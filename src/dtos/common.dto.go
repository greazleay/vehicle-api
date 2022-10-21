package dtos

// uri is used for binding request params
type EntityID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type SuccessResponseDto struct {
	StatusText string      `json:"statusText"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type FailedResponseDto struct {
	StatusText string
	StatusCode int
	ErrorType  string
	Error      string
}
