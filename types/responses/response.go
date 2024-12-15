package responses

import (
	"errors"
	"net/http"
)

type Response struct {
	Status int
	Msg    string
	Error  error
	Data   interface{}
}

func (r *Response) IsError() bool {
	return r.Error != nil
}

func NewUnsupportedMediaTypeResponse(msg string) *Response {
	return &Response{
		Status: http.StatusUnsupportedMediaType,
		Msg:    msg,
		Error:  errors.New(msg),
	}
}

func NewUnsupportedMethodResponse() *Response {
	return &Response{
		Status: http.StatusUnsupportedMediaType,
		Msg:    "Method not allowed",
		Error:  errors.New("Method not allowed"),
	}
}

func NewBadRequestResponse(msg string) *Response {
	return &Response{
		Status: http.StatusBadRequest,
		Msg:    msg,
		Error:  errors.New(msg),
	}
}

func NewInternalServerErrorResponse(msg string) *Response {
	return &Response{
		Status: http.StatusInternalServerError,
		Msg:    msg,
		Error:  errors.New(msg),
	}
}

func NewNotFoundResponse(msg string) *Response {
	return &Response{
		Status: http.StatusNotFound,
		Msg:    msg,
		Error:  errors.New(msg),
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Status: http.StatusOK,
		Data:   data,
	}
}
