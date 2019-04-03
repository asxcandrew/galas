package errors

import "errors"

var (
	NotFoundError   = errors.New("Not found")
	BadRequestError = errors.New("Bad request")
)
