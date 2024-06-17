package service

import (
	"github.com/andrelopes-code/simple-picpay-golang/repository"
	"gorm.io/gorm"
)

// TransactionService is a service that handles transaction operations
type TransactionService struct {
	repository *repository.Repository
}

// NewTransactionService creates a new TransactionService
func NewTransactionService(db *gorm.DB) *TransactionService {
	repository := repository.NewRepository(db)
	return &TransactionService{
		repository,
	}
}

// UserService is a service that handles user operations
type UserService struct {
	repository *repository.Repository
}

// NewUserService creates a new UserService
func NewUserService(db *gorm.DB) *UserService {
	repository := repository.NewRepository(db)
	return &UserService{
		repository,
	}
}
