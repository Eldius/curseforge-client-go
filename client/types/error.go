package types

import "fmt"

const (
    ErrAPIBadRequestMsg  = "api returned bad request"
    ErrRequestErrorMsg   = "could not make request"
    ErrAPIServerErrorMsg = "server error"
)

var (
    // ErrAPIBadRequest represents a Curseforge API bad request error
    ErrAPIBadRequest = Wrap(nil, ErrAPIBadRequestMsg, 400)
    // ErrAPIServerError represents a Curseforge API bad request error
    ErrAPIServerError = Wrap(nil, ErrAPIServerErrorMsg, 500)
    // ErrRequestError represents general a connectivity error
    ErrRequestError = Wrap(nil, ErrRequestErrorMsg, 0)
)

type CurseforgeAPIError struct {
    Status  int
    Message string
    Err     error
}

func (e CurseforgeAPIError) Error() string {
    return fmt.Sprintf("%s: %s", e.Message, e.Err)
}

func (e CurseforgeAPIError) Is(err error) bool {
    t, ok := err.(*CurseforgeAPIError)
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
