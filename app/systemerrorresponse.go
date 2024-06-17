package app

type ErrorResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

func (errorResponse *ErrorResponse) status(status string) *ErrorResponse {
	errorResponse.Status = status
	return errorResponse
}

func (errorResponse *ErrorResponse) message(message string) *ErrorResponse {
	errorResponse.Message = message
	return errorResponse
}

func (errorResponse *ErrorResponse) path(path string) *ErrorResponse {
	errorResponse.Path = path
	return errorResponse
}

func (errorResponse *ErrorResponse) timestamp(timestamp string) *ErrorResponse {
	errorResponse.Timestamp = timestamp
	return errorResponse
}
