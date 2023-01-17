package repository

import "fmt"

type errKind int

var (
	ErrUniqueViolation = Error{kind: uniqueViolation}
	ErrDataNotFound    = Error{kind: dataNotFound}
	ErrUnknown         = Error{kind: unknown}
)

const (
	_ errKind = iota
	uniqueViolation
	dataNotFound
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case uniqueViolation:
		return fmt.Sprintf("Unique Violation %v", e.err)
	case dataNotFound:
		return fmt.Sprintf("Data Not Found %v", e.err)
	default:
		return fmt.Sprintf("Unknown Error %v", e.err)
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
