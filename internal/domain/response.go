package domain

type PaginatedResponse struct {
	Count   int         `json:"count"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
	Results interface{} `json:"results"`
}

type ErrorResponse struct {
	Detail   string `json:"detail"`
	Codename string `json:"codename"`
}

func NewPaginationErrorResponse(detail string) *ErrorResponse {
	return &ErrorResponse{
		Detail:   detail,
		Codename: "pagination_error",
	}
}

var (
	InternalServerError = ErrorResponse{
		Detail:   "Internal Server Error",
		Codename: "internal_server_error",
	}
)
