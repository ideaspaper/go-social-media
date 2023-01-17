package interceptor

import (
	"errors"
	"userservice/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i Interceptor) ErrorHandler(err error) error {
	code := codes.Unknown
	message := "Unknown"
	if errors.Is(err, &usecase.ErrFailToValidate) {
		code = codes.InvalidArgument
		message = errors.Unwrap(err).Error()
	} else if errors.Is(err, &usecase.ErrUserAlreadyExists) {
		code = codes.AlreadyExists
		message = "Email already been registered"
	} else if errors.Is(err, &usecase.ErrUserNotFound) {
		code = codes.NotFound
		message = "User not found"
	}
	return status.Error(code, message)
}
