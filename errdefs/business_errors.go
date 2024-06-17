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
		"invalid user: ony individual users can transact",
	)
	ErrTransactionUnauthorized = errors.New(
		"unauthorized: transaction not authorized",
	)
)
