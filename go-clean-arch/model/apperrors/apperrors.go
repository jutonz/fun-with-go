package apperrors

import (
	"net/http"
)

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}

func Status(err error) int {
	if e, ok := err.(*Error); ok {
		return e.Status
	}
	return http.StatusInternalServerError
}

func NewNotFound(message string) *Error {
	return &Error{
		Type:    "NOT_FOUND",
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func NewBadRequest() *Error {
	return &Error{
		Type:    "BAD_REQUEST",
		Status:  http.StatusBadRequest,
		Message: "Invalid request",
	}
}

func NewInternalError() *Error {
	return &Error{
		Type:    "INTERNAL_ERROR",
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
	}
}
