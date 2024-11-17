package repositories

import (
	"fmt"
	"inwpuun/proxy_assignment/config"
	"inwpuun/proxy_assignment/internal/crypto/models"
	"inwpuun/proxy_assignment/pkg/httputils"
)

type cryptoRepository struct {
	config *config.Config
}

func NewCryptoRepository(conf *config.Config) CryptoRepository {
	return &cryptoRepository{
		config: conf,
	}
}

func (r *cryptoRepository) BroadcastTransaction(in models.BroadcastTransactionRequestDto) (models.BroadcastTransactionResponseDto, error) {
	url := r.config.Server2ServerConfig.Url + "/broadcast"

	resp, err := httputils.PostWithOnlyResponseBody[models.BroadcastTransactionRequestDto, models.BroadcastTransactionResponseDto](url, in)
	if err != nil {
		return models.BroadcastTransactionResponseDto{}, fmt.Errorf("can't broadcast transaction in url %s from crypto repository: %w", url, err)
	}

	return resp, nil
}

func (r *cryptoRepository) CheckStatus(hash string) (models.CheckStatusResponseDto, error) {
	url := r.config.Server2ServerConfig.Url + "/check/" + hash

	resp, err := httputils.GetWithOnlyResponseBody[models.CheckStatusResponseDto](url)
	if err != nil {
		return models.CheckStatusResponseDto{}, fmt.Errorf("can't check status in url %s from crypto repository: %w", url, err)
	}

	return resp, nil
}
