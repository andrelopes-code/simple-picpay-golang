package entity

type UserType int8

const (
	CommonUser   UserType = 1
	MerchantUser UserType = 2
)

// Struct that represents a user in the database
type User struct {
	ID             uint64   `gorm:"primaryKey"  json:"id"`
	FullName       string   `                   json:"fullName"`
	DocumentID     string   `gorm:"uniqueIndex" json:"document_id"`
	Email          string   `gorm:"uniqueIndex" json:"email"`
	Password       string   `                   json:"password"`
	Type           UserType `                   json:"type"`
	BalanceInCents int64    `                   json:"balance_in_cents"`

	SenderTransactions   []Transaction `gorm:"foreignKey:SenderID;references:ID"   json:"sender_transactions"`
	ReceiverTransactions []Transaction `gorm:"foreignKey:ReceiverID;references:ID" json:"receiver_transactions"`
}

// Debit debits the user by the given amount
func (w *User) Debit(amount int64) *User {
	w.BalanceInCents -= amount
	return w
}

// Credit credits the user by the given amount
func (w *User) Credit(amount int64) *User {
	w.BalanceInCents += amount
	return w
}
