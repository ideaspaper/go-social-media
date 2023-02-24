package usecase

import "fmt"

type errKind int

var (
	ErrUserAlreadyExists    = Error{kind: userAlreadyExists}
	ErrFailHashingPassword  = Error{kind: failHashingPassword}
	ErrFailToValidate       = Error{kind: failToValidate}
	ErrUserNotFound         = Error{kind: userNotFound}
	ErrWrongEmailOrPassword = Error{kind: wrongEmailOrPassword}
	ErrFailSigningJWT       = Error{kind: failSigningJWT}
	ErrUnknown              = Error{kind: unknown}
)

const (
	_ errKind = iota
	userAlreadyExists
	failHashingPassword
	failToValidate
	userNotFound
	wrongEmailOrPassword
	failSigningJWT
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case userAlreadyExists:
		return fmt.Sprintf("User already exists %v", e.err)
	case failHashingPassword:
		return fmt.Sprintf("Fail hashing password %v", e.err)
	case failToValidate:
		return fmt.Sprintf("%v", e.err)
	case userNotFound:
		return fmt.Sprintf("User not found %v", e.err)
	case wrongEmailOrPassword:
		return fmt.Sprintf("Wrong email or password %v", e.err)
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
