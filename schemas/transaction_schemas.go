package schemas

type TransactionResponse struct {
	ID uint64 `json:"id"`
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
