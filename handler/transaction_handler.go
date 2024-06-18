package handler

import (
	"net/http"

	"github.com/andrelopes-code/simple-picpay-golang/errdefs"
	"github.com/andrelopes-code/simple-picpay-golang/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Create Transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body schemas.TransactionCreateRequest true "Transaction data"
// @Success 201 {object} schemas.TransactionResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 404 {object} schemas.ErrorResponse
// @Failure 403 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /transaction [post]
func CreateTransaction(ctx *gin.Context) {
	// binds the req body to TransactionCreateRequest struct
	// if a validation error occurs, it sends an error message
	req := new(schemas.TransactionCreateRequest)
	err := ctx.ShouldBind(req)
	if err != nil {
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
	sendSuccess(ctx, http.StatusCreated, "create-transaction", newTransaction)
}
