package handler

import (
	"net/http"

	"github.com/andrelopes-code/simple-picpay-golang/errdefs"
	"github.com/andrelopes-code/simple-picpay-golang/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(ctx *gin.Context) {
	// binds the req body to TransactionCreateRequest struct
	// if a validation error occurs, it sends an error message
	req := new(schemas.TransactionCreateRequest)
	err := ctx.ShouldBind(req)
	if err != nil {
		logger.Debugf("%+v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newTransaction, err := transactionService.Create(req)
	if err != nil {
		// If an error occurs, it checks the type of error and sends an error message
		switch err {
		case errdefs.ErrRecordNotFound:
			sendError(ctx, http.StatusNotFound, err.Error())
		case errdefs.ErrInsufficientFunds:
			sendError(ctx, http.StatusBadRequest, err.Error())
		case errdefs.ErrSameUser:
			sendError(ctx, http.StatusBadRequest, err.Error())
		case errdefs.ErrInvalidUser:
			sendError(ctx, http.StatusBadRequest, err.Error())
		case errdefs.ErrTransactionUnauthorized:
			sendError(ctx, http.StatusForbidden, err.Error())

		default:
			logger.Error(err)
			sendError(
				ctx, http.StatusInternalServerError, "unexpected error",
			)
		}
		return
	}

	// If no error occurs, it sends a success message with the created transaction
	sendSuccess(ctx, http.StatusOK, "create-transaction", newTransaction)
}
