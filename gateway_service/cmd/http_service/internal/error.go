package internal

import (
	"fmt"
)

type errKind int

var (
	ErrCors           = Error{kind: cors}
	ErrNoRoute        = Error{kind: noRoute}
	ErrBadParams      = Error{kind: badParams}
	ErrFailToValidate = Error{kind: failToValidate}
	ErrUnknown        = Error{kind: unknown}
)

const (
	_ errKind = iota
	cors
	noRoute
	badParams
	failToValidate
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case cors:
		return "CORS error"
	case noRoute:
		return "No route"
	case badParams:
		return "Bad params"
	case failToValidate:
		return fmt.Sprintf("%v", e.err)
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
