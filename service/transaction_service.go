package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andrelopes-code/simple-picpay-golang/cfg"
	"github.com/andrelopes-code/simple-picpay-golang/entity"
	"github.com/andrelopes-code/simple-picpay-golang/errdefs"
	"github.com/andrelopes-code/simple-picpay-golang/schemas"
)

// Create creates a new transaction in the database
func (s *TransactionService) Create(
	transaction *schemas.TransactionCreateRequest,
) (*entity.Transaction, error) {

	// Get sender and receiver users and check if they exist
	sender, receiver, err := s.getSenderAndReceiverUsers(transaction)
	if err != nil {
		return nil, err
	}

	// Check if transaction is valid
	err = s.validateTransaction(sender, receiver, transaction)
	if err != nil {
		return nil, err
	}

	// Check if transaction is authorized
	err = s.authorizeTransaction()
	if err != nil {
		return nil, err
	}

	sender.Debit(transaction.ValueInCents)
	receiver.Credit(transaction.ValueInCents)

	newTransaction := &entity.Transaction{
		ValueInCents: transaction.ValueInCents,
		SenderID:     transaction.SenderID,
		ReceiverID:   transaction.ReceiverID,
	}

	// Saves the transaction in the database atomically
	err = s.repository.SaveTransactionAtomically(
		newTransaction,
		sender,
		receiver,
	)
	if err != nil {
		return nil, err
	}

	// Send notification and return the new transaction
	go s.notificationService.SendNotification(newTransaction)
	return newTransaction, nil
}

// validateTransaction checks if the transaction is a valid transaction according to the business rules
func (s *TransactionService) validateTransaction(
	sender *entity.User,
	receiver *entity.User,
	transaction *schemas.TransactionCreateRequest,
) error {
	// Check if sender user is a merchant
	if sender.Type == entity.MerchantUser {
		return errdefs.ErrInvalidUser
	}
	// Check if sender has enough funds
	if sender.BalanceInCents < transaction.ValueInCents {
		return errdefs.ErrInsufficientFunds
	}
	// Check if sender and receiver are the same
	if sender.ID == receiver.ID {
		return errdefs.ErrSameUser
	}
	return nil
}

// GetSenderAndReceiverUsers returns the sender and receiver users
// RESPECTIVELY (Sender, Receiver). Also retuns an error if the users do not exist
func (s *TransactionService) getSenderAndReceiverUsers(
	transaction *schemas.TransactionCreateRequest,
) (*entity.User, *entity.User, error) {
	senderUser, err := s.repository.FindUserByID(transaction.SenderID)
	if err != nil {
		return nil, nil, errdefs.ErrRecordNotFound
	}
	receiverUser, err := s.repository.FindUserByID(transaction.ReceiverID)
	if err != nil {
		return nil, nil, errdefs.ErrRecordNotFound
	}
	return senderUser, receiverUser, nil
}

// AuthorizeTransaction checks if the transaction is authorized.
// If no error is returned, the transaction is authorized
func (s *TransactionService) authorizeTransaction() error {
	response, err := http.Get(cfg.Settings.AuthURL)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	auth := schemas.Authorization{}
	err = json.Unmarshal(data, &auth)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if !auth.Data.Authorization {
		return errdefs.ErrTransactionUnauthorized
	}

	return nil
}
