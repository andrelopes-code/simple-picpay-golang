package errdefs

import "errors"

var (
	ErrRecordNotFound = errors.New(
		"record not found: sender or receiver does not exist",
	)
)
