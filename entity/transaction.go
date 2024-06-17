package entity

import (
	"time"

	"gorm.io/gorm"
)

// Struct that represents a transaction in the database
type Transaction struct {
	ID           uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `                  json:"createdAt"`
	UpdatedAt    time.Time      `                  json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index"      json:"deletedAt"`
	ValueInCents int64          `                  json:"value_in_cents"`
	SenderID     uint64         `gorm:"foreignKey" json:"sender_id"`
	ReceiverID   uint64         `gorm:"foreignKey" json:"receiver_id"`
}
