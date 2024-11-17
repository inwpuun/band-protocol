package domain

import (
	"fmt"
	"inwpuun/proxy_assignment/internal/crypto/models"
	"inwpuun/proxy_assignment/internal/crypto/repositories"
)

type cryptoDomainImpl struct {
	cryptoRepository repositories.CryptoRepository
}

func NewCryptoDomainImpl(
	cryptoRepository repositories.CryptoRepository,
) CryptoDomain {
	return &cryptoDomainImpl{
		cryptoRepository: cryptoRepository,
	}
}

func (u *cryptoDomainImpl) BroadcastTransaction(in models.BroadcastTransactionRequestDto) (models.BroadcastTransactionResponseDto, error) {
	txhash, err := u.cryptoRepository.BroadcastTransaction(in)
	if err != nil {
		return models.BroadcastTransactionResponseDto{}, fmt.Errorf("can't broadcast transaction from crypto domain: %w", err)
	}

	return txhash, nil
}

func (u *cryptoDomainImpl) CheckStatus(hash string) (models.CheckStatusResponseDto, error) {
	status, err := u.cryptoRepository.CheckStatus(hash)
	if err != nil {
		return models.CheckStatusResponseDto{}, fmt.Errorf("can't check status from crypto domain: %w", err)
	}

	return status, nil
}
