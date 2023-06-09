package response

import (
	"github.com/crosect/cc-go/exception"
	"net/http"
)

type Response struct {
	Meta Meta        `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func New(code int, message string, data interface{}) Response {
	return Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

func Ok(data interface{}) Response {
	return New(http.StatusOK, "Successful", data)
}

func Created(data interface{}) Response {
	return New(http.StatusCreated, "Resource created", data)
}

func Error(err error) Response {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	if e, ok := err.(exception.Exception); ok {
		code = int(e.Code())
		message = e.Message()
	}
	return Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
	}
}
