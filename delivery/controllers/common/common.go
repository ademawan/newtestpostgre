package common

import "net/http"

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func InternalServerError() Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Message: "There is some error on server",
		Data:    nil,
	}
}

func BadRequest() Response {
	return Response{
		Code:    http.StatusBadRequest,
		Message: "There is some problem from input",
		Data:    nil,
	}
}

func Success(code int, msg string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
func NotFound(code int, msg string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
