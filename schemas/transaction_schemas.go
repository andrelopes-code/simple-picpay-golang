package schemas

import "time"

type TransactionCreateRequest struct {
	ValueInCents int64  `json:"value_in_cents" binding:"required,min=1"`
	SenderID     uint64 `json:"sender_id"      binding:"required"`
	ReceiverID   uint64 `json:"receiver_id"    binding:"required"`
}

type TransactionResponse struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	ID           uint64    `json:"id"`
	ValueInCents int64     `json:"value_in_cents"`
	SenderID     uint64    `json:"sender_id"`
	ReceiverID   uint64    `json:"receiver_id"`
}

type Authorization struct {
	Status string `json:"status"`
	Data   struct {
		Authorization bool `json:"authorization"`
	}
}
