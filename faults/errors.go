package faults

const (
	NotFoundError     = "Not found"
	ForbiddenError    = "Forbidden action"
	UnauthorisedError = "Unauthorised action"
	BadRequestError   = "Bad request"
)

type RichError struct {
	original error
	text     string
}

type IRichError interface {
	Error() string
	Original() string
}

func (e RichError) Error() string {
	return e.text
}

func (e RichError) Original() string {
	return e.original.Error()
}

func BuildRichError(template string, original error) IRichError {
	return RichError{
		text:     template,
		original: original,
	}
}
