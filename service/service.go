package service

import (
	"net/http"

	"github.com/andrelopes-code/simple-picpay-golang/cfg"
	"github.com/andrelopes-code/simple-picpay-golang/entity"
	"github.com/andrelopes-code/simple-picpay-golang/repository"
	"gorm.io/gorm"
)

var (
	log = cfg.GetLogger("service")
)

// TransactionService is a service that handles transaction operations
type TransactionService struct {
	repository          *repository.Repository
	notificationService *NotificationService
}

// NewTransactionService creates a new TransactionService
func NewTransactionService(db *gorm.DB) *TransactionService {
	repository := repository.NewRepository(db)
	return &TransactionService{
		repository,
		NewNotificationService(),
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

// NotificationService is a service that handles notification operations
type NotificationService struct{}

// NewNotificationService creates a new NotificationService
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendNotification(transaction *entity.Transaction) {
	resp, err := http.Get(cfg.Settings.NotificationURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Errorf("Error sending notification: %s", err)
	}
	defer resp.Body.Close()
}
