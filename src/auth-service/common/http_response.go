package common

type SuccesResponse struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewHttpSuccessResponse(status int, message string, data interface{}) SuccesResponse {
	return SuccesResponse{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}
}
