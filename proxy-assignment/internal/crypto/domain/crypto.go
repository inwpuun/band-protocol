package domain

import "inwpuun/proxy_assignment/internal/crypto/models"

type CryptoDomain interface {
	BroadcastTransaction(in models.BroadcastTransactionRequestDto) (models.BroadcastTransactionResponseDto, error)
	CheckStatus(hash string) (models.CheckStatusResponseDto, error)
}
