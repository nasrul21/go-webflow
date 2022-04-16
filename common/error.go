package common

import (
	"encoding/json"
	"net/http"
)

const (
	APIValidationError string = "API_VALIDATION_ERROR"
	GoErrCode          string = "GO_ERROR"
)

type Error struct {
	Code     int         `json:"code"`
	Message  string      `json:"msg"`
	Err      string      `json:"err"`
	Name     string      `json:"name,omitempty"`
	Path     string      `json:"path,omitempty"`
	Problems interface{} `json:"problems,omitempty"`
}

// FromGoErr generates xendit.Error from generic go errors
func FromGoErr(err error) *Error {
	return &Error{
		Code:    http.StatusTeapot,
		Err:     GoErrCode,
		Message: err.Error(),
	}
}

// FromHTTPErr generates xendit.Error from http errors with non 2xx status
func FromHTTPErr(status int, respBody []byte) *Error {
	var httpError *Error
	if err := json.Unmarshal(respBody, &httpError); err != nil {
		return FromGoErr(err)
	}
	httpError.Code = status

	return httpError
}
