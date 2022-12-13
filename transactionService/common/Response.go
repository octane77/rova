package common

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err string) Response {
	return Response{
		Status:  false,
		Message: message,
		Error:   err,
		Data:    nil,
	}
}
