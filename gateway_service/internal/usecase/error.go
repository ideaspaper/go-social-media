package usecase

import "fmt"

type errKind int

var (
	ErrClientService  = Error{kind: clientService}
	ErrFailToValidate = Error{kind: failToValidate}
	ErrFailSigningJWT = Error{kind: failSigningJWT}
	ErrUnknown        = Error{kind: unknown}
)

const (
	_ errKind = iota
	clientService
	failToValidate
	failSigningJWT
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case clientService:
		return fmt.Sprintf("Client service error %v", e.err)
	case failToValidate:
		return fmt.Sprintf("Fail to validate %v", e.err)
	case failSigningJWT:
		return fmt.Sprintf("Fail signing JWT %v", e.err)
	default:
		return fmt.Sprintf("Unknown error %v", e.err)
	}
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Is(err error) bool {
	target, ok := err.(*Error)
	if !ok {
		return false
	}
	return target.kind == e.kind
}

func (e *Error) SetError(err error) *Error {
	e.err = err
	return e
}
