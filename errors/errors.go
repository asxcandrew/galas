package errors

import "errors"

var (
	NotFoundError   = errors.New("Not found")
	ForbiddenError  = errors.New("Forbidden action")
	BadRequestError = errors.New("Bad request")
)
