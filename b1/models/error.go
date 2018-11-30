package models

type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func GetGenericStatusResponse(StatusCode string, Message string) StatusResponse {

	return StatusResponse{
		Code:    StatusCode,
		Message: Message,
	}
}
