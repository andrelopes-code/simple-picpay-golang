package handler

import (
	"github.com/andrelopes-code/simple-picpay-golang/cfg"
	"github.com/andrelopes-code/simple-picpay-golang/service"
)

var (
	logger             *cfg.Logger
	transactionService *service.TransactionService
)

func Init() {
	logger = cfg.GetLogger("handler")
	logger.Info("Initializing handler...")
	db := cfg.GetDB()
	transactionService = service.NewTransactionService(db)
}
