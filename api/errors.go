package api

import (
	"errors"
	"fmt"
)

// There are list of possible API errors.
var (
	ErrEncodeBody = errors.New("could not encode request body")
	ErrPrepareReq = errors.New("could not create request")
	ErrSendReq    = errors.New("could not send request")
	ErrReqTimeout = errors.New("request timeout")
	ErrResponse   = errors.New("server returned unexpected result")
	ErrDecodeBody = errors.New("could not parse response body")
)

// Err represents API error general struct.
type Err struct {
	apiErr  error
	baseErr error
}

// Unwrap unwraps API error.
func (e Err) Unwrap() error {
	return e.apiErr
}

func (e Err) Error() string {
	return fmt.Sprintf("%v: %v", e.apiErr, e.baseErr)
}

func newErr(apiErr, baseErr error) error {
	return Err{
		apiErr:  apiErr,
		baseErr: baseErr,
	}
}
