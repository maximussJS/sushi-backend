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

func NewContentTooLargeResponse() *Response {
	return &Response{
		Status: http.StatusRequestEntityTooLarge,
		Msg:    "Request entity too large",
		Error:  errors.New("Request entity too large"),
	}
}

func NewBadRequestResponse(msg string) *Response {
	return &Response{
		Status: http.StatusBadRequest,
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

func NewForbiddenResponse() *Response {
	return &Response{
		Status: http.StatusForbidden,
		Msg:    "Forbidden",
		Error:  errors.New("forbidden"),
	}
}

func NewUnauthorizedResponse(msg string) *Response {
	return &Response{
		Status: http.StatusUnauthorized,
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
