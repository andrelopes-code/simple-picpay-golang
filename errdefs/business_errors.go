package errdefs

import "errors"

var (
	ErrInsufficientFunds = errors.New(
		"insufficient funds: sender does not have enough funds",
	)
	ErrSameUser = errors.New(
		"same user: sender and receiver cannot be the same",
	)
	ErrInvalidUser = errors.New(
		"invalid user: merchants is not allowed to trasfer",
	)
	ErrTransactionUnauthorized = errors.New(
		"unauthorized: transaction not authorized",
	)
)
