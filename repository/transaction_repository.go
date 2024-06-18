package repository

import (
	"github.com/andrelopes-code/simple-picpay-golang/entity"
	"gorm.io/gorm"
)

// FindByID returns a transaction by its ID
func (r *Repository) FindTransactionByID(
	id uint64,
) (*entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Create creates a new transaction in the database
func (r *Repository) SaveTransaction(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}

// TXSaveTransaction saves a transaction in the database atomically
func (r *Repository) SaveTransactionAtomically(
	transaction *entity.Transaction, sender *entity.User, receiver *entity.User,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(transaction).Error; err != nil {
			return err
		}
		if err := tx.Save(sender).Error; err != nil {
			return err
		}
		if err := tx.Save(receiver).Error; err != nil {
			return err
		}
		return nil
	})
}
