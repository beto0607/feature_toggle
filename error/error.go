package error

import (
	"net/http"
	"toggler/responses"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

const (
	NotFound      string = "NotFound"
	BadRequest    string = "BadRequest"
	InternalError string = "InternalError"
	Forbidden     string = "Forbidden"
	Unauthorized  string = "Unauthorized"
)

func SendException(c *gin.Context, exception string, message string) {
	error := exceptionToError(exception, message)
	SendError(c, error)
}

func SendError(c *gin.Context, error Error) {
	c.JSON(error.Status,
		responses.ErrorResponse{
			Status:  error.Status,
			Message: error.Message,
			Data:    error.Data})
}

func exceptionToError(exception string, message string) Error {
	error := Error{
		Message: exception,
	}
	if len(message) > 0 {
		error.Message = message
	}
	switch exception {
	case NotFound:
		error.Status = http.StatusNotFound
	case BadRequest:
		error.Status = http.StatusBadRequest
	case Forbidden:
		error.Status = http.StatusForbidden
	case Unauthorized:
		error.Status = http.StatusUnauthorized
	case InternalError:
		error.Status = http.StatusInternalServerError
	default:
		error.Status = http.StatusInternalServerError
	}

	return error
}
