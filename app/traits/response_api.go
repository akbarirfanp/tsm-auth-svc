package traits

import "github.com/goravel/framework/contracts/http"

type ResponseAPI struct{}

type SuccessResponse struct {
	Error   bool        `json:"error"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error   bool        `json:"error"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type PaginatedData struct {
	Items      interface{} `json:"items"`
	Pagination interface{} `json:"pagination"`
}

type PaginatedResponse struct {
	Error   bool          `json:"error"`
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    PaginatedData `json:"data"`
}

func (r ResponseAPI) Success(ctx http.Context, data any, message string) http.Response {
	return ctx.Response().Json(200, SuccessResponse{
		Error:   false,
		Status:  200,
		Message: message,
		Data:    data,
	})
}

func (r ResponseAPI) Error(ctx http.Context, code int, message string, errors any) http.Response {
	return ctx.Response().Json(code, ErrorResponse{
		Error:   true,
		Status:  code,
		Message: message,
		Errors:  errors,
	})
}

func (r ResponseAPI) PaginatedResponse(ctx http.Context, items any, total, page, pageSize int, message string) http.Response {
	totalPages := (total + pageSize - 1) / pageSize

	return ctx.Response().Json(200, PaginatedResponse{
		Error:   false,
		Status:  200,
		Message: message,
		Data: PaginatedData{
			Items: items,
			Pagination: map[string]any{
				"total":       total,
				"page":        page,
				"page_size":   pageSize,
				"total_pages": totalPages,
			},
		},
	})
}
