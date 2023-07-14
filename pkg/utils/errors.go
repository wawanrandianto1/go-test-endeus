package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Error struct {
	Success bool                   `json:"success"`
	Message interface{}            `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

func NewError(err error) Error {
	e := Error{Success: false}
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Message = v.Message
	default:
		e.Message = v.Error()
	}
	return e
}

func NewValidatorError(err error) Error {
	e := Error{Message: "validator error", Success: false}
	e.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Errors[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return e
}

func NotFound() Error {
	e := Error{Success: false, Message: "resource not found"}
	return e
}
