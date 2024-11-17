package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type WithError struct {
	Error Error `json:"error"`
}

type WithData[T any] struct {
	Data T `json:"data"`
}

func Ok[T any](c echo.Context, data T) error {
	res := WithData[T]{
		Data: data,
	}
	return c.JSON(http.StatusOK, res)
}

func SendError(c echo.Context, status int, err Error) error {
	res := WithError{
		Error: err,
	}
	return c.JSON(status, res)
}

func BadInput(c echo.Context) error {
	return SendError(c, http.StatusBadRequest, Error{
		Code:    "BAD_INPUT",
		Message: "The input is invalid, please check again",
	})
}

func InternalServerError(c echo.Context) error {
	return SendError(c, http.StatusInternalServerError, Error{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Something went wrong, please try again later",
	})
}
