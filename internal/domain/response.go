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

func NewValidationErrorResponse(detail string) *ErrorResponse {
	return &ErrorResponse{
		Detail:   detail,
		Codename: "validation_error",
	}
}

func NewUniqueConstraintErrorResponse(detail string) *ErrorResponse {
	return &ErrorResponse{
		Detail:   detail,
		Codename: "unique_constraint_error",
	}
}

var (
	InternalServerErrorResponse = ErrorResponse{
		Detail:   "Internal Server Error",
		Codename: "internal_server_error",
	}
	NotFoundResponse = &ErrorResponse{
		Detail:   "Not found",
		Codename: "not_found",
	}
	ForbiddenResponse = &ErrorResponse{
		Detail:   "Forbidden",
		Codename: "forbidden",
	}
	InvalidURLParamErrorResponse = &ErrorResponse{
		Detail:   "Invalid url param",
		Codename: "invalid_url_param",
	}
	InvalidAuthCredentialsErrorResponse = &ErrorResponse{
		Detail:   "Invalid authentication credentials provided",
		Codename: "invalid_auth_credentials_error",
	}
)
