package utils

import "net/http"

// ErrorResponse представляет формат ошибки в JSON
//
// swagger:model
type ErrorResponse struct {
	Message string `json:"message" example:"Error message"`
}

// newError создает новый объект ошибки
func newError(msg string) *ErrorResponse {
	return &ErrorResponse{
		Message: msg,
	}
}

// BadRequestError возвращает 400 Bad Request
func BadRequestError() (int, *ErrorResponse) {
	return http.StatusBadRequest, newError("Bad Request")
}

// NotFoundError возвращает 404 Not Found
func NotFoundError() (int, *ErrorResponse) {
	return http.StatusNotFound, newError("Not Found")
}

// InternalServerError возвращает 500 Internal Server Error
func InternalServerError(msg string) (int, *ErrorResponse) {
	return http.StatusInternalServerError, newError(msg)
}

// ConflictError возвращает 409 Conflict
func ConflictError() (int, *ErrorResponse) {
	return http.StatusConflict, newError("Some fields can't be modified")
}

// MultipleLoginError возвращает ошибку уникальности логина
func MultipleLoginError() *ErrorResponse {
	return &ErrorResponse{
		Message: "Login should be unique",
	}
}

// UnauthorizedError возвращает 401 Unauthorized
func UnauthorizedError() (int, *ErrorResponse) {
	return http.StatusUnauthorized, newError("Unauthorized")
}

// ForbiddenError возвращает 403 Forbidden
func ForbiddenError() (int, *ErrorResponse) {
	return http.StatusForbidden, newError("Forbidden")
}
