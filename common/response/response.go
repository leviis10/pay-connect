package response

type SuccessResponse struct {
	Status  string
	Message string
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string
	Error   string
	Message string `json:"message,omitempty"`
}
