package util

type ErrorResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

func (errorResponse *ErrorResponse) SetStatus(status string) *ErrorResponse {
	errorResponse.Status = status
	return errorResponse
}

func (errorResponse *ErrorResponse) SetMessage(message string) *ErrorResponse {
	errorResponse.Message = message
	return errorResponse
}

func (errorResponse *ErrorResponse) SetPath(path string) *ErrorResponse {
	errorResponse.Path = path
	return errorResponse
}

func (errorResponse *ErrorResponse) SetTimestamp(timestamp string) *ErrorResponse {
	errorResponse.Timestamp = timestamp
	return errorResponse
}
