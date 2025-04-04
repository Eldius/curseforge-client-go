package types

import (
	"errors"
	"fmt"
)

const (
	ErrAPIBadRequestMsg        = "api returned bad request"
	ErrRequestErrorMsg         = "could not make request"
	ErrAPIServerErrorMsg       = "server error"
	ErrInvalidRequestParamsMsg = "invalid request params"
)

var (
	// ErrAPIBadRequest represents a Curseforge API bad request error
	ErrAPIBadRequest error = Wrap(nil, ErrAPIBadRequestMsg, 400)
	// ErrAPIServerError represents a Curseforge API bad request error
	ErrAPIServerError error = Wrap(nil, ErrAPIServerErrorMsg, 500)
	// ErrRequestError represents general a connectivity error
	ErrRequestError error = Wrap(nil, ErrRequestErrorMsg, 0)
	// ErrInvalidRequestParams represents a validation error in the parameters
	ErrInvalidRequestParams error = Wrap(nil, ErrInvalidRequestParamsMsg, 0)
)

type CurseforgeAPIError struct {
	error
	Status  int
	Message string
	Err     error
}

func (e CurseforgeAPIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Err)
}

func (e CurseforgeAPIError) Is(err error) bool {
	var t *CurseforgeAPIError
	ok := errors.As(err, &t)
	if !ok {
		return false
	}
	return t.Message == e.Message || t.Status == e.Status
}

func Wrap(err error, msg string, status int) CurseforgeAPIError {
	return CurseforgeAPIError{
		Status:  status,
		Message: msg,
		Err:     err,
	}
}
