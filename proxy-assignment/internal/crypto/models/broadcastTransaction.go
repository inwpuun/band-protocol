package models

type BroadcastTransactionRequestDto struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type BroadcastTransactionResponseDto struct {
	Hash string `json:"tx_hash"`
}
