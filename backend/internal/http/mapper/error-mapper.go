package mapper

import "givemegoodcoffee/internal/http/response"

type ErrorMapper struct{}

func NewErrorMapper() *ErrorMapper {
	return &ErrorMapper{}
}

func (m ErrorMapper) ToErrorResponse(error string) *response.ErrorResponse {
	return &response.ErrorResponse{Message: error}
}
