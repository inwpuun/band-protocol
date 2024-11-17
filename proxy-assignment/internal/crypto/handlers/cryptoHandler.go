package handlers

import "github.com/labstack/echo/v4"

type CryptoHandler interface {
	BroadcastTransaction(c echo.Context) error
	CheckStatus(c echo.Context) error
}
