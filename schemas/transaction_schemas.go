package schemas

import "github.com/andrelopes-code/simple-picpay-golang/entity"

type TransactionResponse struct {
	entity.Transaction
}

type TransactionCreateRequest struct {
	ValueInCents int64  `json:"value_in_cents" binding:"required,min=1"`
	SenderID     uint64 `json:"sender_id"      binding:"required"`
	ReceiverID   uint64 `json:"receiver_id"    binding:"required"`
}

type TransactionUpdateRequest struct {
}

type Authorization struct {
	Status string `json:"status"`
	Data   struct {
		Authorization bool `json:"authorization"`
	}
}