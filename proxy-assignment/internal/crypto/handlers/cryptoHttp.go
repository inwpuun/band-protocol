package handlers

import (
	"inwpuun/proxy_assignment/internal/api"
	"inwpuun/proxy_assignment/internal/crypto/domain"
	"inwpuun/proxy_assignment/internal/crypto/models"

	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
)

type cryptoHttpHandler struct {
	cryptoDomain domain.CryptoDomain
}

func NewCryptoHttpHandler(cryptoDomain domain.CryptoDomain) CryptoHandler {
	return &cryptoHttpHandler{
		cryptoDomain: cryptoDomain,
	}
}

func (h *cryptoHttpHandler) BroadcastTransaction(c echo.Context) error {
	reqBody := new(models.BroadcastTransactionRequestDto)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return api.BadInput(c)
	}

	data, err := h.cryptoDomain.BroadcastTransaction(*reqBody)
	if err != nil {
		log.Error("Error broadcasting transaction: ", err)
		return api.InternalServerError(c)
	}

	response := models.BroadcastTransactionResponseDto{
		Hash: data.Hash,
	}

	return api.Ok(c, response)
}

func (h *cryptoHttpHandler) CheckStatus(c echo.Context) error {
	hash := c.Param("hash")

	// Validate the path parameter if necessary
	if hash == "" {
		log.Errorf("Error: Missing or invalid path parameter 'id'")
		return api.BadInput(c)
	}

	data, err := h.cryptoDomain.CheckStatus(hash)
	if err != nil {
		log.Error("Error checking status: ", err)
		return api.InternalServerError(c)
	}

	response := models.CheckStatusResponseDto{
		Status: data.Status,
	}

	return api.Ok(c, response)
}
