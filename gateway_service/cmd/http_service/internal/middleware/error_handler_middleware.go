package middleware

import (
	"errors"
	"gatewayservice/cmd/http_service/internal"
	"gatewayservice/internal/dto/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var grpcToHttp = map[codes.Code]int{
	codes.InvalidArgument:  http.StatusBadRequest,
	codes.Unauthenticated:  http.StatusUnauthorized,
	codes.PermissionDenied: http.StatusForbidden,
	codes.NotFound:         http.StatusNotFound,
	codes.AlreadyExists:    http.StatusConflict,
	codes.Unknown:          http.StatusInternalServerError,
}

func (m Middleware) ErrorHandler(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) == 0 {
		return
	}
	code := http.StatusInternalServerError
	body := &resp.StandardDto{
		Code:    code,
		Message: http.StatusText(code),
		Data:    nil,
	}
	firstErr := ctx.Errors[0].Err
	if errors.Is(firstErr, &internal.ErrBadParams) {
		code = http.StatusBadRequest
		body = &resp.StandardDto{
			Code:    code,
			Message: firstErr.Error(),
			Data:    nil,
		}
	} else if errors.Is(firstErr, &internal.ErrFailToValidate) {
		code = http.StatusBadRequest
		body = &resp.StandardDto{
			Code:    code,
			Message: errors.Unwrap(firstErr).Error(),
			Data:    nil,
		}
	} else if errors.Is(firstErr, &internal.ErrNoRoute) {
		code = http.StatusNotFound
		body = &resp.StandardDto{
			Code:    code,
			Message: "Oops... nothing here",
			Data:    nil,
		}
	} else {
		grpcStatus, ok := status.FromError(firstErr)
		if ok {
			code = grpcToHttp[grpcStatus.Code()]
			body = &resp.StandardDto{
				Code:    code,
				Message: grpcStatus.Message(),
				Data:    nil,
			}
		}
	}
	ctx.AbortWithStatusJSON(code, body)
}
