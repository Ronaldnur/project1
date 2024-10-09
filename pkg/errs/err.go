package errs

import "net/http"

type HttpStatus int

const (
	StatusOK                  HttpStatus = 200
	StatusBadRequest          HttpStatus = 400
	StatusUnauthorized        HttpStatus = 401
	StatusForbidden           HttpStatus = 403
	StatusNotFound            HttpStatus = 404
	StatusInternalServerError HttpStatus = 500
	StatusUnprocessableEntity HttpStatus = 422
)

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type ErrorData struct {
	ErrMessage string     `json:"message"`
	ErrStatus  HttpStatus `json:"status"`
	ErrError   string     `json:"error"`
}

func (e *ErrorData) Message() string {
	return e.ErrMessage
}

func (e *ErrorData) Status() int {
	return int(e.ErrStatus)
}

func (e *ErrorData) Error() string {
	return e.ErrError
}

func NewUnauthorizedError(message string) MessageErr {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  StatusForbidden,
		ErrError:   "NOT_AUTHORIZED",
	}
}

func NewUnauthenticatedError(message string) MessageErr {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  StatusUnauthorized,
		ErrError:   "NOT_AUTHENTICATED",
	}
}

func NewNotFoundError(message string) MessageErr {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "NOT_FOUND",
	}
}

func NewBadRequest(message string) MessageErr {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) MessageErr {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError, //500
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "INVALID_REQUEST_BODY",
	}
}
